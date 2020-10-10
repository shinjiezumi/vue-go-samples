
-- +migrate Up
DROP TABLE IF EXISTS `todos`;
CREATE TABLE `todos`
(
    id          bigint unsigned AUTO_INCREMENT PRIMARY KEY COMMENT 'TodoId',
    user_id     bigint unsigned COMMENT 'ユーザーID',
    title       varchar(128) NOT NULL COMMENT 'タイトル',
    memo        varchar(255) COMMENT 'メモ',
    limit_date  date         NOT NULL COMMENT '期日',
    finished_at datetime COMMENT '完了日',
    created_at  datetime DEFAULT CURRENT_TIMESTAMP COMMENT '作成日時',
    updated_at  datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新日時',
    KEY `idx_user_id` (`user_id`)
) ENGINE = INNODB
  CHARSET = utf8mb4 COMMENT ='Todo';

INSERT INTO todos (user_id, title, memo, limit_date, finished_at, created_at, updated_at)
VALUES (1, 'スマホ表示改善', '未ログイン時のヘッダーがイマイチなので直す', '2020-02-19', '2020-03-05 11:03:19', '2020-02-18 10:26:00',
        '2020-10-03 07:34:37');
INSERT INTO todos (user_id, title, memo, limit_date, finished_at, created_at, updated_at)
VALUES (1, 'サンプルアプリ作る', '・機能は認証＋CRUDでとりあえず形にする
・フロントエンドはvue/vuex/vuetifyでやる
・サーバーサイドはgo/ginでやる
', '2020-02-11', '2020-03-04 18:50:16', '2020-03-04 08:55:21', '2020-09-26 14:20:42');
INSERT INTO todos (user_id, title, memo, limit_date, finished_at, created_at, updated_at)
VALUES (1, 'CloudRun/CloudSQLにデプロイ', '・マルチステージビルド使えばDockerfileはそれぞれ１つで済みそう
・CloudRunとCloudSQLはTCP接続できないらしいので要修正', '2020-02-16', '2020-03-04 18:50:50', '2020-03-04 08:56:12',
        '2020-03-06 19:01:59');
INSERT INTO todos (user_id, title, memo, limit_date, finished_at, created_at, updated_at)
VALUES (1, 'ローディング表示対応', '', '2020-03-06', null, '2020-03-05 05:00:04', '2020-09-15 17:36:14');
INSERT INTO todos (user_id, title, memo, limit_date, finished_at, created_at, updated_at)
VALUES (1, '自動デプロイ、Dockerfile見直し', '・CircleCIで自動デプロイ
・Dockerfileの最適化', '2020-03-10', '2020-03-11 18:01:50', '2020-03-05 05:01:04', '2020-03-11 09:01:50');

-- +migrate Down
DROP TABLE `todos`;
