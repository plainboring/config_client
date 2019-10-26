package cfgclient

import (
	"context"
	"time"

	"github.com/pingcap/errors"
	"github.com/pingcap/kvproto/pkg/configpb"
	"github.com/pingcap/log"
	pd "github.com/pingcap/pd/client"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
)

const (
	dialTimeout = 5 * time.Second
)

// ConfigClient is the interface of a config client.
type ConfigClient interface {
	Get(comp string, storeID uint64) (string, error)
	Update(comp string, subs []string, name, value string, storeID uint64) error
}

// configClient the client of a TiDB/TiKV Config service.
type configClient struct {
	ctx       context.Context
	pdClient  pd.Client
	pdAddr    string
	clusterID uint64
}

// NewConfigClient creates a new ConfigClient.
func NewConfigClient(ctx context.Context, pdAddr string) (ConfigClient, error) {
	log.Info("connect pd", zap.String("addr", pdAddr))
	pdClient, err := pd.NewClient([]string{pdAddr}, pd.SecurityOption{})
	if err != nil {
		return nil, errors.Trace(err)
	}

	clusterID := pdClient.GetClusterID(ctx)
	return &configClient{
		ctx:       ctx,
		pdClient:  pdClient,
		pdAddr:    pdAddr,
		clusterID: clusterID,
	}, nil
}

// Client returns a ConfigClient.
func (cli *configClient) Client(pdAddr string) (configpb.ConfigClient, error) {
	opt := grpc.WithInsecure()
	dailCtx, cancel := context.WithTimeout(cli.ctx, dialTimeout)
	keepAlive := 10
	keepAliveTimeout := 3
	conn, err := grpc.DialContext(
		dailCtx,
		pdAddr,
		opt,
		grpc.WithBackoffMaxDelay(time.Second*3),
		grpc.WithKeepaliveParams(keepalive.ClientParameters{
			Time:                time.Duration(keepAlive) * time.Second,
			Timeout:             time.Duration(keepAliveTimeout) * time.Second,
			PermitWithoutStream: true,
		}),
	)
	cancel()
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return configpb.NewConfigClient(conn), nil
}

// Get config
func (cli *configClient) Get(comp string, storeID uint64) (string, error) {
	ctx, cancel := context.WithCancel(cli.ctx)
	defer cancel()
	client, err := cli.Client(cli.pdAddr)
	if err != nil {
		return "", errors.Trace(err)
	}
	req := &configpb.GetRequest{
		Header: &configpb.RequestHeader{
			ClusterId: cli.clusterID,
		},
		Component: ParseComponent(comp),
	}
	resp, err := client.Get(ctx, req)
	if err != nil {
		return "", errors.Trace(err)
	}
	log.Info("config get", zap.String("config", resp.Config))
	return resp.Config, nil
}

// Update config
func (cli *configClient) Update(
	comp string, subs []string, name, value string, storeID uint64) error {
	ctx, cancel := context.WithCancel(cli.ctx)
	defer cancel()
	client, err := cli.Client(cli.pdAddr)
	if err != nil {
		return errors.Trace(err)
	}
	req := &configpb.UpdateRequest{
		Header: &configpb.RequestHeader{
			ClusterId: cli.clusterID,
		},
		Component: ParseComponent(comp),
		Entry: &configpb.ConfigEntry{
			Subsystem: subs,
			Name:      name,
			Value:     value,
		},
	}
	log.Info("config update", zap.Reflect("request", req))
	resp, err := client.Update(ctx, req)
	if err != nil {
		return errors.Trace(err)
	}
	log.Info("config update", zap.Reflect("response", resp))
	return nil
}

// ParseComponent parse a Component string
func ParseComponent(str string) configpb.Component {
	if str == "tikv" {
		return configpb.Component_TiKV
	} else if str == "pd" {
		return configpb.Component_PD
	} else {
		log.Fatal("unknown component", zap.String("component", str))
		return configpb.Component_UNKNOWN
	}
}
