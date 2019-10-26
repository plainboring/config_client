package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

// NewConfigommand return a cfg subcommand.
func NewConfigommand() *cobra.Command {
	bp := &cobra.Command{
		Use:   "cfg",
		Short: "Get/update cluster config!",
	}
	bp.AddCommand(
		newUpdateCfgCommand(),
		newGetCfgCommand(),
	)
	return bp
}

// newUpdateCfgCommand return a update config subcommand.
func newUpdateCfgCommand() *cobra.Command {
	command := &cobra.Command{
		Use:   "update",
		Short: "update the given config",
		RunE: func(command *cobra.Command, _ []string) error {
			comp, err := command.Flags().GetString("comp")
			if err != nil {
				return err
			}
			sub, err := command.Flags().GetString("subsystem")
			if err != nil {
				return err
			}
			subs := strings.Split(sub, ",")
			name, err := command.Flags().GetString("name")
			if err != nil {
				return err
			}
			value, err := command.Flags().GetString("value")
			if err != nil {
				return err
			}
			store, err := command.Flags().GetUint64("store")
			if err != nil {
				return err
			}
			err = defaultConfigClient.Update(comp, subs, name, value, store)
			if err != nil {
				return err
			}
			return nil
		},
	}

	command.Flags().StringP("comp", "c", "tikv", "update component config")
	command.Flags().StringP("subsystem", "b", "raftstore", "update a subsystem config")
	command.Flags().StringP("name", "k", "", "update the given config")
	command.Flags().StringP("value", "v", "",
		"update the given config with the value")
	command.Flags().Uint64P("store", "s", 0,
		"update the given store ids value")
	command.MarkFlagRequired("comp")
	command.MarkFlagRequired("name")
	command.MarkFlagRequired("value")
	return command
}

// newGetCfgCommand return a update config subcommand.
func newGetCfgCommand() *cobra.Command {
	command := &cobra.Command{
		Use:   "get",
		Short: "get config",
		RunE: func(command *cobra.Command, _ []string) error {
			comp, err := command.Flags().GetString("comp")
			if err != nil {
				return err
			}
			store, err := command.Flags().GetUint64("store")
			if err != nil {
				return err
			}
			// mock, err := command.Flags().GetBool("mock")
			// if err != nil {
			// 	return err
			// }
			config, err := defaultConfigClient.Get(comp, store)
			fmt.Println(config)
			if err != nil {
				return err
			}
			return nil
		},
	}

	command.Flags().StringP("comp", "c", "tikv", "update component config")
	command.Flags().Uint64P("store", "s", 1,
		"update the given store ids value")
	return command
}
