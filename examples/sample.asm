## spaces and comments
set ra, rb   
set ra, rb		
#set ra, rb
#set ra, rb	
#set ra, rb 
set	 ra	 ,	 rb	 #	 some comment
SeT	 rA	 ,	 Rb	 #	 some comment
SeT	 rA	 	 Rb	 #	 some comment
SeT,	 rA,	 	 Rb,	 #	 some comment

## valid numeric formats
set ra 55
set ra 055 // octal
set ra 111b
set ra 111B
set ra 55h
set ra 55H
set ra 0x5a
set ra 0X5A
set ra 0x5A

## invalid numeric formats
set ra 0xx5A
set ra 123b
set ra 0xzb
set ra 0
set ra 02
set ra 55hh
set ra h
set ra 0x
set ra b

## extra or unkown mnemonics and registers
set set ra 0x5A
set ra rb rc rd
garbage
set ra, garbage

## regular code sample (testing all the mnemonics)
start
set ra 0x01
set rb ra
set rc 0x02
set rd rc
put ra
put 0x44
get ra
add rd 0x44
add rd ra
sub rd 0x44
sub rd ra
dec ra
inc ra
not ra
neg ra
and ra rb
and ra 0x55
or  ra rb
or  ra 0x55
xor ra rb
xor ra 0x55
end