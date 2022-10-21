## techblog-linter
日本語の文章をlintしてくれるオレオレ設定

## 必要なもの
- nodejsの実行環境
- [mattn/efm-langserver](https://github.com/mattn/efm-langserver)

## 設定
### textlintもろもろのインストール
```sh
$ npm ci
```

### efm-langserverの設定を置く
`~/.config/efm-langserver/config.yaml` を自分で定義する or [これ](https://raw.githubusercontent.com/kei-s16/dotfiles/master/.config/efm-langserver/config.yaml)を持ってくる
