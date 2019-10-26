package cmd

import (
	"context"
	"sync"

	"github.com/pingcap/errors"
	"github.com/plainboring/config_client/pkg/cfgclient"
	"github.com/spf13/cobra"
)

var (
	initOnce       = sync.Once{}
	defaultContext context.Context

	defaultConfigClient *cfgclient.ConfigClient
)

const (
	// FlagPD is the name of url flag.
	FlagPD = "pd"
)

// AddFlags adds flags to the given cmd.
func AddFlags(cmd *cobra.Command) {
	cmd.PersistentFlags().StringP(FlagPD, "u", "127.0.0.1:2379", "PD address")
	cmd.MarkFlagRequired(FlagPD)
}

// Init ...
func Init(ctx context.Context, cmd *cobra.Command) (err error) {
	initOnce.Do(func() {
		defaultContext = ctx
		var addr string
		addr, err = cmd.Flags().GetString(FlagPD)
		if err != nil {
			return
		}
		if addr == "" {
			err = errors.Errorf("pd address can not be empty")
			return
		}
		defaultConfigClient, err = cfgclient.NewConfigClient(defaultContext, addr)
		if err != nil {
			return
		}
	})
	return
}
