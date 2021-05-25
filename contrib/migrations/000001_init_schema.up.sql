CREATE TABLE users (
    id BIGSERIAL PRIMARY KEY,
    uuid TEXT NOT NULL,
    role text NOT NULL,
    email TEXT NOT NULL,
    password TEXT NOT NULL,
    created_at TIMESTAMP WITHOUT TIME ZONE NOT NULL,
    last_updated_at TIMESTAMP WITHOUT TIME ZONE NOT NULL
);

INSERT INTO users  (uuid, role, email, password, created_at, last_updated_at) values ('180338c8-53a9-4305-6f40-c8e437a3806a', 'admin','admin@gmail.com', 'WmNxMMGGEH4S-1XFC_xFTNXHdtE=', now(), now())
