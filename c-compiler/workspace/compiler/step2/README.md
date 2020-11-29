# 加減算のできるコンパイラの作成

5+20-4のような式は、コンパイルするときに計算して、その結果の数（この場合21）をアセンブリに埋め込むこともできますが、それだとコンパイラではなくインタープリタのようになってしまうので、加減算を実行時に行うアセンブリを出力する必要があります。

加算と減算を行うアセンブリ命令はaddとsubです。

addは、2つのレジスタを受け取って、その内容を加算し、結果を第1引数のレジスタに書き込みます。

subはaddと同じですが、減算を行います。

これらの命令を使うと、5+20-4は次のようにコンパイルすることができます。

```s
.intel_syntax noprefix
.globl main

main:
        mov rax, 5
        add rax, 20
        sub rax, 4
        ret
```

上記のアセンブリでは、movでRAXに5をセットし、そのあとRAXに20を足して、そして4を引いています。

retが実行される時点でのRAXの値は5+20-4すなわち21になるはずです。


動作確認
```bash
$ cc -o tmp tmp.s
$ ./tmp
$ echo $?
21
```

この加減算のある式を「言語」として考えてみると、この言語は次のように定義することができます。

最初に数字が1つある
そのあとに0個以上の「項」が続いている
項というのは、+の後に数字が来ているものか、-の後に数字が来ているものである
この定義を素直にCのコードに落としてみると、次のようなプログラムになります。

```c
#include <stdio.h>
#include <stdlib.h>

int main(int argc, char **argv) {
  if (argc != 2) {
    fprintf(stderr, "引数の個数が正しくありません\n");
    return 1;
  }

  char *p = argv[1];

  printf(".intel_syntax noprefix\n");
  printf(".globl main\n");
  printf("main:\n");
  printf("  mov rax, %ld\n", strtol(p, &p, 10));

  while (*p) {
    if (*p == '+') {
      p++;
      printf("  add rax, %ld\n", strtol(p, &p, 10));
      continue;
    }

    if (*p == '-') {
      p++;
      printf("  sub rax, %ld\n", strtol(p, &p, 10));
      continue;
    }

    fprintf(stderr, "予期しない文字です: '%c'\n", *p);
    return 1;
  }

  printf("  ret\n");
  return 0;
}
```

makeでのコンパイラ作成

```bash
$ make
$ ./9cc '5+20-4'
.intel_syntax noprefix
.globl main
main:
  mov rax, 5
  add rax, 20
  sub rax, 4
  ret
```
