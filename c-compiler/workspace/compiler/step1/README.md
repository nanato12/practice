# コンパイラ本体の作成

```c
#include <stdio.h>
#include <stdlib.h>

int main(int argc, char **argv) {
  if (argc != 2) {
    fprintf(stderr, "引数の個数が正しくありません\n");
    return 1;
  }

  printf(".intel_syntax noprefix\n");
  printf(".globl main\n");
  printf("main:\n");
  printf("  mov rax, %d\n", atoi(argv[1]));
  printf("  ret\n");
  return 0;
}
```

動作確認
```bash
$ cc -o 9cc 9cc.c
$ ./9cc 123 > tmp.s
$ cat tmp.s
.intel_syntax noprefix
.globl main
main:
  mov rax, 123
  ret
```

Unixにおいてはcc（あるいはgcc）は、CやC++だけではなく多くの言語のフロントエンドということになっていて、与えられたファイルの拡張子で言語を判定してコンパイラやアセンブラを起動するということになっています。

したがってここでは9ccをコンパイルしたときと同じように、.sという拡張子のアセンブラファイルをccに渡すと、アセンブルをすることができます。

```
$ cc -o tmp tmp.s
$ ./tmp
$ echo $?
123
```

## テストファイルの作成

`chmod a+x test.sh`

```bash
#!/bin/bash
assert() {
  expected="$1"
  input="$2"

  ./9cc "$input" > tmp.s
  cc -o tmp tmp.s
  ./tmp
  actual="$?"

  if [ "$actual" = "$expected" ]; then
    echo "$input => $actual"
  else
    echo "$input => $expected expected, but got $actual"
    exit 1
  fi
}

assert 0 0
assert 42 42

echo OK
```

```bash
$ ./test.sh
0 => 0
42 => 42
OK
```

本書を通して読者のみなさんは9ccを何百回、あるいは何千回もビルドすることになるでしょう。9ccの実行ファイルを作成して、その後にテストスクリプトを走らせる作業は毎回同じなので、ツールに任せると便利です。こうした用途で標準的に使われているのがmakeコマンドです。

makeは、実行されるとカレントディレクトリのMakefileという名前のファイルを読み込んで、そこに書かれているコマンドを実行します。Makefileは、コロンで終わるルールと、そのルールのためのコマンドの列という構成になっています。次のMakefileはこのステップで実行したいコマンドを自動化するためのものです。

```makefile
CFLAGS=-std=c11 -g -static

9cc: 9cc.c

test: 9cc
        ./test.sh

clean:
        rm -f 9cc *.o *~ tmp*

.PHONY: test clean
```
