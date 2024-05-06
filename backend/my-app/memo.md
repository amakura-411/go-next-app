
# 処理の流れのかくにん

1. main.go が実行し、/router/router.go にあるルーティングを読み込む
2. routingを行った後、entityに



## memo

- :=は型推論
- domainディレクトリとは
  - domainディレクトリは、クリーンアーキテクチャにおける「ドメイン層」に該当します。
  - このディレクトリには、ビジネスドメインに関するロジックが含まれます。
  - 具体的には、アプリケーションのビジネスルールやエンティティ（domain model）が定義されます。
  - わかりやすい参考資料：[Goでエンティティを実装（「ドメイン駆動設計入門」Chapter3）](https://zenn.dev/msksgm/articles/20220225-go-itddd-03-entity)

- time.Timeはgoの日付型
- passwordの作成
  - ハッシュ化するのに参考になった記事:[【Go】bcrypt を使ってパスワードのハッシュ値を生成して検証する](https://zenn.dev/kou_pg_0131/articles/go-digest-and-compare-by-bcrypt)

