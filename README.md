# vue-go-samples
Vue/Vuex + Go/Gin + MySQLを使ったサンプルアプリ

## Todolist
![todolist](https://user-images.githubusercontent.com/41136277/74259276-00336980-4d3b-11ea-81cd-ea75b3742971.png)

認証ありの簡単なTODOリストアプリ。

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

## セットアップ
```
$ git clone https://github.com/shinjiezumi/vue-go-samples.git
$ docker-compose up -d
$ docker-compose exec api dep ensure
$ docker-compose exec front npm install
```

## 実行
```
$ docker-compose exec api gin -i -p 8080 run
$ docker-compose exec front npm run serve
```

http://localhost:8080 開く

