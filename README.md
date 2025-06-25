# Romantic Situation Generator

シチュエーションをランダムに生成するWebアプリケーション。場所、時間、キャラクター、状況、ムードを組み合わせて、ユニークなシチュエーションを作成します。

## 機能

- ランダムなシチュエーション生成
- 要素別の表示（場所、時間、キャラクター、状況、ムード）
- 物語形式での表示
- レスポンシブデザイン
- JSON APIエンドポイント

## 技術スタック

- Go 1.24.4
- HTML/CSS/JavaScript
- Docker

## ローカルでの実行方法

```bash
# リポジトリのクローン
git clone https://github.com/yourusername/romantic-situation.git
cd romantic-situation

# アプリケーションの実行
go run .
```

デフォルトでは `http://localhost:8080` でアプリケーションが起動します。
ポート番号は環境変数 `PORT` で変更可能です：

```bash
PORT=3000 go run .
```

## API エンドポイント

### GET /api/generate

新しいシチュエーションを生成します。

レスポンス例：
```json
{
    "location": "山奥の旅館",
    "time": "深夜2時",
    "character": "無口な後輩",
    "situation": "一緒の部屋になる",
    "mood": "癒し",
    "story_text": "ある深夜2時のこと。山奥の旅館で無口な後輩と二人きりになった。\nそんな時、突然一緒の部屋になる…\n雰囲気は癒しに包まれていった…"
}
```

## Docker での実行

```bash
# イメージのビルド
docker build -t romantic-situation .

# コンテナの実行
docker run -p 8080:8080 romantic-situation
```

## ライセンス

MIT 