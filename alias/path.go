package alias

import (
	"fmt"
	"os"
	"path/filepath"
)

func GetAliasDefinitionDataPath() string{
	dataPath, result := os.LookupEnv("MASKCMD_DATA_PATH")
	if !result{
		home, err := os.UserHomeDir()
		if err != nil{
			fmt.Println(err.Error())
			os.Exit(1)
		}
		dataPath = filepath.Join(home, ".config", "maskcmd", "data.json")
	}
	dbAbsPath, err := filepath.Abs(dataPath)
	if err != nil{
		fmt.Println(err.Error())
		os.Exit(1)
	}
	return dbAbsPath
}
