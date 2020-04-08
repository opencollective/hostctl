package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/guumaster/hostctl/pkg/host"
)

// makeListStatusCmd represents the list enabled command
var makeListStatusCmd = func(cmd string) *cobra.Command {
	status := ""
	switch cmd {
	case "enabled":
		status = "on"
	case "disabled":
		status = "off"
	}
	return &cobra.Command{
		Use:     cmd,
		Aliases: []string{status},
		Short:   fmt.Sprintf("Shows list of %s profiles on your hosts file.", cmd),
		Long: fmt.Sprintf(`
Shows a detailed list of %s profiles on your hosts file with name, ip and host name.
`, cmd),
		RunE: func(cmd *cobra.Command, args []string) error {
			src, _ := cmd.Flags().GetString("host-file")
			raw, _ := cmd.Flags().GetBool("raw")
			cols, _ := cmd.Flags().GetStringSlice("column")

			err := host.ListProfiles(src, &host.ListOptions{
				RawTable:     raw,
				Columns:      cols,
				StatusFilter: status,
			})

			return err
		},
	}
}
