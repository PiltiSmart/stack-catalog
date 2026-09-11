package cmd

import (
	"fmt"
	"runtime"

	"github.com/piltismart/tb-stack-ps-cli/pkg/ui"
	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Show ps CLI version and build info",
	Run: func(cmd *cobra.Command, args []string) {
		ui.PrintBanner("ThingsBoard PS CLI Version")
		fmt.Printf("  - CLI Version   : %s1.0.0 (Phase 1)%s\n", ui.ColorGreen, ui.ColorReset)
		fmt.Printf("  - Framework     : Cobra v1.8.1\n")
		fmt.Printf("  - Go Runtime    : %s\n", runtime.Version())
		fmt.Printf("  - OS / Arch     : %s / %s\n\n", runtime.GOOS, runtime.GOARCH)
	},
}
