# はじめに

このプロジェクトは、タスク管理アプリとなっている。

# API について

| エンドポイント                     | HTTP メソッド | 説明                               |
| ---------------------------------- | ------------- | ---------------------------------- |
| /v1/users                          | POST          | アカウントの作成を行う             |
| /v1/users                          | GET           | アカウント一覧を取得する           |
| /v1/user/{{user_id}}               | GET           | アカウント情報を取得する           |
| /v1/user/{{user_id}}               | POST          | アカウントの情報を更新する         |
| /v1/user/{{user_id}}               | DELETE        | アカウントを削除する               |
| /v1/auth/login                     | POST          | ログイン                           |
| /v1/auth/logout                    | DELETE        | ログアウト                         |
| /v1/auth/                          | GET           | ログインされているか               |
| /v1/project                        | POST          | プロジェクトの作成                 |
| /v1/projects                       | GET           | プロジェクト一覧を取得する         |
| /v1/project/{{project_id}}         | GET           | プロジェクトの情報を取得する       |
| /v1/project/{{project_id}}         | POST          | プロジェクトの情報を更新する       |
| /v1/project/{{project_id}}         | DELETE        | プロジェクトを削除する             |
| /v1/project/{{project_id}}/members | POST          | プロジェクトにメンバーを追加する   |
| /v1/project/{{project_id}}/members | DELETE        | プロジェクトからメンバーを削除する |

# リクエストとレスポンス

## アカウント関連

### アカウントの作成

`Request`

```bash
curl -i --request POST 'http://localhost:1323/v1/users' \
--header 'Content-Type: application/json' \
--data-raw '{
  "name": "Test",
  "password": "xxxxx123"
}'
```

- ID は UUID で生成する
- password は半角英数含む８文字以上のみ
- このバリデーションは、フロントでも行う
- パスワードはハッシュ化して保存する

`Response`

```json
{
  "id": "5cf59c6c-0047-4b13-a118-65878313e329",
  "name": "Test",
  "icon": "",
  "created_at": "2020-11-02T14:50:46Z",
  "updated_at": "2020-11-02T14:50:46Z"
}
```

### アカウントの一覧を取得する

`Request`

```bash
curl -i --request GET 'http://localhost:1323/v1/users'
```

`Response`

```json
[
  {
    "id": "5cf59c6c-0047-4b13-a118-65878313e329",
    "name": "Test",
    "icon": "",
    "created_at": "2020-11-02T14:50:46Z",
    "updated_at": "2020-11-02T14:50:46Z"
  }
]
```

### アカウント情報を取得する

```bash
curl -i --request GET 'http://localhost:1323/v1/users/{{user_id}}' \
--header 'Authorization: Bearer <session_id>'
```

`Response`

```json
{
  "id": "5cf59c6c-0047-4b13-a118-65878313e329",
  "name": "Test",
  "icon": "",
  "created_at": "2020-11-02T14:50:46Z",
  "updated_at": "2020-11-02T14:50:46Z"
}
```

### アカウントの情報を更新する

`Request`

```bash
curl -i --request POST 'http://localhost:1323/v1/user/{{user_id}}' \
--header 'Authorization: Bearer <session_id>' \
--header 'Content-Type: application/json' \
--data-raw '{
  "name": "Updated Test"
}'
```

`Response`

```json
{
  "id": "5cf59c6c-0047-4b13-a118-65878313e329",
  "name": "Test",
  "icon": "xx.jpg",
  "created_at": "2020-11-02T14:50:46Z",
  "updated_at": "2020-11-02T14:50:46Z"
}
```

### アカウントを削除する

`Request`

```bash
curl -i --request DELETE 'http://localhost:1323/v1/user/{{user_id}}' \
--header 'Authorization: Bearer <session_id>'
```

`Response`

```json
{
  "message": "user deleted successfully"
}
```

---

## 認証関連

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

- username と password を受け取り、DB と照合する
- パスワードをハッシュ化して、DB と照合する

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

`Request`

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

- not logged in の場合は、403 エラーののち、ログイン画面にリダイレクトする
- Logged in の場合は、アクセスを許可する

---

## プロジェクト関連

### プロジェクトの作成

`Request`

```bash
curl -i --request POST 'http://localhost:1323/v1/project' \
--header 'Authorization: Bearer <session_id>' \
--header 'Content-Type: application/json' \
--data-raw '{
  "user_id" : "5cf59c6c-0047-4b13-a118-65878313e329",
  "project_title": "Test Project",
  "project_description": "This is a test project",
  "goal_date": "2020-11-02T14:50:46Z",
  "display_flag": TRUE,
}'
```

