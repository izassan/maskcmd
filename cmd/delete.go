package cmd

import (
	"fmt"
	"os"
	"slices"

	"github.com/izassan/maskcmd/alias"
	"github.com/spf13/cobra"
)

var deleteCmdArgs struct{
	deleteAlias string
	targetShell string
}


var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete alias",
	Run: func(cmd *cobra.Command, args []string) {
		ads, err := alias.LoadAliasDefinitionData()
		if err != nil{
			fmt.Println(err.Error())
			os.Exit(1)
		}
		for _, ad := range ads.AliasDefinitions{
			if ad.Shell != deleteCmdArgs.targetShell{
				continue
			}
			if !slices.Contains(ad.Aliases, deleteCmdArgs.deleteAlias){
				continue
			}

			index := slices.Index(ad.Aliases, deleteCmdArgs.deleteAlias)

			fmt.Printf("Remove Alias '%s'", deleteCmdArgs.deleteAlias)
			ad.Aliases = slices.Delete(ad.Aliases, index, index+1)
			fmt.Println(ad.Aliases)
		}
		if err := alias.SaveAliasDefinitionData(ads); err != nil{
			fmt.Println(err.Error())
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
	deleteCmd.Flags().StringVarP(&deleteCmdArgs.deleteAlias, "alias", "a", "", "delete target alias")
	deleteCmd.Flags().StringVarP(&deleteCmdArgs.targetShell, "shell", "s", "all", "delete target shell")
}

func removeAlias(){

}
