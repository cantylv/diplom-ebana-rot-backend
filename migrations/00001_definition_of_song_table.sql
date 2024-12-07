-- +goose Up
-- +goose StatementBegin
CREATE TABLE song (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name TEXT,
    release_date DATE,
    text JSONB,
    link TEXT,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT now(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT now()
);

ALTER TABLE song
ALTER COLUMN name SET NOT NULL,
ALTER COLUMN release_date SET NOT NULL,
ALTER COLUMN text SET NOT NULL,
ALTER COLUMN link SET NOT NULL;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE song;
-- +goose StatementEnd
