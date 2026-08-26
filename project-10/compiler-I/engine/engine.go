package engine

type Engine struct {
	inputFilePath  string
	outputFilePath string
}

func NewCompilationEngine(inputFilePath string, outputFilePath string) Engine {
	return Engine{inputFilePath: inputFilePath, outputFilePath: outputFilePath}
}

func (e *Engine) compileClass() {
}
func (e *Engine) compileClassVarDec() {
}
func (e *Engine) compileSubroutine() {
}
func (e *Engine) compileParameterList() {
}
func (e *Engine) compileVarDec() {
}
func (e *Engine) compileStatements() {
}
func (e *Engine) compileDo() {
}
func (e *Engine) compileLet() {
}
func (e *Engine) compileWhile() {
}
func (e *Engine) compileReturn() {
}
func (e *Engine) compileIf() {
}
func (e *Engine) compileExpression() {
}
func (e *Engine) compileTerm() {
}
func (e *Engine) compileExpressionList() {
}
