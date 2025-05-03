package alias

import (
	"encoding/json"
	"os"
)

func LoadAliasDefinitionData() (*MaskCmdData, error){
	dataPath := GetAliasDefinitionDataPath()
	dataByte, err := os.ReadFile(dataPath)
	if err != nil{
		return nil, err
	}

	var maskCmdData *MaskCmdData
	if err := json.Unmarshal(dataByte, &maskCmdData); err != nil{
		return nil, err
	}
	return maskCmdData, nil
}

func SaveAliasDefinitionData(maskCmdData *MaskCmdData) error{
	dataPath := GetAliasDefinitionDataPath()

	output, err := json.MarshalIndent(&maskCmdData, "", "    ")
	if err != nil{
		return err
	}

	if err := os.WriteFile(dataPath, output, 0755); err != nil{
		return err
	}

	return nil
}
