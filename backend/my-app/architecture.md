## はじめに

このプロジェクトは、タスク管理アプリとなっている。


## APIについて

| エンドポイント | HTTPメソッド | 説明 |
| ---- | ---- | ---- |
| /v1/accounts | POST | アカウントの作成を行う |
| /v1/accounts | GET | アカウント一覧を取得する |
| /v1/account/{{account_id}} | GET | アカウント情報を取得する |
| /v1/account/{{account_id}} | POST | アカウントの情報を更新する |
| /v1/account/{{account_id}} | DELETE | アカウントを削除する |
| /v1/auth/login | POST | ログイン |
| /v1/auth/logout | DELETE | ログアウト |
| /v1/auth/ | GET | ログインされているか |


## リクエストとレスポンス

### アカウントの作成

`Request`
```bash
curl -i --request POST 'http://localhost:1323/v1/accounts' \
--header 'Content-Type: application/json' \
--data-raw '{
  "name": "Test",
  "password": "xxxxx123"
}'
```

- IDはUUIDで生成する
- passwordは半角英数含む８文字以上のみ
- このバリデーションは、フロントでも行う
- パスワードはハッシュ化して保存する


`Response`
```json
{
  "id":"5cf59c6c-0047-4b13-a118-65878313e329",
  "name":"Test",
  "created_at":"2020-11-02T14:50:46Z",
  "updated_at":"2020-11-02T14:50:46Z"
}
```

### アカウントの一覧を取得する

`Request`

```bash
curl -i --request GET 'http://localhost:1323/v1/accounts'
```

`Response`
```json
[
  {
    "id":"5cf59c6c-0047-4b13-a118-65878313e329",
    "name":"Test",
  }
]
```


### アカウント情報を取得する

```bash
curl -i --request GET 'http://localhost:1323/v1/accounts/{{account_id}}' \
--header 'Authorization: Bearer <session_id>'
```

`Response`
```json

{
  "id":"5cf59c6c-0047-4b13-a118-65878313e329",
  "name":"Test",
  "created_at":"2020-11-02T14:50:46Z",
  "updated_at":"2020-11-02T14:50:46Z"
}

```

### アカウントの情報を更新する
`Request`

```bash
curl -i --request POST 'http://localhost:1323/v1/account/{{account_id}}' \
--header 'Authorization: Bearer <session_id>' \
--header 'Content-Type: application/json' \
--data-raw '{
  "name": "Updated Test"
}'
```
`Response`

```json
{
  "id":"5cf59c6c-0047-4b13-a118-65878313e329",
  "name":"Updated Test",
  "created_at":"2020-11-02T14:50:46Z",
  "updated_at":"2020-11-02T15:30:15Z"
}
```

### アカウントを削除する
`Request`

```bash
curl -i --request DELETE 'http://localhost:1323/v1/account/{{account_id}}' \
--header 'Authorization: Bearer <session_id>'
```

`Response`

```json
{
  "message": "Account deleted successfully"
}
```


### ログイン
`Request`
```bash
curl -i --request POST 'http://localhost:1323/v1/auth/login' \
--header 'Content-Type: application/json' \
--data-raw '{
  "name": "Test",
  "password": "xxxxx123"
}'

```

  - usernameとpasswordを受け取り、DBと照合する
  - パスワードをハッシュ化して、DBと照合する

`Response`

```json
{
  "session_id": "xxxxx-xxxx-xxxx-xxxx-xxxx",
  "message": "Login successful"
}
```

### ログアウト

`Request`

```bash
curl -i --request DELETE 'http://localhost:1323/v1/auth/logout' \
--header 'Authorization: Bearer <session_id>'
```

`Response`

```json
{
  "message": "Logout successful"
}
```


### ログイン状態の確認

```bash
curl -i --request GET 'http://localhost:1323/v1/auth/' \
--header 'Authorization: Bearer <session_id>'
```

`Response`

```json
{
  "message": "Logged in" or "Not logged in"
}
```

- not logged inの場合は、403エラーののち、ログイン画面にリダイレクトする
- Logged inの場合は、アクセスを許可する




