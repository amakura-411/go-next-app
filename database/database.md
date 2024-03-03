


```bash

mysql> SHOW COLUMNS FROM USERS;
+------------+--------------+------+-----+-------------------+-------------------+
| Field      | Type         | Null | Key | Default           | Extra             |
+------------+--------------+------+-----+-------------------+-------------------+
| id         | varchar(255) | NO   | PRI | NULL              |                   |
| username   | varchar(255) | NO   |     | NULL              |                   |
| password   | varchar(255) | NO   |     | NULL              |                   |
| created_at | timestamp    | NO   |     | CURRENT_TIMESTAMP | DEFAULT_GENERATED |
+------------+--------------+------+-----+-------------------+-------------------+
4 rows in set (0.02 sec)

mysql> SHOW COLUMNS FROM PROJECTS;
+-------------+--------------+------+-----+-------------------+-------------------+
| Field       | Type         | Null | Key | Default           | Extra             |
+-------------+--------------+------+-----+-------------------+-------------------+
| id          | varchar(255) | NO   | PRI | NULL              |                   |
| title       | varchar(255) | NO   |     | NULL              |                   |
| description | text         | YES  |     | NULL              |                   |
| created_at  | timestamp    | NO   |     | CURRENT_TIMESTAMP | DEFAULT_GENERATED |
| updated_at  | timestamp    | NO   |     | CURRENT_TIMESTAMP | DEFAULT_GENERATED |
+-------------+--------------+------+-----+-------------------+-------------------+
5 rows in set (0.00 sec)

mysql> SHOW COLUMNS FROM COMMENTS;
+------------+--------------+------+-----+-------------------+-------------------+
| Field      | Type         | Null | Key | Default           | Extra             |
+------------+--------------+------+-----+-------------------+-------------------+
| id         | varchar(255) | NO   | PRI | NULL              |                   |
| task_id    | varchar(255) | NO   | MUL | NULL              |                   |
| user_id    | varchar(255) | NO   | MUL | NULL              |                   |
| content    | text         | NO   |     | NULL              |                   |
| created_at | timestamp    | NO   |     | CURRENT_TIMESTAMP | DEFAULT_GENERATED |
| updated_at | timestamp    | NO   |     | CURRENT_TIMESTAMP | DEFAULT_GENERATED |
+------------+--------------+------+-----+-------------------+-------------------+
6 rows in set (0.00 sec)

mysql> SHOW COLUMNS FROM TASKS;
+------------------+-----------------------------+------+-----+-------------------+-------------------+
| Field            | Type                        | Null | Key | Default           | Extra             |
+------------------+-----------------------------+------+-----+-------------------+-------------------+
| id               | varchar(255)                | NO   | PRI | NULL              |                   |
| project_id       | varchar(255)                | NO   | MUL | NULL              |                   |
| assigned_user_id | varchar(255)                | NO   | MUL | NULL              |                   |
| title            | varchar(255)                | NO   |     | NULL              |                   |
| description      | text                        | YES  |     | NULL              |                   |
| deadline         | timestamp                   | NO   |     | NULL              |                   |
| status           | enum('todo','doing','done') | YES  |     | todo              |                   |
| progress         | int                         | NO   |     | 0                 |                   |
| priority         | enum('low','middle','high') | YES  |     | middle            |                   |
| created_at       | timestamp                   | NO   |     | CURRENT_TIMESTAMP | DEFAULT_GENERATED |
| updated_at       | timestamp                   | NO   |     | CURRENT_TIMESTAMP | DEFAULT_GENERATED |
+------------------+-----------------------------+------+-----+-------------------+-------------------+
11 rows in set (0.00 sec)
```