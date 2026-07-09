-- +goose Up
CREATE TABLE course(
    id serial primary key,
    name VARCHAR(255),
    description TEXT,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE chapter(
    id serial primary key,
    name VARCHAR(255),
    description TEXT,
    chapter_order INTEGER,
    course_id INTEGER,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (course_id) REFERENCES course(id) ON DELETE CASCADE
);

CREATE TABLE lesson(
    id serial primary key,
    name VARCHAR(255),
    description TEXT,
    content TEXT,
    lesson_order INTEGER,
    chapter_id INTEGER,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (chapter_id) REFERENCES chapter(id) ON DELETE CASCADE
);

-- +goose Down
DROP TABLE IF EXISTS lesson CASCADE;
DROP TABLE IF EXISTS chapter CASCADE;
DROP TABLE IF EXISTS course CASCADE;