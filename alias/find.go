package alias

import "errors"

var NoRecord = errors.New("no record")

func FindByCommand(maskCmdData *MaskCmdData, searchCommand string)(*AliasDefinition, error){
	for _, ad := range maskCmdData.AliasDefinitions{
		if ad.Command == searchCommand{
			return ad, nil
		}
	}
	return nil, NoRecord
}
