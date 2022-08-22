# 典原（Dictland）

**典原（Dictland）** 是一個開放且公共的面向所有詞典區塊鏈項目，基於 Cosmos SDK 和 Tendermint，並使用 [Ignite CLI](https://ignite.com/cli) 構建。

## 開始

```
ignite chain serve
```

`serve` 命令安裝依賴、構建、初始化和以開發模式啟動 Dictland。

### 配置

開發模式中的區塊鏈可通過 `config.yml` 配置。更多內容參見 [Ignite CLI docs](https://docs.ignite.com)。

## 發佈

要發佈區塊鏈新版本的話，創建並推送一個新的以 `v` 爲前綴的標籤即可。

```
git tag v0.1
git push origin v0.1
```

### 安裝

執行下面的命令來安裝最新版本的 Dictland：

```
curl https://get.ignite.com/tsunhua/dictland@latest! | sudo bash
```

## 參考

- [Ignite CLI docs](https://docs.ignite.com)
- [Cosmos SDK docs](https://docs.cosmos.network)
