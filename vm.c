#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <math.h>

#include "vmisa.h"

/* 
    Virtual machine description

    Memory:
        0x000   .stack  size: 512b 0x200, perms: rw-
        0x200   .text   size: 512b 0x200, perms: r-x
    
    CPU Registers:
        r0-rF   16-bit general purpouse
        sp      16-bit stack pointer
        pc      16-bit program counter
        lr      16-bit link register

    ISA (ARMv6 inspired instruction set):
        vstr
            this is the only instruction capable of storing into memory
            vstr r0, [r1]       store r0 into *r1

        vldr
            this is the only instruction capable of loading from memory
            vldr r0, [r1]       laod r0 from *r1
        
        vmov
            copy immediate value or register content to another register
            vmov r0, r1         copy values from r1 into r0
            vmov r0, #5         copy immediate value #5 into r0

        vadd
            add immediate value or register content with another register
            stores result into 1st operand
            vadd r0, r0, r1     r0=r0+r1
            vadd r0, r1, #5     r0=r1+5

        veor
            perform bitwise exclusive or (XOR) on an immediate value or register content with another register
            stores result into 1st operand
            veor r0, r0, r1     r0=r0^r1
            veor r0, r1, #5     r0=r1^5

        vand
            perform bitwise AND on an immediate value or register content with another register
            stores result into 1st operand
            vand r0, r0, r1     r0=r0&r1
            vand r0, r1, #5     r0=r1&5

        vor
            perform bitwise OR on an immediate value or register content with another register
            stores result into 1st operand
            vor  r0, r0, r1     r0=r0|r1
            vor  r0, r1, #5     r0=r1|5
        
        vb
            branches to another location (updates pc register)
            value can be an offset or register
            vb   r0             branches to *r0
            vb   #5             branches to offset +5
        
        vbl
            branch and link to another location (updates pc register and link register)
            value can be an offset or register
            vbl   r0            branch and link to *r0
            vbl   #5            branch and link to offset +5

        vret
            branch to link register content (updates pc register and empties the link register value)
            vret

*/

void printVMState(uint32_t* mem, uint32_t* gp_regs, uint32_t* pc){
    printf("============================================\n");
    printf("============= Current VM State =============\n");
    printf("============================================\n");
    
    printf(">>>> Registers\n");
    printf("pc: 0x%08x\n", *pc);
    for (int i = 0; i < 2; i++) printf("0x%08x : 0x%08x\n", i, gp_regs[i]);
    printf("============================================\n");
    
    printf(">>>> Memory\n");
    for (int i = 0; i < 7; i++) printf("0x%08x : 0x%08x\n", i, mem[i]);
    printf("============================================\n");

    fflush(stdout);
}

void printInstructionDetails(uint8_t* opcode, uint8_t* first_op, uint8_t* second_op, uint8_t* dst_op){
    printf(">>>> ins:0x%02x op1:0x%02x op2:0x%02x dst:0x%02x\n", *opcode, *first_op, *second_op, *dst_op);
    fflush(stdout);
}

int main(int argc, char const *argv[])
{
    uint32_t mem[7] = {
        0x00000000, // vnop
        0x01000000, // vadd r0, r1, r0
        0x02000000, // vsub r0, r1, r0
        0x03000000, // vand r0, r1, r0
        0x04000000, // vor  r0, r1, r0
        0x05000000, // veor r0, r1, r0
        0x06000000, // vmov r0, r1
    };

    uint32_t gp_regs[0x10] = {
        0xffffffff, // r0
        0x55555555, // r1
        0x00000000, // r2
        0x00000000, // r3
        0x00000000, // r4
        0x00000000, // r5
        0x00000000, // r6
        0x00000000, // r7
        0x00000000, // r8
        0x00000000, // r9
        0x00000000, // rA
        0x00000000, // rB
        0x00000000, // rC
        0x00000000, // rD
        0x00000000, // rE
        0x00000000, // rF
    };

    uint32_t pc = 0x00000000;
    uint32_t sp = 0x00000000;
    uint32_t lr = 0x00000000;

    size_t instructions_q = sizeof(mem)/sizeof(mem[0]);
    printf("Will execute %lu instructions\n", instructions_q);

    printf("Printing Initial VM sate...\n");
    printVMState(mem, gp_regs, &pc);

    uint8_t executionState = 1;
    while(executionState == 1)
    {
        uint32_t opcode = mem[pc];
        uint8_t instruction     = opcode >> 24;
        uint8_t first_operand   = opcode >> 16 & 0x00ff;
        uint8_t second_operand  = opcode >> 8  & 0x0000ff;
        uint8_t destination     = opcode       & 0x000000ff;

        printInstructionDetails(&instruction, &first_operand, &second_operand, &destination);

        switch (instruction)
        {
            case 0x00:
                printf("Performing vnop operation\n");
                vnop();
                break;
           
           case 0x01:
                printf("Performing vadd operation\n");
                vadd(gp_regs[0], gp_regs[1], &gp_regs[destination]);
                break;

            case 0x02:
                printf("Performing vsub operation\n");
                vsub(gp_regs[0], gp_regs[1], &gp_regs[destination]);
                break;

            case 0x03:
                printf("Performing vand operation\n");
                vand(gp_regs[0], gp_regs[1], &gp_regs[destination]);
                break;

            case 0x04:
                printf("Performing vor operation\n");
                vor(gp_regs[0], gp_regs[1], &gp_regs[destination]);
                break;

            case 0x05:
                printf("Performing veor operation\n");
                veor(gp_regs[0], gp_regs[1], &gp_regs[destination]);
                break;

            case 0x06:
                printf("Performing vmov operation\n");
                vmov(&gp_regs[0], &gp_regs[1]);
                break;                                

            default:
                break;
        }
        
        pc++;
        printVMState(mem, gp_regs, &pc);

        if (pc >= instructions_q) executionState = 0;
    }
    
    printf("VM Execution ended!\n");
    return 0;
}
