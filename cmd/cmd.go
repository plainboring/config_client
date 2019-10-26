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

	defaultConfigClient cfgclient.ConfigClient
)

const (
	// FlagPD is the name of url flag.
	FlagPD = "pd"
)

// AddFlags adds flags to the given cmd.
func AddFlags(cmd *cobra.Command) {
	cmd.PersistentFlags().StringP(FlagPD, "u", "", "PD address")
	cmd.PersistentFlags().BoolP("mock", "m", false, "use mock config client")
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
		mock, err := cmd.Flags().GetBool("mock")
		if err != nil {
			return
		}
		if addr == "" && !mock {
			err = errors.Errorf("pd address can not be empty")
			return
		}

		if mock {
			defaultConfigClient, err = cfgclient.NewMockClient()
		} else {
			defaultConfigClient, err = cfgclient.NewConfigClient(defaultContext, addr)
		}
		if err != nil {
			return
		}
	})
	return
}
