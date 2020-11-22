# チュートリアル

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


## hello
- [hello](./hello)

## crate
- [crate](./crate-example)
