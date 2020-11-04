# Secret DB

## 概要

Cosmos SDK を利用して秘匿化ブロックチェーンデータベースを作成する．

- コンセンサスアルゴリズム : Tendermint
- 秘匿化 : Intel SGX + Graphene
- データベース : MongoDB + MapReduce

## 構成

```bash
secretdb/
├── app
│   ├── app.go
│   ├── export.go
│   └── prefix.go
├── cmd : コマンド関連
│   ├── secretdbcli
│   │   └── main.go
│   └── secretdbd
│       ├── genaccounts.go
│       └── main.go
├── config.yml
├── go.mod
├── go.sum
├── vue : フロントエンド
│   ├── README.md
│   ├── babel.config.js
│   ├── package-lock.json
│   ├── package.json
│   ├── public
│   │   ├── favicon.ico
│   │   └── index.html
│   ├── src
│   │   ├── App.vue
│   │   ├── main.js
│   │   ├── router
│   │   │   └── index.js
│   │   ├── store
│   │   │   └── index.js
│   │   └── views
│   │       └── Index.vue
│   └── vue.config.js
└── x
    ├── mongodb : mongo-driverのWrapper
    │   ├── connection.go : LocalのMongoDBとのコネクション管理
    │   ├── query.go : MongoDBに対するクエリ関連
    │   └── types.go : 結果格納用の構造体定義
    └── secretdb : 本体
        ├── abci.go
        ├── client
        │   ├── cli : secretcbcli
        │   │   ├── query.go
        │   │   ├── queryItem.go
        │   │   ├── tx.go
        │   │   └── txItem.go
        │   └── rest : REST
        │       ├── queryItem.go
        │       ├── rest.go
        │       └── txItem.go
        ├── genesis.go
        ├── handler.go : メッセージハンドラ
        ├── handlerMsgDeleteItem.go
        ├── handlerMsgDeleteItems.go
        ├── handlerMsgStoreItem.go
        ├── handlerMsgUpdateItem.go
        ├── handlerMsgUpdateItems.go
        ├── keeper : データベースとのやり取りを管理
        │   ├── item.go
        │   ├── keeper.go
        │   ├── params.go
        │   └── querier.go
        ├── module.go
        ├── spec
        │   └── README.md
        └── types : 構造体等の定義
            ├── MsgDeleteItem.go
            ├── MsgDeleteItems.go
            ├── MsgStoreItem.go
            ├── MsgUpdateItem.go
            ├── MsgUpdateItems.go
            ├── TypeItem.go
            ├── codec.go
            ├── errors.go
            ├── events.go
            ├── expected_keepers.go
            ├── genesis.go
            ├── key.go
            ├── msg.go
            ├── params.go
            ├── querier.go
            └── types.go
```
