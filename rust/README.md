# rust-practice
rustのお勉強

## docker
```bash
$ docker-compose build
$ docker-compose up -d
$ docker exec -it rust1.47.0 bash

nanato12@rust1.47.0:~/workspace$ rustc --version
rustc 1.47.0 (18bf6b4f0 2020-10-07)
nanato12@rust1.47.0:~/workspace$ cargo --version
cargo 1.47.0 (f3c7e066a 2020-08-28)
nanato12@rust1.47.0:~/workspace$ rustup --version
rustup 1.22.1 (b01adbbc3 2020-07-08)
```

## hello
```bash
nanato12@rust1.47.0:~/workspace$ cd hello/
nanato12@rust1.47.0:~/workspace/hello$ cargo run
   Compiling hello v0.1.0 (/home/nanato12/workspace/hello)
    Finished dev [unoptimized + debuginfo] target(s) in 4.86s
     Running `target/debug/hello`
Hello, world!
```

## cargo
Rustビルドツールおよびパッケージマネージャ
- プロジェクトのビルドには`cargo build`
- プロジェクトの実行には`cargo run`
- プロジェクトのテストには`cargo test`
- プロジェクトのドキュメントのビルドには`cargo doc`
- ライブラリをcrates.ioに公開するには`cargo publish`

## クレート [crate]
Rustのパッケージレジストリであるcrates.io上で、あらゆる種類のライブラリを見つけることができます。Rustではパッケージのことをよく「クレート」と呼びます。

## 共存関係の追加
Cargo.tomlの`[dependencies]`に記載する。
`cargo build`で依存関係のインストール

## 参考
- https://www.rust-lang.org/ja/learn/get-started
