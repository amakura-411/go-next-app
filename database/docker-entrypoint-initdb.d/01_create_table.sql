-- users の作成
-- ユーザーがいる。
CREATE TABLE USERS (
  id VARCHAR(255) PRIMARY KEY,
  username VARCHAR(255) NOT NULL,
  password VARCHAR(255) NOT NULL,
  icon VARCHAR(255),
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
  updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- プロジェクトがある
CREATE TABLE PROJECTS (
  id VARCHAR(255) PRIMARY KEY,
  project_title VARCHAR(255) NOT NULL,
  project_description TEXT,
  goal_date : TIMESTAMP NOT NULL,
  -- 表示・非表示
  display_flag BOOLEAN DEFAULT TRUE,
  -- ダッシュボードで利用
  task_counts INTEGER NOT NULL DEFAULT 0,
  task_done_counts INTEGER NOT NULL DEFAULT 0
  -- 自動連番
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
);

-- ユーザーはプロジェクトに参加する
-- 複数のユーザーがプロジェクトに参加する
-- プロジェクトにはオーナーがおり、オーナーはプロジェクトを作成したユーザーがデフォルトとなる
CREATE TABLE PROJECTS_MEMBERS (
  id VARCHAR(255) PRIMARY KEY,
  project_id VARCHAR(255) NOT NULL,
  user_id VARCHAR(255) NOT NULL,
  role ENUM('owner', 'member') DEFAULT 'member',
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  FOREIGN KEY (project_id) REFERENCES PROJECTS(id),
  FOREIGN KEY (user_id) REFERENCES USERS(id)
);

-- プロジェクトにはタスクがある
CREATE TABLE TASK (
  id VARCHAR(255) PRIMARY KEY,
  task_title VARCHAR(255) NOT NULL,
  task_description TEXT,
  deadline TIMESTAMP NOT NULL,
  status ENUM('todo', 'doing', 'done') DEFAULT 'todo',
  progress INTEGER NOT NULL DEFAULT 0,
  priority ENUM('low', 'middle', 'high') DEFAULT 'middle',
  estimated_time INTEGER NOT NULL DEFAULT 0,
  actual_time INTEGER NOT NULL DEFAULT 0
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  assigned_user_id VARCHAR(255) NOT NULL,
  FOREIGN KEY (assigned_user_id) REFERENCES USERS(id),
);


-- １つのプロジェクトには、複数のタスクがある
CREATE TABLE PROJECTS_TASKS (
  id VARCHAR(255) PRIMARY KEY,
  project_id VARCHAR(255) NOT NULL,
  task_id VARCHAR(255) NOT NULL,
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  FOREIGN KEY (project_id) REFERENCES PROJECTS(id),
  FOREIGN KEY (task_id) REFERENCES TASKS(id)
);

-- タスクにはコメントがある
CREATE TABLE TASK_COMMENTS (
  id VARCHAR(255) PRIMARY KEY,
  task_id VARCHAR(255) NOT NULL,
  user_id VARCHAR(255) NOT NULL,
  content TEXT NOT NULL,
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  FOREIGN KEY (task_id) REFERENCES TASKS(id),
  FOREIGN KEY (user_id) REFERENCES USERS(id)
);


CREATE TABLE TASK_NOTIFICATIONS (
  id VARCHAR(255) PRIMARY KEY,
  notification_time TIME,
  monday_notification BOOLEAN DEFAULT FALSE,
  tuesday_notification BOOLEAN DEFAULT FALSE,
  wednesday_notification BOOLEAN DEFAULT FALSE,
  thursday_notification BOOLEAN DEFAULT FALSE,
  friday_notification BOOLEAN DEFAULT FALSE,
  saturday_notification BOOLEAN DEFAULT FALSE,
  sunday_notification BOOLEAN DEFAULT FALSE,
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

