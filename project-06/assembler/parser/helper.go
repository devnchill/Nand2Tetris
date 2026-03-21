package parser

func getCommandTypeInString(command TCommandType) string {
	if command == 0 {
		return "ACommand"
	}
	if command == 1 {
		return "CCommand"
	}
	return "LCommand"
}
