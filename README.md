[![Test](https://github.com/cohky16/invite/actions/workflows/cicd.yml/badge.svg)](https://github.com/cohky16/invite/actions/workflows/cicd.yml)

# 概要

discordのボイスチャンネルへの招待を送信できます

- 招待先のボイスチャンネルに既に2人以上いた場合は、空いているボイスチャンネルへの招待を送信します

## 各種コマンド

### invite

ユーザーにボイスチャンネルへの招待情報を送信します

- 招待者が既にボイスチャンネルに参加している場合は、招待者が参加しているボイスチャンネルへの招待情報を送信します

```shell
# 通常コマンド
!invite @hoge @fuga

# スラッシュコマンド
/invite to: @hoge @fuga
```

### invite $(RoomNo)

ユーザーに特定のボイスチャンネルへの招待情報を送信します

```shell
# 通常コマンド
!invite 1 @hoge @fuga

# スラッシュコマンド
/invite to: @hoge @fuga channel: Piyo
```

### help

ヘルプを表示します
