-- users の作成
CREATE TABLE USERS (
    id VARCHAR(255) PRIMARY KEY,
    username VARCHAR(255) NOT NULL,
    password VARCHAR(255) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- プロジェクトの作成
CREATE TABLE PROJECTS (
    id VARCHAR(255) PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    description TEXT,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- タスクの作成
CREATE TABLE TASKS (
    id VARCHAR(255) PRIMARY KEY,
    project_id VARCHAR(255) NOT NULL,
    assigned_user_id VARCHAR(255) NOT NULL,
    title VARCHAR(255) NOT NULL,
    description TEXT,
    deadline TIMESTAMP NOT NULL,
    status ENUM('todo', 'doing', 'done') DEFAULT 'todo',
    progress INTEGER NOT NULL DEFAULT 0,
    priority ENUM('low', 'middle', 'high') DEFAULT 'middle',
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (project_id) REFERENCES PROJECTS(id),
    FOREIGN KEY (assigned_user_id) REFERENCES USERS(id)
);

-- コメントの作成
CREATE TABLE COMMENTS (
    id VARCHAR(255) PRIMARY KEY,
    task_id VARCHAR(255) NOT NULL,
    user_id VARCHAR(255) NOT NULL,
    content TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (task_id) REFERENCES TASKS(id),
    FOREIGN KEY (user_id) REFERENCES USERS(id)
);
