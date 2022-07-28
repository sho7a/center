all: center

center:
	mkdir bin
	gcc src/main.c -o bin/center

center.exe:
	if not exist bin ( mkdir bin )
	gcc src\main.c -o bin\center.exe