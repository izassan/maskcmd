package alias

type AliasDefinition struct {
	Command string `json:"command"`
	Aliases []string `json:"aliases"`
	Shell string `json:"shell"`
}

type MaskCmdData struct{
	AliasDefinitions []*AliasDefinition `json:"alias_definitions"`
}


func NewAliasDefinition(cmd string, aliases []string, shell string)(*AliasDefinition){
	return &AliasDefinition{
		Command: cmd,
		Aliases: aliases,
		Shell: shell,
	}
}
