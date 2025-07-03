-- +goose Up
CREATE TABLE IF NOT EXISTS users (
    id BIGINT PRIMARY KEY
);

CREATE TABLE IF NOT EXISTS segments (
    name TEXT PRIMARY KEY
);

CREATE TABLE IF NOT EXISTS user_segments (
    user_id BIGINT REFERENCES users(id) ON DELETE CASCADE,
    segment_name TEXT REFERENCES segments(name) ON UPDATE CASCADE ON DELETE CASCADE,
    PRIMARY KEY (user_id, segment_name)
);

-- +goose Down
DROP TABLE IF EXISTS user_segments;
DROP TABLE IF EXISTS segments;
DROP TABLE IF EXISTS users; 