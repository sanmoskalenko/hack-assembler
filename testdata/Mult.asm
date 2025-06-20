// This file is part of www.nand2tetris.org
// and the book "The Elements of Computing Systems"
// by Nisan and Schocken, MIT Press.
// File name: projects/4/Mult.asm

// Multiplies R0 and R1 and stores the result in R2.
// (R0, R1, R2 refer to RAM[0], RAM[1], and RAM[2], respectively.)
// The algorithm is based on repetitive addition.

  @R2
  M=0               // Initialize result to 0

  @R1
  D=M               // Load R1 (first operand) into D register
  @END
  D;JEQ             // If R1 = 0, jump to END (optimization: 0 * anything = 0)
  @multiplier
  M=D               // Store R1 as multiplier

  @R0
  D=M               // Load R0 (second operand) into D register
  @END
  D;JEQ             // If R0 = 0, jump to END (optimization: anything * 0 = 0)
  @counter
  M=D               // Store R0 as counter (how many times to add)

  @IS_NEGATIVE
  D;JLE             // If counter <= 0, handle negative numbers

(STEP_0)
  @counter
  M=D-1             // Decrement counter by 1
  D=M               // Load decremented counter into D

(LOOP)
    @multiplier
    D=M            // Load multiplier into D register

    @R2
    M=D+M          // Add multiplier to current result (R2 = R2 + multiplier)  

    @counter
    M=M-1         // Decrement loop counter
    D=M           // Load counter value into D

    @LOOP
    D;JGE         // If counter >= 0, continue loop

(END)
    @END
    0;JMP        // Infinite loop to halt program

(IS_NEGATIVE)
  @multiplier
  M=-M          // Make multiplier positive (negate it)

  @counter
  M=-M          // Make counter positive (negate it)
  D=M           // Load positive counter into D

  @STEP_0
  D;JGT         // If counter > 0, proceed to multiplication loop