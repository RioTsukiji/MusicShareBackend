# MusicShareBackend
This is the repository for music share app.

## ディレクトリ構成
```
.
├── README.md
├── main.go
├── go.mod
├── go.sum
├── interfaces // インターフェース層
│   ├── handler
│   │   ├── music.go
│   │   └── user.go
│   │
│   └── response
│       └── response.go
│       
├── usecase // アプリケーション層
│   ├── music.go
│   └── user.go
│
├── domain // ドメイン層
│   ├── repository
│   │   ├── music_repository.go
│   │   └── user_repository.go
│   │
│   ├── music.go
│   └── user.go
│
├── infrastructure // インフラ層
│   └── persistence
│       ├── music.go
│       └── user.go
│
└── internal // 内部パッケージ
    └── config
        └── database.go //DBの起動

```
参考：https://qiita.com/tono-maron/items/345c433b86f74d314c8d#%E5%AF%BE%E8%B1%A1%E8%AA%AD%E8%80%85

## データベース