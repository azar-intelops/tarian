package cmd

import (
	"os"

	"github.com/kube-tarian/tarian/cmd/tarian-server/cmd/dgraph"
	"github.com/kube-tarian/tarian/cmd/tarian-server/cmd/flags"
	"github.com/kube-tarian/tarian/pkg/log"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var globalFlags *flags.GlobalFlags

func newRootCommand(logger *logrus.Logger) *cobra.Command {
	var rootCmd = &cobra.Command{
		Use:           "tarian-server",
		Short:         "The Tarian Server is the central component which manages config DB, users, etc.",
		Version:       versionStr,
		SilenceUsage:  true,
		SilenceErrors: true,
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			err := globalFlags.ValidateGlobalFlags()
			if err != nil {
				return err
			}

			logLevel, _ := logrus.ParseLevel(globalFlags.LogLevel)
			logger.SetLevel(logrus.Level(logLevel))

			if globalFlags.LogFormatter == "json" {
				logger.SetFormatter(&logrus.JSONFormatter{})
			}
			return nil
		},
	}

	return rootCmd
}

func buildRootCommand(logger *logrus.Logger) *cobra.Command {
	rootCmd := newRootCommand(logger)
	rootCmd.SetVersionTemplate("tarianctl version: {{.Version}}\n")

	// Add global flags
	persistentFlags := rootCmd.PersistentFlags()
	globalFlags = flags.SetGlobalFlags(persistentFlags)

	// Add subcommand to the root command
	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(newRunCommand(globalFlags))
	rootCmd.AddCommand(dgraph.NewDgraphCommand(globalFlags))

	return rootCmd
}

func Execute() {
	logger := log.GetLogger()
	rootCmd := buildRootCommand(logger)
	err := rootCmd.Execute()
	if err != nil {
		logger.Errorf("command failed: %s", err)
		os.Exit(1)
	}
}