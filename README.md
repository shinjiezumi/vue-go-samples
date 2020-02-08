# vue-go-todo

# セットアップ
```
$ git clone https://github.com/shinjiezumi/vue-go-todo.git
$ docker-compose up -d
$ docker-compose exec api dep ensure
$ docker-compose exec front npm install
```

# 実行
```
$ docker-compose exec api gin -i -p 8080 run
$ docker-compose exec front npm run serve
```