# tips

## プロジェクト初期化・依存関係管理

| コマンド                    | 説明                                                                              |
| --------------------------- | --------------------------------------------------------------------------------- |
| `go mod init <module-path>` | 新しい Go モジュールを初期化する（例: `go mod init github.com/username/project`） |
| `go mod tidy`               | 依存関係を整理（不要なものを削除、必要なものを追加）                              |
| `go get <package>`          | パッケージを取得し依存関係に追加する                                              |
| `go get -u <package>`       | パッケージを最新バージョンに更新する                                              |
| `go mod vendor`             | 依存パッケージを vendor ディレクトリにコピーする                                  |
| `go mod graph`              | モジュール依存関係のグラフを出力                                                  |
| `go list -m all`            | 全ての依存モジュールを一覧表示                                                    |

## ビルド・実行

| コマンド               | 説明                                           |
| ---------------------- | ---------------------------------------------- |
| `go build`             | パッケージをコンパイルする                     |
| `go build -o <output>` | 実行ファイル名を指定してビルド                 |
| `go run <file.go>`     | Go プログラムを直接実行する                    |
| `go run .`             | カレントディレクトリの Go プログラムを実行する |
| `go install`           | バイナリをコンパイルして GOPATH/bin に配置     |
| `go clean`             | オブジェクトファイルを削除                     |

## テスト・解析

| コマンド                          | 説明                                         |
| --------------------------------- | -------------------------------------------- |
| `go test`                         | テストを実行する                             |
| `go test -v`                      | 詳細なテスト出力を表示                       |
| `go test -cover`                  | テストカバレッジを表示                       |
| `go test -coverprofile=cover.out` | カバレッジプロファイルを出力                 |
| `go tool cover -html=cover.out`   | カバレッジを HTML 形式で表示                 |
| `go vet`                          | コードの潜在的な問題を検出                   |
| `go fmt`                          | コードを標準フォーマットに整形               |
| `golint`                          | スタイルに関するベストプラクティスをチェック |

## クロスコンパイル

| コマンド                             | 説明                                         |
| ------------------------------------ | -------------------------------------------- |
| `GOOS=linux GOARCH=amd64 go build`   | Linux (amd64)用にビルド                      |
| `GOOS=windows GOARCH=amd64 go build` | Windows (amd64)用にビルド                    |
| `GOOS=darwin GOARCH=amd64 go build`  | macOS (amd64)用にビルド                      |
| `go tool dist list`                  | サポートされているプラットフォームを一覧表示 |

## ドキュメント・情報表示

| コマンド                      | 説明                                                                           |
| ----------------------------- | ------------------------------------------------------------------------------ |
| `go doc <package>`            | パッケージのドキュメントを表示                                                 |
| `go doc <package>.<function>` | 特定の関数のドキュメントを表示                                                 |
| `godoc -http=:6060`           | ローカルで Godoc サーバーを起動（ブラウザで http://localhost:6060 にアクセス） |
| `go env`                      | Go 環境変数を表示                                                              |
| `go version`                  | Go のバージョンを表示                                                          |

## モジュール・ツール関連

| コマンド                            | 説明                                 |
| ----------------------------------- | ------------------------------------ |
| `go install <module>@latest`        | 最新バージョンのツールをインストール |
| `go list -f '{{.Dir}}' -m <module>` | モジュールのパスを表示               |

## gRPC 関連

| コマンド                                                                                                            | 説明                                                    |
| ------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------- |
| `go install google.golang.org/protobuf/cmd/protoc-gen-go@latest`                                                    | Protocol Buffers 用の Go コード生成ツールをインストール |
| `go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest`                                                   | gRPC 用の Go コード生成ツールをインストール             |
| `protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative <proto-file>` | .proto ファイルから Go コードを生成                     |

## パフォーマンス解析

| コマンド                                            | 説明                                     |
| --------------------------------------------------- | ---------------------------------------- |
| `go test -bench=.`                                  | ベンチマークを実行                       |
| `go test -benchmem`                                 | メモリアロケーションも含めたベンチマーク |
| `go tool pprof <profile-file>`                      | プロファイルデータを解析                 |
| `go test -cpuprofile cpu.prof -memprofile mem.prof` | CPU とメモリのプロファイルを生成         |

## データベースマイグレーション (外部ツール)

| コマンド                                                             | 説明                           |
| -------------------------------------------------------------------- | ------------------------------ |
| `go install github.com/golang-migrate/migrate/v4/cmd/migrate@latest` | migrate ツールをインストール   |
| `migrate create -ext sql -dir migrations -seq create_users_table`    | マイグレーションファイルを作成 |
| `migrate -database "${DB_URL}" -path migrations up`                  | マイグレーションを実行         |

## コード品質・静的解析 (外部ツール)

| コマンド                                                                | 説明                                           |
| ----------------------------------------------------------------------- | ---------------------------------------------- |
| `go install golang.org/x/tools/cmd/goimports@latest`                    | インポート文を自動整理するツールをインストール |
| `go install honnef.co/go/tools/cmd/staticcheck@latest`                  | 高度な静的解析ツールをインストール             |
| `go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest` | Go の総合的な Lint ツールをインストール        |
| `golangci-lint run`                                                     | コードをリント（多数の Linter をまとめて実行） |
