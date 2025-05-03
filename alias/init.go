package alias

import (
	"encoding/json"
	"os"
	"path/filepath"
)

func InitAliasDefinitionData() error{
	dataPath := GetAliasDefinitionDataPath()
	if IsExistDataFile(dataPath){
		return nil
	}

	if err := os.MkdirAll(filepath.Dir(dataPath), 0755); err != nil{
		return err
	}

	initData := &MaskCmdData{
		AliasDefinitions: []*AliasDefinition{},
	}
	output, err := json.MarshalIndent(&initData, "", "    ")
	if err != nil{
		return err
	}

	if err := os.WriteFile(dataPath, output, 0755); err != nil{
		return err
	}

	return nil
}

func IsExistDataFile(dataPath string) bool{
	_, err := os.Stat(dataPath)
	return err == nil
}
