package cmd

import (
	"errors"
	"fmt"
	"os"
	"slices"

	"github.com/izassan/maskcmd/alias"
	"github.com/spf13/cobra"
)

var addCmdArgs struct{
	alias string
	command string
	shell string
}


var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add alias",
	Run: func(cmd *cobra.Command, args []string) {
		ads, err := alias.LoadAliasDefinitionData()
		if err != nil{
			fmt.Println(err.Error())
			os.Exit(1)
		}
		record, err := alias.FindByCommand(ads, addCmdArgs.command)
		if err != nil{
			if err == alias.NoRecord{
				ad := alias.NewAliasDefinition(
					addCmdArgs.command,
					[]string{},
					addCmdArgs.shell,
				)
				ads.AliasDefinitions = append(ads.AliasDefinitions, ad)
				record = ad
			}else{
				fmt.Println(err.Error())
				os.Exit(1)
			}
		}

		if record.Shell != addCmdArgs.shell{
			fmt.Println(errors.New("shell validate error").Error())
			os.Exit(1)
		}

		if slices.Contains(record.Aliases, addCmdArgs.alias){
			fmt.Println(errors.New("duplicate alias").Error())
			os.Exit(1)

		}

		record.Aliases = append(record.Aliases, addCmdArgs.alias)
		if err := alias.SaveAliasDefinitionData(ads); err != nil{
			fmt.Println(err.Error())
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
	addCmd.Flags().StringVarP(&addCmdArgs.command, "command", "c", "", "execute actual command")
	addCmd.Flags().StringVarP(&addCmdArgs.alias, "alias", "a", "", "set alias")
	addCmd.Flags().StringVarP(&addCmdArgs.shell, "shell", "s", "all", "output shell")
}
