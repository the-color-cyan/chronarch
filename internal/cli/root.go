package cli

import (
	"github.com/spf13/cobra"
)

// base command when called without any subcommands
func NewRootCmd() *cobra.Command {
	root := &cobra.Command{
		Use:   "chronarch",
		Short: "time tracking and management.",
		Long: `time tracking and management
BOTTOM TEXT`,
		// Run: func(cmd *cobra.Command, args []string) { },
	}

	root.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	AddChildCommands(root)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.chronarch.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.

	return root
}

func AddChildCommands(root *cobra.Command) {
	root.AddCommand(NewProjectCmd())
	root.AddCommand(NewSessionCmd())
}
