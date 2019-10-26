package cfgclient

import (
	"bytes"

	"github.com/BurntSushi/toml"
	"github.com/pingcap/errors"
	"github.com/plainboring/config_client/pkg/tikv"
)

type mockClient struct {
	// TODO: support PD
	tikvConfig tikv.Config
}

// NewMockClient creates a new ConfigClient.
func NewMockClient() (ConfigClient, error) {
	var tikvConfig tikv.Config
	_, err := toml.Decode(tikv.ConfigFixture, &tikvConfig)
	if err != nil {
		return nil, errors.Trace(err)
	}
	return &mockClient{
		tikvConfig: tikvConfig,
	}, nil
}

// Get config
func (cli *mockClient) Get(comp string, storeID uint64) (string, error) {
	writer := bytes.NewBuffer(make([]byte, 0, len(tikv.ConfigFixture)))
	encoder := toml.NewEncoder(writer)
	encoder.Encode(cli.tikvConfig)
	return writer.String(), nil
}

// Update config
func (cli *mockClient) Update(
	comp string, subs []string, name, value string, storeID uint64) error {

	return nil
}
