package cfgclient

import (
	"bytes"

	"github.com/BurntSushi/toml"
	"github.com/pingcap/errors"
	"github.com/plainboring/config_client/pkg/pd"
	"github.com/plainboring/config_client/pkg/tikv"
)

type mockClient struct {
	// TODO: support PD
	tikvConfig tikv.Config
	pdConfig   pd.Config
}

// NewMockClient creates a new ConfigClient.
func NewMockClient() (ConfigClient, error) {
	var tikvConfig tikv.Config
	_, err := toml.Decode(tikv.ConfigFixture, &tikvConfig)
	if err != nil {
		return nil, errors.Trace(err)
	}
	var pdConfig pd.Config
	_, err = toml.Decode(tikv.ConfigFixture, &pdConfig)
	if err != nil {
		return nil, errors.Trace(err)
	}
	return &mockClient{
		tikvConfig: tikvConfig,
		pdConfig:   pdConfig,
	}, nil
}

// Get config
func (cli *mockClient) Get(comp string, storeID uint64) (string, error) {
	if comp == "tikv" {
		kvcfg := &tikv.Config{}
		reader := bytes.NewReader([]byte(tikv.ConfigFixture))
		_, err := toml.DecodeReader(reader, kvcfg)
		if err != nil {
			return "", errors.Trace(err)
		}
		writer := bytes.NewBuffer(make([]byte, 0, len(tikv.ConfigFixture)))
		encoder := toml.NewEncoder(writer)
		err = encoder.Encode(kvcfg)
		if err != nil {
			return "", errors.Trace(err)
		}
		return writer.String(), nil
	}
	pdcfg := &pd.Config{}
	reader := bytes.NewReader([]byte(pd.ConfigFixture))
	_, err := toml.DecodeReader(reader, pdcfg)
	if err != nil {
		return "", errors.Trace(err)
	}
	writer := bytes.NewBuffer(make([]byte, 0, len(pd.ConfigFixture)))
	encoder := toml.NewEncoder(writer)
	err = encoder.Encode(pdcfg)
	if err != nil {
		return "", errors.Trace(err)
	}
	return writer.String(), nil
}

// Update config
func (cli *mockClient) Update(
	comp string, subs []string, name, value string, storeID uint64) error {

	return nil
}
