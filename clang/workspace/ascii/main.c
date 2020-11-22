#include <stdio.h>

void main()
{
    int x, y;
    char c;

    for (y = 0; y < 16; y++)
    {
        for (x = 2; x < 8; x++)
        {
            c = x * 16 + y;
            printf("%2X:%3d: %c | ", c, c, c);
        }
        printf("\n");
    }
}
