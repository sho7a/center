#include <stdio.h>

#ifdef _WIN32
#include <Windows.h>
#endif

void printc(char *text, int width) {
    printf("%*s", -width, printf("%*s", (width + strlen(text)) / 2, text));
}

int main(int argc, char *argv[])  {
    #ifdef _WIN32
        CONSOLE_SCREEN_BUFFER_INFO csbi;
        GetConsoleScreenBufferInfo(GetStdHandle(STD_OUTPUT_HANDLE), &csbi);
        int width = csbi.srWindow.Right - csbi.srWindow.Left + 1;
        for (int i = 1; i < argc; ++i) {
            printc(argv[i], width);
        }
    #endif
    return 0;
}