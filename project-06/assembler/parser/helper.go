package parser

func GetCommandTypeInString(command TCommandType) string {
	if command == 0 {
		return "ACommand"
	}
	if command == 1 {
		return "CCommand"
	}
	return "LCommand"
}
