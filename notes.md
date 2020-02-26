# Fibonacci

### Features
- Variables
- Stack
- Aritmetic ops
- Routines
- Recursion

### Sample C code
```
int f(int n) {
    if (n == 0 || n == 1) {
        return n;
    }
    return (f(n-1) + f(n-2));
}
```

### Rough intel x86 equivalent
```
str
mov  rbx, 5     # or just "push 5"
push rbx        #
mov  sp, pc # call _rutine
hlt

_rutine:
    pop rbx
    cmp rbx, 0
    je  _rutine_end
    cmp rbx, 1
    je  _rutine_end

    dec rbx
    push rbx
    call _rutine    # rax = f(n-1)
    
    pop rbx
    dec rbx
    push rax
    push rbx
    call _rutine    # rax = f(n-2)
    pop rbx
    pop rbx
    add rax, rbx

    _rutine_end:
        ret         # analize how to do this
```

## Architecture and mnemonics needed

Registers:
    3 8bit general purpose registers: ra, rb, rc
    1 8bit stack pointer: sp
    1 8bit program counter: pc
    1 8bit condition register: cr (like flags but more high level)

Encoding:
    operation
     0 0 0 0   0 0 0 0     0 0 0 0   0 0 0 0    0 0 0 0   0 0 0 0
    |-------| |-------|   |-----------------|  |-----------------|
     imm val? operation     reg val1 / zero    reg val2 || imm val

    imm val? == either b0000 or b1111 (0x0 or 0xf)
    operation ranges from 0x0 to 0x7 right now as we don't have more ops defined
    reg val1 can be 0x00 (ra), 0x01 (rb), 0x02 (rc), 0x88 (sp), 0xff (pc)
    imm val on the last byte can range from 0x00 to 0xff

Mnemonics:

    start:  Reset all register values to 0
            start               00 00 00

    end:    Stop/finish execution
            end                 ff 00 00

    set:    Copy immediate or register value into destination
            set  ra, rb         01 00 01
            set  ra, 0xff       f1 00 ff
    
    put:    Saves the value into the stack
            put ra              02 00 00
            put rb              02 01 00
            put rc              02 02 00
            put 0xff            f2 00 ff
    
    get:    Retreives the value from the stack into the register
            get ra              03 00 00
            get rb              03 01 00
            get rc              03 02 00

    cmp:    Compares values
            cmp ra, 0xff        f4 00 ff
            cmp ra, rb          04 01 00

    if:     Execute following instruction if the condition is met
            if  0xff            f5 00 ff
            if  ra              05 00 00

    dec:    Decrement register vale
            dec ra              06 00 00
            dec ra, 1           f6 00 01
            dec rb              06 01 00
    
    add:    Add values
            add ra, rb          07 00 01
            add rb, ra          07 01 00
            add rb, rc          07 01 02
            add ra, 0xff        f7 00 ff
            add rb, 0xff        f7 01 ff
    
    call:   
            put pc              02 00 ff
            put sp              02 00 88
            set pc, 0xff        f1 ff ff

    ret:
            get sp              03 00 88
            get pc              03 00 ff