- user_id はログイン中のユーザーの ID
- goal_date は期限日
- display_flag は表示するかどうか

`Response`

```json
{
  "id": "5cf59c6c-0047-4b13-a118-65878313e329",
  "user_id": "5cf59c6c-0047-4b13-a118-65878313e329",
  "project_title": "Test Project",
  "project_description": "This is a test project",
  "goal_date": "2020-11-02T14:50:46Z",
  "display_flag": TRUE,
  "created_at": "2020-11-02T14:50:46Z",
  "updated_at": "2020-11-02T14:50:46Z"
}
```

### プロジェクトの一覧を取得する

- 一覧を表示する
- ただし、display_flag が false かつメンバーでない時、非表示

```bash
curl -I --request GET 'http://localhost:1323/v1/projects' \
--header 'Authorization Bearer <session_id>'
--header 'Content-Type: application/json' \
--data-raw '{
  "user_id" : "5cf59c6c-0047-4b13-a118-65878313e329"
}'
```

`Response`

```json
[
  {
    "id": "5cf59c6c-0047-4b13-a118-65878313e329",
    "project_title": "Test Project",
    "project_description": "This is a test project",
    "goal_date": "2020-11-02T14:50:46Z",
    "display_flag": TRUE,
    "created_at": "2020-11-02T14:50:46Z",
    "updated_at": "2020-11-02T14:50:46Z",
    "members" : [
      {
        "user_id": "5cf59c6c-0047-4b13-a118-65878313e329",
        "user_name": "Test",
      }
    ]
  }
]
```

### プロジェクトの情報を取得する

- project の詳細を表示する
- ただし、閲覧権限があるかどうかは、project のメンバーかどうかで判断する
- もし、メンバーでない場合、403 エラーを返し、リダイレクトする

```bash
curl -I --request GET 'http://localhost:1323/v1/project/{{project_id}}' \
--header 'Authorization Bearer <session_id>'
--header 'Content-Type: application/json' \
--data-raw '{
  "user_id" : "5cf59c6c-0047-4b13-a118-65878313e329"
}'
```

`Response`

```json
{
  "id": "5cf59c6c-0047-4b13-a118-65878313e329",
  "project_title": "Test Project",
  "project_description": "This is a test project",
  "goal_date": "2020-11-02T14:50:46Z",
  "display_flag": TRUE,
  "created_at": "2020-11-02T14:50:46Z",
  "updated_at": "2020-11-02T14:50:46Z",
  "members" : [
    {
      "user_id": "5cf59c6c-0047-4b13-a118-65878313e329",
      "user_name": "Test",
    }
  ]
}
```

### プロジェクトの情報を更新する

```bash
curl -I --request POST 'http://localhost:1323/v1/project/{{project_id}}' \
--header '
Authorization Bearer <session_id>'
--header 'Content-Type: application/json' \
--data-raw '{
  "project_title": "Updated Test Project",
  "project_description": "This is an updated test project",
  "goal_date": "2020-11-02T14:50:46Z",
  "display_flag": TRUE,
}'
```

`Response`

```json
{
  "id": "5cf59c6c-0047-4b13-a118-65878313e329",
  "project_title": "Updated Test Project",
  "project_description": "This is an updated test project",
  "goal_date": "2020-11-02T14:50:46Z",
  "display_flag": TRUE,
  "created_at": "2020-11-02T14:50:46Z",
  "updated_at": "2020-11-02T14:50:46Z",
  "members" : [
    {
      "user_id": "5cf59c6c-0047-4b13-a118-65878313e329",
      "user_name": "Test",
    }
  ]
}
```

### プロジェクトを削除する

- project を削除する
- ただし、削除権限があるかどうかは、project のオーナーかどうかで判断する

```bash
curl -I --request DELETE 'http://localhost:1323/v1/project/{{project_id}}' \
--header 'Authorization Bearer <session_id>'
--header 'Content-Type: application/json' \
--data-raw '{
  "user_id" : "5cf59c6c-0047-4b13-a118-65878313e329"
}'
```

`Response`

```json
{
  "message": "project deleted successfully" or "You are not the owner of this project"
}
```

### プロジェクトにメンバーを追加する

- project にメンバーを追加する
- ただし、追加権限があるかどうかは、project のオーナーかどうかで判断する
- 一度に複数人の追加も可能

