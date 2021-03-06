# Cとそれに対応するアセンブラ
簡単な例

```c
int main() {
  return 42;
}
```

```bash
$ cc -o test1 test1.c
$ ./test1
$ echo $?
42
```
ここでは正しく42が返されていることがわかる。


Cプログラムに対応するアセンブリプログラム
```s
.intel_syntax noprefix
.globl main
main:
        mov rax, 42
        ret
```
グローバルなラベルmainが定義されていて、ラベルのあとにmain関数のコードが続いている。

ここでは42という値を、RAXというレジスタにセットし、mainからリターンしている。

整数を入れられるレジスタはRAXを含めて合計で16個あるが、関数からリターンしたときにRAXに入っている値が関数の返り値という約束になっているので、ここでは値をRAXにセットしている。

```bash
$ cc -o test2 test2.s
$ ./test2
$ echo $?
42
```

大雑把にいうと、Cコンパイラは、test1.cのようなCコードを読み込んだ時に、test2.sのようなアセンブリを出力するプログラムということになる。
