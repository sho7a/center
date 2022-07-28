#include <stdio.h>

#ifdef _WIN32
#include <Windows.h>
#endif

int main(int argc, char *argv[])  {
    #ifdef _WIN32
        CONSOLE_SCREEN_BUFFER_INFO csbi;
        GetConsoleScreenBufferInfo(GetStdHandle(STD_OUTPUT_HANDLE), &csbi);
        int width = csbi.srWindow.Right - csbi.srWindow.Left + 1;

        for (int i = 1; i < argc; ++i) {
            printf("%s\n", argv[i]);
        }
    #endif
	return 0;
}