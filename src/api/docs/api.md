# Group User
ユーザー

## ユーザー [/user]

### 個人データ取得 [GET]

+ Response 200 (application/json)

    + Body

        ```json
        {
          "user": {
            "id": 123456789,
            "name": "the40san",
            "avater_id": 1,
            "rank": 123,
            "exp": 12345678
          }
        }
        ```

### ユーザー設定変更 [PATCH]

+ Attributes

    + name: `` (string, optional) - ユーザー名
    + avater_id: `` (string, optional) - アバターID

+ Request example ユーザー名変更 (application/json)

    + Body

        ```json
        {
          "name": "つよい"
        }
        ```
+ Request example アバターID変更 (application/json)

    + Body

        ```json
        {
          "avater_id": 2
        }
        ```

+ Response 200 (application/json)

    + Body

        ```json
        {
          "user": {
            "id": 123456789,
            "name": "the40san",
            "avater_id": 1,
            "rank": 123,
            "exp": 12345678
          }
        }
        ```