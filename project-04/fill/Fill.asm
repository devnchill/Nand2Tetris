// Runs an infinite loop that listens to the keyboard input. 
// When a key is pressed (any key), the program blackens the screen,
// i.e. writes "black" in every pixel. When no key is pressed, 
// the screen should be cleared.

(LOOP)
  @LOOP
  // keypressed is represented by word of 16 bit length at KBD/RAM[24576] and screen is represented at RAM[16384]/SCREEN
  @KBD
  D=M

  @CLEAR
  D;JEQ

  @BLACK
  0;JMP

(BLACK)

  @i
  M=0

  @8192
  D=A

  @totalwords
  M=D

  @16384
  D=A 

  @ptr
  M=D

(BLACK_LOOP)
  // Dereference pointer to access screen memory
  @ptr
  A=M
  M=-1       

  @ptr
  M=M+1

  @i
  M=M+1

  @i
  D=M
  @totalwords
  D=D-M
  @BLACK_LOOP
  D;JLT

  @LOOP
  0;JMP

(CLEAR)

  @i
  M=0

  @8192
  D=A
  @totalwords
  M=D

  @SCREEN
  D=A
  @ptr
  M=D

(CLEAR_LOOP)
  @ptr
  A=M
  M=0       // write all 16 bits = 0 → clear

  @ptr
  M=M+1

  @i
  M=M+1

  @i
  D=M
  @totalwords
  D=D-M
  @CLEAR_LOOP
  D;JLT

  @LOOP
  0;JMP


