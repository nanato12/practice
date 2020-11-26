#include <stdio.h>

int ctoi(char c)
{
    return c - '0';
}

void main()
{
    printf("そのまま >>> %d\n", '5');
    printf("ctoi >>> %d\n", ctoi('5'));
}
