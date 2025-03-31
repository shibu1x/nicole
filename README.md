# Nicole

ニコニコ動画のランキングビューアです。1時間ごとにランキングを取得し、ブラウザで閲覧できます。

## 機能

- ニコニコ動画のランキングを定期的に取得
- カテゴリごとのランキング表示
- 投稿者単位での非表示機能
- 有料動画の表示/非表示
- レスポンシブデザイン対応

## 技術スタック

### バックエンド

- Go 1.24
- Dependencies:
  - goquery: HTMLスクレイピング
  - gjson: JSON解析

### フロントエンド

- Vue.js 3
- Vite
- Dependencies:
  - axios: HTTP通信
  - sass: スタイリング

## 開発環境のセットアップ

### バックエンド

```bash
cd back
go mod download
go run main.go
```

### フロントエンド

```bash
cd front
npm install
npm run dev
```

## デプロイメント

GitHubリポジトリにpushすると、GitHub Actionsによって以下が自動的に実行されます：

1. バックエンドのDockerイメージのビルドとGitHub Container Registryへのプッシュ
2. フロントエンドのビルドとGitHub Pagesへのデプロイ
3. 1時間ごとのランキング更新（AWS S3へのアップロード）

## ライセンス

MIT License