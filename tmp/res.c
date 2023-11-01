#include <stdio.h>
#include <sys/ioctl.h>
#include <unistd.h>

int main() {
    struct winsize size;

    if (ioctl(STDOUT_FILENO, TIOCGWINSZ, &size) == 0) {
        int columns = size.ws_col;
        int rows = size.ws_row;
        printf("Terminal size: %d columns x %d rows\n", columns, rows);
    } else {
        perror("ioctl");
        return 1;
    }

    return 0;
}

