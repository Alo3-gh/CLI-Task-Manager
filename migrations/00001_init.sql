-- +goose Up
-- +goose StatementBegin
CREATE TABLE tasks (
    id          BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    title       TEXT NOT NULL,
    description TEXT NOT NULL DEFAULT '',
    status      TEXT,
    created_at  TIMESTAMP NOT NULL DEFAULT now(),
    deadline    TIMESTAMP
);

COMMENT ON TABLE tasks IS 'Таблица задач';
COMMENT ON COLUMN tasks.id IS 'ID задачи';
COMMENT ON COLUMN tasks.title IS 'Заголовок задачи';
COMMENT ON COLUMN tasks.description IS 'Описание задачи';
COMMENT ON COLUMN tasks.status IS 'Статус задачи';
COMMENT ON COLUMN tasks.created_at IS 'Дата создания';
COMMENT ON COLUMN tasks.deadline IS 'Дата окончания';
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE tasks;
-- +goose StatementEnd
