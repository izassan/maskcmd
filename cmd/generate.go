package cmd

import (
	"fmt"
	"os"
	"slices"
	"strings"

	"github.com/google/uuid"
	"github.com/izassan/maskcmd/alias"
	"github.com/spf13/cobra"
)

var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "generate alias shellscript",
}

var generatePowerShellCmd = &cobra.Command{
	Use:   "powershell",
	Short: "generate alias powershell shellscript",
	Run: func(cmd *cobra.Command, args []string) {
		ads, err := alias.LoadAliasDefinitionData()
		if err != nil{
			fmt.Println(err.Error())
			os.Exit(1)
		}

		enableShellStr := []string{
			"all", "pwsh",
		}
		for _, ad := range ads.AliasDefinitions{
			if slices.Contains(enableShellStr, ad.Shell){
				newPowerShellAlias(ad)
			}
		}
	},
}

var generateBashCmd = &cobra.Command{
	Use:   "bash",
	Short: "generate alias bash shellscript",
	Run: func(cmd *cobra.Command, args []string) {
		ads, err := alias.LoadAliasDefinitionData()
		if err != nil{
			fmt.Println(err.Error())
			os.Exit(1)
		}

		enableShellStr := []string{
			"all", "sh", "bash",
		}
		for _, ad := range ads.AliasDefinitions{
			if slices.Contains(enableShellStr, ad.Shell){
				newLinuxAlias(ad)
			}
		}
	},
}

var generateZshCmd = &cobra.Command{
	Use:   "zsh",
	Short: "generate alias zsh shellscript",
	Run: func(cmd *cobra.Command, args []string) {
		ads, err := alias.LoadAliasDefinitionData()
		if err != nil{
			fmt.Println(err.Error())
			os.Exit(1)
		}

		enableShellStr := []string{
			"all", "sh", "zsh",
		}
		for _, ad := range ads.AliasDefinitions{
			if slices.Contains(enableShellStr, ad.Shell){
				newLinuxAlias(ad)
			}
		}
	},
}


func init() {
	rootCmd.AddCommand(generateCmd)
	generateCmd.AddCommand(generatePowerShellCmd)
	generateCmd.AddCommand(generateBashCmd)
	generateCmd.AddCommand(generateZshCmd)
}

func newLinuxAlias(ad *alias.AliasDefinition){
	for _, alias := range ad.Aliases{
		fmt.Printf("alias %s=\"%s\"\n", alias, ad.Command)
	}
}

func newPowerShellAlias(ad *alias.AliasDefinition){
	funcName, err := uuid.NewRandom()
	if err != nil{
		fmt.Println(err.Error())
		os.Exit(1)
	}
	for _, alias := range ad.Aliases{
		if strings.Contains(ad.Command, " "){
			fmt.Printf(
				"function %s(){ %s $Args }\nSet-Alias %s %s\n",
				funcName, ad.Command, alias, funcName,
			)
		}else{
			fmt.Printf(
				"Set-Alias %s %s\n",
				alias, ad.Command,
			)

		}
	}
}
