CREATE TABLE IF NOT EXISTS todos (
    id INTEGER,
    name TEXT NOT NULL,
    description TEXT NOT NULL,
    status TEXT DEFAULT 'to do',
    created_at DATE DEFAULT CURRENT_TIMESTAMP,
    last_modified_date DATE DEFAULT CURRENT_TIMESTAMP,
    active BOOLEAN DEFAULT TRUE,
    CONSTRAINT pk_todos PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS users (
    id INTEGER,
    first_name TEXT NOT NULL,
    last_name TEXT NOT NULL,
    role TEXT DEFAULT 'user',
    email TEXT NOT NULL UNIQUE,
    password TEXT NOT NULL,
    active BOOLEAN DEFAULT TRUE,
    created_at DATE DEFAULT CURRENT_TIMESTAMP,
    last_modified_date DATE DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT pk_users PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS todos_users (
    todo_id INTEGER,
    user_id INTEGER,
    CONSTRAINT fk_todo FOREIGN KEY (todo_id) REFERENCES todos (id),
    CONSTRAINT fk_user FOREIGN KEY (user_id) REFERENCES users (id)
);
