package cmd

import (
	"os"

	"github.com/piltismart/tb-stack-ps-cli/pkg/ui"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "ps",
	Short: "PiltiSmart Enterprise CLI - ThingsBoard Stack Orchestration",
	Long: `==================================================================
  PiltiSmart 'ps' CLI Utility (Go + Cobra Framework)
  Enterprise automation tool to inspect host health and deploy
  the 3-Component ThingsBoard Microservice Stack from GitHub.
==================================================================`,
	Run: func(cmd *cobra.Command, args []string) {
		ui.PrintBanner("ThingsBoard Stack Automation Engine")
		cmd.Help()
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		ui.Error("%v", err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(doctorCmd)
	rootCmd.AddCommand(installCmd)
	rootCmd.AddCommand(catalogCmd)
	rootCmd.AddCommand(stackCmd)
	rootCmd.AddCommand(statusCmd)
	rootCmd.AddCommand(versionCmd)
}
