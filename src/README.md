# プログラミング
読み込んだファイルの各行を並列に処理して、処理結果を元の行の並び通りに別のファイルに出力するプログラムを作ってください。処理の内容は行データのSHA256チェックサムのHEXダンプとします。

採点ポイントは以下です。

並列処理が正しく実装できていること
チェックサムの処理は抽象に依存し、「ハッシュアルゴリズムの追加・変更や出力先の追加・変更」に対して、解放閉鎖の原則を満たすこと。
メソッドの呼び出し部分と実装部分を切り離したユニットテストができること
domainのもつメソッドは外部に依存せず、変更されにくく、単体でテストが行えること。
変数や関数の命名が明瞭で、適切なコメントアウトを含み、マジックナンバーや黒魔術を含まないなど、可読性の高いプログラムであること。
使用可能言語：
Go, Java, Kotlin, Rust

# 動作手順

```
$ go build -o hashdump

$ ./hashdump
  or
$ ./hashdump -o ./out.txt
```