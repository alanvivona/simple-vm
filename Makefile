compile:
	clang vm.c -o ./out/vm && ./out//vm

clean:
	rm ./out/*

dep:
	sudo apt update && and sudo apt -y install llvm build-essential clang