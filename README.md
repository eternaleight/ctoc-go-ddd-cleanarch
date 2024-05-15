# go-pack

### Project structure

```
.
├── README.md                      // プロジェクトの基本情報、セットアップ手順、使用方法などを記載
├── api                            // API関連の主要なコードやツールを格納するディレクトリ
│   ├── handlers                   // Webリクエストを処理するためのディレクトリ
│   │   ├── auth.go                // 認証処理（ユーザー登録、ログイン、ログアウトなど）
│   │   ├── posts.go               // 投稿に関する処理
│   │   ├── user.go                // ユーザー情報に関する処理
│   │   ├── products.go            // 商品情報に関する処理
│   │   └── purchase.go            // 購入情報に関する処理
│   ├── middlewares                // リクエストやレスポンスの前後で実行される処理を格納するディレクトリ
│   │   └── isAuthenticated.go     // 認証状態の確認と処理を行うミドルウェア
│   └── responses                  // APIからの応答を生成するヘルパー関数を格納するディレクトリ
│       ├── error.go               // エラー応答の生成を行うファイル
│       └── success.go             // 成功応答の生成を行うファイル
├── go.mod                         // プロジェクトの依存関係やモジュール情報を定義するファイル
├── go.sum                         // 依存関係の確認用のチェックサムデータを含むファイル
├── main.go                        // アプリケーションの開始点。サーバーの設定や初期化を含む
├── models                         // データベースのテーブルと一致するGoの構造体を格納するディレクトリ
│   └── models.go                  // データ構造の定義を行うファイル
└── stores                         // データベースとのやり取りを行う関数を格納するディレクトリ
    ├── auth_store.go              // ユーザーの認証や登録に関するデータベース処理
    ├── post_store.go              // 投稿に関するデータベース処理
    ├── user_store.go              // ユーザー情報に関するデータベース処理
    ├── product_store.go           // 商品情報に関するデータベース処理
    └── purchase_store.go          // 購入情報に関するデータベース処理
```

ディレクトリ構造とその対応

```
.
├── README.md
├── app                         # アプリケーション層
│   └── usecases               # ユースケース
│       ├── auth_usecase.go
│       ├── post_usecase.go
│       ├── product_usecase.go
│       ├── purchase_usecase.go
│       └── user_usecase.go
├── config                      # 設定
│   ├── config.go
│   └── database.go
├── domain                      # ドメイン層
│   ├── models                 # ドメインモデル
│   │   └── models.go
│   └── rules                  # ドメインルール
│       ├── auth_rules.go
│       ├── post_rules.go
│       ├── product_rules.go
│       ├── purchase_rules.go
│       └── user_rules.go
├── go.mod                      # Goモジュール設定
├── go.sum                      # Goモジュール依存関係
├── infra                       # インフラ層
│   └── stores                 # データストア
│       ├── auth_store.go
│       ├── post_store.go
│       ├── product_store.go
│       ├── purchase_store.go
│       └── user_store.go
├── interfaces                  # インターフェース層（プレゼンテーション層）
│   ├── api                    # APIインターフェース
│   │   ├── handlers           # ハンドラ
│   │   │   ├── auth.go
│   │   │   ├── posts.go
│   │   │   ├── products.go
│   │   │   ├── purchase.go
│   │   │   └── user.go
│   │   ├── middlewares        # ミドルウェア
│   │   │   └── isAuthenticated.go
│   │   └── responses          # レスポンス
│   │       ├── error.go
│   │       └── success.go
│   └── router.go              # ルータ設定
└── main.go                    # エントリーポイント
```
