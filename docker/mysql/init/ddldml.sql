-- DDL
create table users
(
    id                bigint unsigned auto_increment primary key comment 'ユーザーID',
    name              varchar(255) not null comment 'ユーザー名',
    email             varchar(255) not null comment 'メールアドレス',
    password          varchar(255) not null comment 'パスワード',
    created_at        timestamp    null comment '作成日時',
    updated_at        timestamp    null comment '更新日時',
    constraint users_name_unique unique (name),
    constraint users_email_unique unique (email)
) collate = utf8mb4_unicode_ci comment='ユーザー';

create table todos
(
    id          bigint unsigned auto_increment primary key comment 'TodoId',
    user_id     bigint unsigned comment 'ユーザーID',
    title       varchar(128) not null comment 'タイトル',
    memo        varchar(255) null comment 'メモ',
    limit_date  date not null comment '期日',
    finished_at timestamp null comment '完了日',
    created_at  timestamp not null comment '作成日時',
    updated_at  timestamp not null comment '更新日時',
    index       todos_user_id_index (user_id)
) collate = utf8mb4_unicode_ci comment='Todoリスト';

-- DML
insert into
    users(name, email, password, created_at, updated_at)
values
    ('test', 'test@test.com', '$2a$10$LvsHcEMnYFFe1taM1sSr1eLjoZ740O7o5M5aFB05RuI6yY0mHa.1u', now(), now());