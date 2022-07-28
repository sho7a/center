.PHONY: all clean

ifeq ($(OS), Windows_NT)
all: clean bin\center.exe
else
all: clean bin/center
endif

bin\center.exe:
	if not exist bin ( mkdir bin )
	gcc src\main.c -o bin\center.exe

bin/center:
	if [ ! -d bin ]; then mkdir bin; fi
	gcc src/main.c -o bin/center

clean:
ifeq ($(OS), Windows_NT)
		if exists bin ( rmdir /s /q bin )
else
		if [ -d bin ]; then rm -r bin; fi
endif