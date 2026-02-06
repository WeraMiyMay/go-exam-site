-- Категории тестирования (рабочий, технолог и т.д.)
CREATE TABLE categories (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT NOT NULL,
    description TEXT
);

-- Темы
CREATE TABLE topics (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    category_id INTEGER NOT NULL,
    name TEXT NOT NULL,
    description TEXT,
    FOREIGN KEY (category_id) REFERENCES categories(id)
);

-- Вопросы
CREATE TABLE questions (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    topic_id INTEGER NOT NULL,
    text TEXT NOT NULL,
    points INTEGER DEFAULT 1,
    FOREIGN KEY (topic_id) REFERENCES topics(id)
);

-- Варианты ответов
CREATE TABLE answers (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    question_id INTEGER NOT NULL,
    text TEXT NOT NULL,
    is_correct BOOLEAN NOT NULL DEFAULT 0,
    FOREIGN KEY (question_id) REFERENCES questions(id)
);

-- Попытки экзамена (без регистрации)
CREATE TABLE exam_attempts (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    full_name TEXT NOT NULL,
    employee_id TEXT NOT NULL,
    position TEXT,
    category_id INTEGER NOT NULL,
    score INTEGER DEFAULT 0,
    max_score INTEGER DEFAULT 0,
    passed BOOLEAN DEFAULT 0,
    started_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    finished_at DATETIME,
    FOREIGN KEY (category_id) REFERENCES categories(id)
);

-- Ответы пользователя
CREATE TABLE exam_answers (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    attempt_id INTEGER NOT NULL,
    question_id INTEGER NOT NULL,
    answer_id INTEGER NOT NULL,
    is_correct BOOLEAN NOT NULL,
    FOREIGN KEY (attempt_id) REFERENCES exam_attempts(id),
    FOREIGN KEY (question_id) REFERENCES questions(id),
    FOREIGN KEY (answer_id) REFERENCES answers(id)
);
