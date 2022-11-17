## techblog-linter
日本語の文章をlintしてくれるオレオレ設定

## 必要なもの
### 共通
- nodejsの実行環境
- (daggerをローカルで動かす場合)goの実行環境

### vim
- [mattn/efm-langserver](https://github.com/mattn/efm-langserver)

### vscode
- [vscode-textlint](https://marketplace.visualstudio.com/items?itemName=taichi.vscode-textlint)

## 設定
### textlintもろもろのインストール
```sh
$ npm ci
```

### efm-langserverの設定を置く
`~/.config/efm-langserver/config.yaml` を自分で定義する or [これ](https://raw.githubusercontent.com/kei-s16/dotfiles/master/.config/efm-langserver/config.yaml)を持ってくる
