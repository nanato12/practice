#include <stdio.h>

#define STACK_SIZE 128

int sp;                /*スタックポインタ*/
int stack[STACK_SIZE]; /*数値を格納する配列*/

void init_stack() /*スタックポインタを初期化する*/
{
  sp = -1;
}

void push(int input) /*スタックにデータを一つ入れる*/
{
  if (input >= 0, input < 10)
  {
    if (sp <= STACK_SIZE)
    {
      sp = sp + 1;
      stack[sp] = input;
      printf("%dをpush\n", input);
    }
    else
    {
      printf("スタックがいっぱいです\n");
    }
  }
  else
  {
    printf("数値を入力してください");
  }
}
void pop() /*スタックからデータを一つ取り出す*/
{
  if (sp >= 0)
  {
    printf("%dをpop\n", stack[sp]);
    sp = sp - 1;
  }
  else
  {
    printf("スタックが空です\n");
  }
}
void show_stack() /*スタックに格納された値をすべて格納する*/
{
  printf("[ ");
  int i;
  for (i = 0; i <= sp; i++)
  {
    printf("%d ", stack[i]);
  }
  printf("]");
  printf("\nsp:%d\n\n", sp);
}

unsigned str_length(const char str[]) /*文字列の長さを返す関数*/
{
  unsigned len = 0;
  while (str[len])
    len++;
  return len;
}

void str_judgement(const char str[]) /*配列を判定し、push,pop,showを行う*/
{
  int x;
  int i; /*次の数値をpushする*/
  int o; /*popする*/
  int s; /*showする*/
  int e; /*入力の終わり*/

  int len = str_length(str);
  for (x = 0; x < len; x++)
  {
    printf("str[x]=%d\n", str[x]);
    if (str[x] == i)
    {
      push(str[x + 1]);
    }
    else if (str[x] == o)
    {
      pop();
    }
    else if (str[x] == s)
    {
      show_stack();
    }
    else if (str[x] == e)
    {
      break;
    }
    else if (str[x] >= 0, str[x] < 10)
    {
      if (str[x - 1] == i)
      {
        return;
      }
      else
      {
        printf("%dは数字が連続している,または配列を間違えている場合があります\n", str[x]);
      }
    }
    else
    {
      printf("%dは入力できる文字ではありません\n", str[x]);
    }
  }
}

int main(void)
{
  init_stack();

  char str[128];

  printf("スタック操作内容を入力してください\n");
  scanf("%s", str);

  str_judgement(str);

  return 0;
}
