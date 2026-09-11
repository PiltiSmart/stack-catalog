package cmd

import (
	"github.com/piltismart/tb-stack-ps-cli/pkg/stack"
	"github.com/spf13/cobra"
)

var (
	stackDir       string
	stackComponent string
)

var stackCmd = &cobra.Command{
	Use:   "stack [status|down|restart]",
	Short: "Manage ThingsBoard stack lifecycle",
	Run: func(cmd *cobra.Command, args []string) {
		action := "status"
		if len(args) > 0 {
			action = args[0]
		}

		switch action {
		case "down":
			_ = stack.Down(stackDir)
		case "restart":
			_ = stack.Restart(stackDir, stackComponent)
		default:
			_ = stack.Status(stackDir)
		}
	},
}

var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "Quick shortcut to check running ThingsBoard stack status",
	Run: func(cmd *cobra.Command, args []string) {
		_ = stack.Status(stackDir)
	},
}

func init() {
	stackCmd.Flags().StringVarP(&stackDir, "dir", "d", "/opt/piltismart/tb-stack", "Deployment directory")
	stackCmd.Flags().StringVarP(&stackComponent, "component", "c", "all", "Component to manage (all|tb-db|tb|edge-tb)")

	statusCmd.Flags().StringVarP(&stackDir, "dir", "d", "/opt/piltismart/tb-stack", "Deployment directory")
}
