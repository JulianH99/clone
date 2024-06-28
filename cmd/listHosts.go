package cmd

import (
	"fmt"

	"github.com/JulianH99/clone/internal"
	"github.com/JulianH99/clone/internal/ui"
	"github.com/charmbracelet/bubbles/table"
	"github.com/spf13/cobra"
)

// listHostsCmd represents the list hosts command
var listHostsCmd = &cobra.Command{
	Use:   "list",
	Short: "List available hosts",
	RunE: func(cmd *cobra.Command, args []string) error {
		hosts, err := internal.SshHosts()
		if err != nil {
			return err
		}

		columns := []table.Column{
			{Title: "Host", Width: 50},
		}

		rows := []table.Row{}

		for _, host := range hosts {
			rows = append(rows, table.Row{string(host)})
		}

		t := table.New(
			table.WithColumns(columns),
			table.WithRows(rows),
			table.WithHeight(len(hosts)),
			table.WithFocused(false),
		)

		t.SetStyles(ui.TableStyles())

		fmt.Print(ui.InContainer(t.View()))

		return nil
	},
}

func init() {
	hostsCmd.AddCommand(listHostsCmd)
}
