CREATE TABLE IF NOT EXISTS topic
(
    id          serial                                             NOT NULL PRIMARY KEY,
    name        varchar(64)                                        NOT NULL,
    language_id int                                                NOT NULL REFERENCES language (id),
    created_at  timestamp WITH TIME ZONE DEFAULT current_timestamp NOT NULL
);
