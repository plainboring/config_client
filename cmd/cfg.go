package cmd

import (
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
			err = defaultConfigClient.Update(comp, name, value, store)
			if err != nil {
				return err
			}
			return nil
		},
	}
	command.Flags().StringP("comp", "c", "tikv", "update component config")
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
