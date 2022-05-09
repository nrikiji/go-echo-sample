# go-echo-starter

`go、echo`で毎回行っている作業や設定をあらかじめテンプレート化したプロジェクトです。

`go`のバージョンは`1.18`

主に含まれていることは以下のとおり。

- `sql-migrate`によるマイグレーション
- `firebase auth`による認証
- `github actions`でのテスト
- `development`、`production`ごとに設定ファイルを切り替える

## セットアップ手順

### プロジェクトのクローン

```
$ git clone https://github.com/nrikiji/go-echo-starter
```

### `config.yml`の編集

DB 設定を環境に合わせて更新

### マイグレーション実行

```bash
# 開発環境
$ sql-migrate up -env development -config config.yml

# テスト環境
$ sql-migrate up -env test -config config.yml
```

### `firebase`設定ファイルの追加

`firebase`コンソールから`firebase_secret_key.json`をダウンロードしてプロジェクトルートに追加（`git`管理対象外）

## 起動

```bash
# 開発
$ go run server.go

# テスト
$ go test ./...

# ビルド
$ go build -o server .
```
