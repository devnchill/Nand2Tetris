// This file is part of www.nand2tetris.org
// and the book "The Elements of Computing Systems"
// by Nisan and Schocken, MIT Press.
// File name: projects/4/Mult.asm

// Multiplies R0 and R1 and stores the result in R2.
// (R0, R1, R2 refer to RAM[0], RAM[1], and RAM[2], respectively.)
// The algorithm is based on repetitive addition.

  // set value of D to R0
  @R0
  D=M

  // R0 <= 0?
  @END
  D;JLE 

  // set value of D to R1
  @R1
  D=M

  // R0 <= 0?
  @END
  D;JLE 

  // i = 0
  @i
  M=0

  // prod = 0
  @prod 
  M=0

(LOOP)
  // set value of memory register to R0
  @R0 
  D=M

  // add R0 to prod 
  @prod 
  M=D+M

  // increament i
  @i 
  M=M+1 
  // store value of D to i
  D=M

  // find difference between i and R1
  @R1 
  D=D-M

  // isi>=R1 ?
  @END 
  D;JGE

  @LOOP
  0;JMP

(END)
  @prod 
  D=M
  @R2
  M=D
  @END
  0;JMP 
