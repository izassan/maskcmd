package main

import (
	"fmt"
	"os"

	"github.com/izassan/maskcmd/alias"
	"github.com/izassan/maskcmd/cmd"
)

func main() {
	if err := alias.InitAliasDefinitionData(); err != nil{
		fmt.Println(err.Error())
		os.Exit(1)
	}
	cmd.Execute()
}

