-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS tasks
(
    id    SERIAL PRIMARY KEY,
    name  TEXT NOT NULL ,
    description  TEXT,
    deadline   TIMESTAMP,
    priority   INT
);

CREATE TABLE IF NOT EXISTS comments
(
    id INT PRIMARY KEY ,
    comment TEXT
);

CREATE INDEX idx_tasks_deadline ON tasks (deadline);
CREATE INDEX idx_tasks_priority ON tasks (priority);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS tasks;
DROP TABLE IF EXISTS comments;
-- +goose StatementEnd
