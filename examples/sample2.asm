start
set ra 0x01
set rb ra
set rc 0x02
set rd rc
put ra
put 0x44
get ra
get rd
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

# prints "it works!"
put 0x49
put 0x54
put 0x20
put 0x57
put 0x4f
put 0x52
put 0x4b
put 0x53
put 0x21
# go to next memory "line"
set sp 0x10

# prints "bye"
put 0x42
put 0x59
put 0x45
get ra
get ra
get ra

put 0x00
put 0x42
put 0x59
put 0x45
get ra
get ra
get ra
get ra

put 0x00
put 0x00
put 0x42
put 0x59
put 0x45
put 0x21

end