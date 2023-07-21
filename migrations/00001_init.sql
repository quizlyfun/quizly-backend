-- +goose Up
-- +goose StatementBegin
BEGIN
;

CREATE TABLE IF NOT EXISTS players (
    nickname varchar(25) PRIMARY KEY NOT NULL,
    email varchar(255) UNIQUE NOT NULL,
    display_name varchar(25),
    email_verified boolean DEFAULT FALSE NOT NULL,
    password text NOT NULL,
    created_at timestamptz NOT NULL,
    updated_at timestamptz
);

CREATE TABLE IF NOT EXISTS media (
    url varchar(2048) NOT NULL PRIMARY KEY,
    type smallint NOT NULL,
    uploaded_by varchar(25) NOT NULL REFERENCES players (nickname),
    created_at timestamptz NOT NULL
);

CREATE TABLE IF NOT EXISTS packages (
    id serial NOT NULL PRIMARY KEY,
    name varchar(64) NOT NULL,
    author varchar(25) NOT NULL REFERENCES players (nickname),
    is_published bool DEFAULT FALSE NOT NULL,
    cover_url varchar(2048) REFERENCES media (url),
    created_at timestamptz NOT NULL,
    updated_at timestamptz NOT NULL,
    round_count smallint,
    topic_count smallint,
    question_count smallint,
    video_count smallint,
    audio_count smallint,
    image_count smallint
);

CREATE TABLE IF NOT EXISTS rounds (
    id serial NOT NULL PRIMARY KEY,
    name varchar(32) NOT NULL,
    position smallint NOT NULL,
    package_id int NOT NULL REFERENCES packages (id)
);

CREATE TABLE IF NOT EXISTS topics (
    id serial NOT NULL PRIMARY KEY,
    title varchar(50) NOT NULL,
    author varchar(25) NOT NULL REFERENCES players (nickname),
    round_id int NOT NULL REFERENCES rounds (id),
    created_at timestamptz NOT NULL
);

CREATE TABLE IF NOT EXISTS answers (
    id serial NOT NULL PRIMARY KEY,
    text varchar(112) NOT NULL,
    author varchar(25) NOT NULL REFERENCES players (nickname),
    media_url varchar(2048) REFERENCES media (url),
    created_at timestamptz NOT NULL
);

CREATE TABLE IF NOT EXISTS questions (
    id serial NOT NULL PRIMARY KEY,
    text varchar(200) NOT NULL,
    answer_id int NOT NULL REFERENCES answers (id),
    author varchar(25) NOT NULL REFERENCES players (nickname),
    media_url varchar(2048) REFERENCES media (url),
    created_at timestamptz NOT NULL
);

CREATE TABLE IF NOT EXISTS round_topics (
    round_id int NOT NULL REFERENCES rounds (id),
    topic_id int NOT NULL REFERENCES topics (id),
    PRIMARY KEY (round_id, topic_id)
);

CREATE TABLE IF NOT EXISTS round_questions (
    id serial NOT NULL PRIMARY KEY,
    round_id int NOT NULL REFERENCES rounds (id),
    topic_id int NOT NULL REFERENCES topics (id),
    question_id int NOT NULL REFERENCES questions (id),
    question_type smallint NOT NULL,
    cost smallint NOT NULL,
    answer_time smallint NOT NULL,
    host_comment text,
    secret_topic varchar(64),
    secret_cost smallint,
    transfer_type smallint,
    is_keepable boolean
);

CREATE TABLE IF NOT EXISTS tags (
    name varchar(16) NOT NULL PRIMARY KEY,
    author varchar(25) NOT NULL REFERENCES players (nickname),
    created_at timestamptz NOT NULL
);

ALTER TABLE
    tags
ADD
    COLUMN ts tsvector GENERATED ALWAYS AS (
        setweight(to_tsvector('russian', coalesce(name, '')), 'A')
    ) STORED;

CREATE INDEX tags_gin_idx ON tags USING GIN (ts);

CREATE TABLE IF NOT EXISTS package_tags (
    package_id int NOT NULL REFERENCES packages (id),
    tag varchar(16) NOT NULL REFERENCES tags (name),
    PRIMARY KEY (package_id, tag)
);

COMMIT;

-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
BEGIN;

DROP TABLE IF EXISTS package_tags CASCADE;
DROP TABLE IF EXISTS round_questions CASCADE;
DROP TABLE IF EXISTS round_topics CASCADE;
DROP TABLE IF EXISTS questions CASCADE;
DROP TABLE IF EXISTS answers CASCADE;
DROP TABLE IF EXISTS topics CASCADE;
DROP TABLE IF EXISTS rounds CASCADE;
DROP TABLE IF EXISTS tags CASCADE;
DROP TABLE IF EXISTS packages CASCADE;
DROP TABLE IF EXISTS media CASCADE;
DROP TABLE IF EXISTS players CASCADE;

COMMIT;
-- +goose StatementEnd