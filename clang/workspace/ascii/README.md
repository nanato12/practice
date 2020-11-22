# ASCIIコード表

練習用

char型は`-127~128`の数字を格納できる。  

`%x` で16進数表示  
`%d` で10進数表示  
`%c` で文字型として表示  

```c
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
```
