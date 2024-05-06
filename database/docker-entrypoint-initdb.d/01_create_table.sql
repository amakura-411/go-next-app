-- ユーザー
CREATE TABLE Users (
  user_id VARCHAR(255) PRIMARY KEY,
  username VARCHAR(255) NOT NULL,
  password VARCHAR(255) NOT NULL,
  icon VARCHAR(255),
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- カテゴリー
CREATE TABLE Categories (
  category_id VARCHAR(255) PRIMARY KEY,
  category_name VARCHAR(255) NOT NULL,
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- プロジェクト
CREATE TABLE Projects (
  project_id VARCHAR(255) PRIMARY KEY,
  assigned_user_id VARCHAR(255) NOT NULL,
  project_title VARCHAR(255) NOT NULL,
  project_description TEXT,
  goal_date TIMESTAMP NOT NULL,
  category_id VARCHAR(255),
  display_flag BOOLEAN DEFAULT TRUE,
  task_counts INTEGER NOT NULL DEFAULT 0,
  task_done_counts INTEGER NOT NULL DEFAULT 0,
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  FOREIGN KEY (assigned_user_id) REFERENCES Users(user_id),
  FOREIGN KEY (category_id) REFERENCES Categories(category_id)
);

-- プロジェクトとユーザーの関係
CREATE TABLE Projects_Users (
  project_user_id VARCHAR(255) PRIMARY KEY,
  project_id VARCHAR(255) NOT NULL,
  user_id VARCHAR(255) NOT NULL,
  FOREIGN KEY (project_id) REFERENCES Projects(project_id),
  FOREIGN KEY (user_id) REFERENCES Users(user_id)
);

-- タスクステータス
CREATE TABLE TaskStatus (
  status_id INT PRIMARY KEY,
  status_name VARCHAR(50) NOT NULL
);

-- タスク優先度
CREATE TABLE TaskPriority (
  priority_id INT PRIMARY KEY,
  priority_name VARCHAR(50) NOT NULL
);

-- タスク
CREATE TABLE Tasks (
  task_id VARCHAR(255) PRIMARY KEY,
  project_id VARCHAR(255) NOT NULL,
  user_id VARCHAR(255) NOT NULL,
  task_title VARCHAR(255) NOT NULL,
  task_description TEXT,
  task_type ENUM('date', 'time') NOT NULL,
  start_datetime TIMESTAMP,
  end_datetime TIMESTAMP,
  deadline TIMESTAMP NOT NULL,
  status_id INT NOT NULL,
  priority_id INT NOT NULL,
  progress INTEGER NOT NULL DEFAULT 0,
  estimated_time INTEGER NOT NULL DEFAULT 0,
  actual_time INTEGER NOT NULL DEFAULT 0,
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  FOREIGN KEY (project_id) REFERENCES Projects(project_id),
  FOREIGN KEY (user_id) REFERENCES Users(user_id),
  FOREIGN KEY (status_id) REFERENCES TaskStatus(status_id),
  FOREIGN KEY (priority_id) REFERENCES TaskPriority(priority_id)
);


-- プロジェクトとタスクの関係
CREATE TABLE Projects_Tasks (
  project_task_id VARCHAR(255) PRIMARY KEY,
  project_id VARCHAR(255) NOT NULL,
  task_id VARCHAR(255) NOT NULL,
  FOREIGN KEY (project_id) REFERENCES Projects(project_id),
  FOREIGN KEY (task_id) REFERENCES Tasks(task_id)
);



-- タスクとコメントの関係
CREATE TABLE Tasks_Comments (
  task_comment_id VARCHAR(255) PRIMARY KEY,
  task_id VARCHAR(255) NOT NULL,
  user_id VARCHAR(255) NOT NULL,
  content TEXT NOT NULL,
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  FOREIGN KEY (task_id) REFERENCES Tasks(task_id),
  FOREIGN KEY (user_id) REFERENCES Users(user_id)
);


CREATE TABLE Schedules(
  schedule_id VARCHAR(255) PRIMARY KEY,
  schedule_title VARCHAR(255) NOT NULL,
  start_time TIMESTAMP NOT NULL,
  end_time TIMESTAMP NOT NULL,
  start_date TIMESTAMP NOT NULL,
  end_date TIMESTAMP NOT NULL,
  priority ENUM('low', 'middle', 'high') DEFAULT 'middle',
  done_flag BOOLEAN DEFAULT FALSE
);


-- タスクと通知の関係
CREATE TABLE Tasks_Notifications (
  task_notification_id VARCHAR(255) PRIMARY KEY,
  task_id VARCHAR(255) NOT NULL,
  notification_time TIME,
  monday_notification BOOLEAN DEFAULT FALSE,
  tuesday_notification BOOLEAN DEFAULT FALSE,
  wednesday_notification BOOLEAN DEFAULT FALSE,
  thursday_notification BOOLEAN DEFAULT FALSE,
  friday_notification BOOLEAN DEFAULT FALSE,
  saturday_notification BOOLEAN DEFAULT FALSE,
  sunday_notification BOOLEAN DEFAULT FALSE,
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  FOREIGN KEY (task_id) REFERENCES Tasks(task_id)
);



-- ユーザーが１日にどれくらいのタスクを行い、完了したか
CREATE TABLE User_Tasks_Daily (
  user_task_daily_id VARCHAR(255) PRIMARY KEY,
  user_id VARCHAR(255) NOT NULL,
  date TIMESTAMP NOT NULL,
  task_counts INTEGER NOT NULL DEFAULT 0,
  task_done_counts INTEGER NOT NULL DEFAULT 0,
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  FOREIGN KEY (user_id) REFERENCES Users(user_id)
);

-- ユーザーが１週間にどれくらいのタスクを行い、完了したか
CREATE TABLE User_Tasks_Weekly (
  user_task_weekly_id VARCHAR(255) PRIMARY KEY,
  user_id VARCHAR(255) NOT NULL,
  week_start_date TIMESTAMP NOT NULL,
  week_end_date TIMESTAMP NOT NULL,
  task_counts INTEGER NOT NULL DEFAULT 0,
  task_done_counts INTEGER NOT NULL DEFAULT 0,
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  FOREIGN KEY (user_id) REFERENCES Users(user_id)
);

-- ユーザーが１ヶ月にどれくらいのタスクを行い、完了したか
CREATE TABLE User_Tasks_Monthly (
  user_task_monthly_id VARCHAR(255) PRIMARY KEY,
  user_id VARCHAR(255) NOT NULL,
  month TIMESTAMP NOT NULL,
  task_counts INTEGER NOT NULL DEFAULT 0,
  task_done_counts INTEGER NOT NULL DEFAULT 0,
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  FOREIGN KEY (user_id) REFERENCES Users(user_id)
);
