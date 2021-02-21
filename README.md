# vue-go-samples
Vue/Vuex + Go/Gin + MySQLを使ったサンプルアプリ

- Todolist

- goroutine/channelを使った検索アプリ(comming soon...)

## Todolist

![Todolist](./src/front/public/img/TodoList.png)

認証ありの簡単なTODOリストアプリ。

## Searcher

![Searcher](./src/front/public/img/Searcher.png)

goroutine/channelを使った複数Webサービスを横断した検索アプリ。

## 技術スタック

### フロントエンド

- vue
- vuex
- vueRouter
- vuetify
- vuelidate

### サーバーサイド
- golang
- gin
- gorm
- gin-gwt
- sql-migrate

## セットアップ
```
$ git clone https://github.com/shinjiezumi/vue-go-samples.git
$ cd vue-go-samples
$ docker-compose up -d
$ docker-compose exec api sql-migrate up
$ docker-compose exec front npm install
```

## サーバー起動
```
$ docker-compose exec api gin -i -p 8080 run
$ docker-compose exec front npm run serve
```

http://localhost:8080 開く