```bash
curl -I --request POST 'http://localhost:1323/v1/project/{{project_id}}/members' \
--header 'Authorization Bearer <session_id>'
--header 'Content-Type: application/json' \
--data-raw '{
  "user_id" : "5cf59c6c-0047-4b13-a118-65878313e329"
  "add_users" :[
    {
      "user_id" : "5cf59c6c-0047-4b13-a118-65878313e329"
    },
    {
      "user_id" : "5cf59c6c-0047-4b13-a118-65878313e329"
    }
  ]
}'
```

`Response`

```json
{
  "message": "member added successfully" or "You are not the owner of this project"
}
```

### プロジェクトからメンバーを削除する

- project に member を削除する
- ただし、削除権限があるかどうかは、project のオーナーまたは本人かどうかで判断する

```bash
curl -I --request DELETE 'http://localhost:1323/v1/project/{{project_id}}/members' \
--header 'Authorization Bearer <session_id>'
--header 'Content-Type: application/json' \
--data-raw '{
  "user_id" : "5cf59c6c-0047-4b13-a118-65878313e329"
  "delete_user": "5cf59c6c-0047-4b13-a118-65878313e320"
}'
```

`Response`

```json
{
  "message": "member deleted successfully" or "You are not the owner of this project or you are not the member of this project"
}
```

### プロジェクトにタスクを追加する

- project に task を追加する
- task は、メンバーであれば誰でも追加・削除・修正が可能
- なお、何らかの処理を行ったときは、メンバーに通知を送る
- 作成者がアサインされる

```bash
curl -I --request POST 'http://localhost:1323/v1/project/{{project_id}}/tasks' \
--header 'Authorization Bearer <session_id>'
--header 'Content-Type: application/json' \
--data-raw '{
  "user_id" : "5cf59c6c-0047-4b13-a118-65878313e329"
  "task_title" : "Test Task",
  "task_description" : "This is a test task",
  "task_deadline" : "2020-11-02T14:50:46Z",
  "task_status" : "not started"
}'
```

`Response`

```json
{
  "id": "5cf59c6c-0047-4b13-a118-65878313e329",
  "user_id": "5cf59c6c-0047-4b13-a118-65878313e329",
  "task_title": "Test Task",
  "task_description": "This is a test task",
  "task_deadline": "2020-11-02T14:50:46Z",
  "task_status": "not started",
  "created_at": "2020-11-02T14:50:46Z",
  "updated_at": "2020-11-02T14:50:46Z"
}
```

### プロジェクトのタスク一覧を取得する

- project の task 一覧を表示する
- メンバーであれば。誰でも閲覧可能

```bash
  curl -I --request GET 'http://localhost:1323/v1/project/{{project_id}}/tasks' \
  --header 'Authorization Bearer <session_id>'
  --header 'Content-Type: application/json' \
  --data-raw '{
    "user_id" : "5cf59c6c-0047-4b13-a118-65878313e329"
  }'
```

`Response`

```json
[
  {
    "id": "5cf59c6c-0047-4b13-a118-65878313e329",
    "user_id": "5cf59c6c-0047-4b13-a118-65878313e329",
    "task_title": "Test Task",
    "task_description": "This is a test task",
    "task_deadline": "2020-11-02T14:50:46Z",
    "task_status": "not started",
    "created_at": "2020-11-02T14:50:46Z",
    "updated_at": "2020-11-02T14:50:46Z"
  }
]
```

### プロジェクトのタスク情報を取得する

```bash
curl -I --request GET 'http://localhost:1323/v1/project/{{project_id}}/tasks/{{task_id}}' \

--header 'Authorization Bearer <session_id>'
--header 'Content-Type: application/json' \
--data-raw '{
  "user_id" : "5cf59c6c-0047-4b13-a118-65878313e329"
}'
```

### プロジェクトのタスク情報を更新する

```bash
curl -I --request POST 'http://localhost:1323/v1/project/{{project_id}}/tasks/{{task_id}}' \
--header 'Authorization Bearer <session_id>'
--header 'Content-Type: application/json' \
--data-raw '{
  "task_title" : "Updated Test Task",
  "task_description" : "This is an updated test task",
  "task_deadline" : "2020-11-02T14:50:46Z",
  "task_status" : "not started"
}'
```

### プロジェクトのタスクを削除する

```bash
curl -I --request DELETE 'http://localhost:1323/v1/project/{{project_id}}/tasks/{{task_id}}' \
--header 'Authorization Bearer <session_id>'
--header 'Content-Type: application/json' \
--data-raw '{
  "user_id" : "5cf59c6c-0047-4b13-a118-65878313e329"
}'
```

`Response`

```json
{
  "message": "task deleted successfully" or "Failed to delete task"
}
```
