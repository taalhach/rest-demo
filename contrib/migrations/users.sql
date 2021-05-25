CREATE TABLE users (
    id BIGSERIAL PRIMARY KEY,
    uuid TEXT NOT NULL,
    role text NOT NULL,
    email TEXT NOT NULL,
    password TEXT NOT NULL,
    created_at TIMESTAMP WITHOUT TIME ZONE NOT NULL,
    last_updated_at TIMESTAMP WITHOUT TIME ZONE NOT NULL
);