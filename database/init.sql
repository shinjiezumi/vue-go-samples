create table users
(
    id                bigint unsigned auto_increment primary key comment 'ユーザーID',
    name              varchar(255) not null comment 'ユーザー名',
    email             varchar(255) not null comment 'メールアドレス',
    password          varchar(255) not null comment 'パスワード',
    created_at        timestamp    null comment '作成日時',
    updated_at        timestamp    null comment '更新日時',
    constraint users_name_unique unique (name)
    constraint users_email_unique unique (email)
) collate = utf8mb4_unicode_ci comment='ユーザー';

