// This file is part of www.nand2tetris.org
// and the book "The Elements of Computing Systems"
// by Nisan and Schocken, MIT Press.
// File name: projects/04/Fill.asm

// Runs an infinite loop that listens to the keyboard input.
// When a key is pressed (any key), the program blackens the screen,
// i.e. writes "black" in every pixel;
// the screen should remain fully black as long as the key is pressed. 
// When no key is pressed, the program clears the screen, i.e. writes
// "white" in every pixel;
// the screen should remain fully clear as long as no key is pressed.

// Put your code here.

(ROOTLOOP)

//flag means if to blacken the screen
@flag
M=0

//check if the key is pressed
@KBD
D=M
//don't setup black if no key is pressed
@WHITE
D;JEQ

@flag
M=-1

(WHITE)

@i
M=0

//8192 and n mean the number of loops
@8192
D=A
@n
M=D

(FILLLOOP)
//compare i and n
@i
D=M
@n
D=M-D
@ROOTLOOP
D;JEQ

//add i and SCREEN address to get current address that should be filled
@i
D=M
@SCREEN
D=A+D

//fill current address using flag
@curaddress
M=D

@flag
D=M

@curaddress
A=M
M=D

//i++
@i
M=M+1

@FILLLOOP
0;JMP

//infinite loop
(END)
@END
0;JMP