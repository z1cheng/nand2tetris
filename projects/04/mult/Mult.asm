// This file is part of www.nand2tetris.org
// and the book "The Elements of Computing Systems"
// by Nisan and Schocken, MIT Press.
// File name: projects/04/Mult.asm

// Multiplies R0 and R1 and stores the result in R2.
// (R0, R1, R2 refer to RAM[0], RAM[1], and RAM[2], respectively.)
//
// This program only needs to handle arguments that satisfy
// R0 >= 0, R1 >= 0, and R0*R1 < 32768.

// Put your code here.

//Will add R1 R0 times
@i
M=0

@R2
M=0

(LOOP)

//R0 means the number of loops
//compare i and R0
@i
D=M

@R0
D=M-D;

//jump to end if i equals R0
@END
D;JEQ

//add R1 and R2, put the result into R2
@R1
D=M
@R2
M=M+D

//i++
@i
M=M+1

@LOOP
0;JMP

//infinite loop
(END)
@END
0;JMP