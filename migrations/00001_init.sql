-- +goose Up
SELECT 'up SQL query';

CREATE TABLE IF NOT EXISTS tasks (
    id integer PRIMARY KEY,
    title varchar(40) NOT NULL,
    description varchar(160),
    status varchar(40) NOT NULL,
    created_at date NOT NULL,
    deadline date
);

-- +goose Down
SELECT 'down SQL query';

DROP TABLE IF EXISTS tasks CASCADE;
