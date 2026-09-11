package cmd

import (
	"fmt"
	"os"

	"github.com/piltismart/tb-stack-ps-cli/pkg/stack"
	"github.com/piltismart/tb-stack-ps-cli/pkg/ui"
	"github.com/spf13/cobra"
)

var (
	deployDir    string
	repoURL      string
	edgeWebPort  int
	edgeMqttPort int
)

var installCmd = &cobra.Command{
	Use:   "install [tb-stack]",
	Short: "Install and orchestrate stacks dynamically from GitHub catalog",
	Long:  "Pulls dynamic compose manifests from GitHub repository (https://github.com/PiltiSmart/stack-catalog) and orchestrates TimescaleDB, ThingsBoard Core, and ThingsBoard Edge.",
	Args:  cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		target := "tb-stack"
		if len(args) > 0 {
			target = args[0]
		}

		if target != "tb-stack" && target != "all" {
			ui.Error("Unknown install target '%s'. Supported target: 'tb-stack'", target)
			fmt.Println("Usage: ps install tb-stack")
			os.Exit(1)
		}

		err := stack.InstallTBStack(deployDir, repoURL, edgeWebPort, edgeMqttPort)
		if err != nil {
			ui.Error("Installation failed: %v", err)
			os.Exit(1)
		}
	},
}

func init() {
	installCmd.Flags().StringVarP(&deployDir, "dir", "d", "/opt/piltismart/tb-stack", "Target directory where compose files are provisioned")
	installCmd.Flags().StringVarP(&repoURL, "repo", "r", stack.DefaultCatalogRepo, "Catalog raw repository URL")
	installCmd.Flags().IntVar(&edgeWebPort, "edge-web-port", 0, "Host port for ThingsBoard Edge Web UI (default: auto 8080 or 8082 if occupied)")
	installCmd.Flags().IntVarP(&edgeMqttPort, "edge-mqtt-port", "p", 0, "Host port for ThingsBoard Edge MQTT broker (default: auto 1884 to prevent clash with Core 1883)")
}
