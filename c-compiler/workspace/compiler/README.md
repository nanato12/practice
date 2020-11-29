# Cコンパイラ作成
Cコンパイラ作成の最初のステップとして、四則演算やそのほかの算術演算子をサポートして、次のような式をコンパイルできるようにします。

```
30 + (4 - 2) * -5
```

**Intel記法とAT&T記法**
```s
mov rbp, rsp   // Intel
mov %rsp, %rbp // AT&T

mov rax, 8     // Intel
mov $8, %rax   // AT&T

mov [rbp + rcx * 4 - 8], rax // Intel
mov %rax, -8(rbp, rcx, 4)    // AT&T
```

## Step1
整数1個をコンパイルする言語の作成

## Step2
加減算のできるコンパイラの作成
