
//push constant 7
@7D=A
@SP 
A=M 
M=D
@SP 
M=M+1

//push constant 8
@8D=A
@SP 
A=M 
M=D
@SP 
M=M+1

//add
@SP
AM=M-1
D=M
A=A-1
M=M+D
