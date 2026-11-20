-- +goose Up
-- +goose StatementBegin
CREATE TABLE tasks (
    id          integer PRIMARY KEY,
    title       TEXT NOT NULL,
    description TEXT,
    status      TEXT NOT NULL,
    created_at  TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now(),
    deadline    TIMESTAMP WITH TIME ZONE
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
