#include <stdio.h>
#include <string.h>

#ifdef _WIN32
#include <Windows.h>
#else
#include <sys/ioctl.h>
#endif

int main(int argc, char *argv[])  {
    if (argc == 1) {
        printf("USAGE:\n");
        printf("    %s <TEXT>...\n", argv[0]);
        return 0;
    }

    int width;

    #ifdef _WIN32
        CONSOLE_SCREEN_BUFFER_INFO csbi;
        GetConsoleScreenBufferInfo(GetStdHandle(STD_OUTPUT_HANDLE), &csbi);
        width = csbi.srWindow.Right - csbi.srWindow.Left + 1;
    #else
        struct winsize ws;
        ioctl(0, TIOCGWINSZ, &ws);
        width = ws.ws_col;
    #endif

    for (int i = 1; i < argc; ++i) {
        printf("%*s\n", (int) ((width + strlen(argv[i])) / 2), argv[i]);
    }

    return 0;
}