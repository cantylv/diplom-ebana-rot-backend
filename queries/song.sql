-- Эта таблица содержит данные о песнях
CREATE TABLE song (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    release_date DATE,
    text JSONB,
    link TEXT,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT now(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT now()
);

ALTER TABLE song
ALTER COLUMN release_date SET NOT NULL,
ALTER COLUMN text SET NOT NULL,
ALTER COLUMN link SET NOT NULL;