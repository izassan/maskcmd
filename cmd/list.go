package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/izassan/maskcmd/alias"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
)

var listCmdArgs struct{
	filterCommand string
	filterShell string
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List Alias",
	Run: func(cmd *cobra.Command, args []string) {
		ads, err := alias.LoadAliasDefinitionData()
		if err != nil{
			fmt.Println(err.Error())
			os.Exit(1)
		}

		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader([]string{"Aliases", "Command", "Shells"})
		for _, ad := range ads.AliasDefinitions{
			if listCmdArgs.filterCommand != "" && !strings.HasPrefix(ad.Command, listCmdArgs.filterCommand) {
				continue
			}
			if listCmdArgs.filterShell != "" && listCmdArgs.filterShell != ad.Shell{
				continue
			}
			if len(ad.Aliases) > 0{
				table.Append(newTableRow(ad))
			}
		}
		table.Render()
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
	listCmd.Flags().StringVarP(&listCmdArgs.filterCommand, "filter-command", "c", "", "base command filter")
	listCmd.Flags().StringVarP(&listCmdArgs.filterShell, "filter-shell", "s", "", "shell filter")
}


func newTableRow(ad *alias.AliasDefinition) []string{
	printAliases := strings.Join(ad.Aliases, ",")
	return []string{printAliases, ad.Command, ad.Shell}
}
