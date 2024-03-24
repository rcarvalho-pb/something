CREATE TRIGGER IF NOT EXISTS update_todo_last_modified_date
    AFTER UPDATE
    ON todos
    BEGIN
        UPDATE todos SET last_modified_date = CURRENT_TIMESTAMP WHERE ROWID = new.ROWID;
    END;

CREATE TRIGGER IF NOT EXISTS update_user_last_modified_date
    AFTER UPDATE
    ON users
    BEGIN
        UPDATE users SET last_modified_date = CURRENT_TIMESTAMP WHERE ROWID = new.ROWID;
    END;
