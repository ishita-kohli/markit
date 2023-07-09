CREATE TABLE IF NOT EXISTS documents (
    id bigserial PRIMARY KEY,
    title varchar NOT NULL,
    body TEXT NOT NULL DEFAULT '',
    created_at timestamptz NOT NULL DEFAULT now(),
    updated_at timestamptz NOT NULL DEFAULT now()
);

CREATE TYPE document_access_roles AS ENUM ('owner', 'editor', 'viewer');

CREATE TABLE IF NOT EXISTS document_access (
    document_id bigserial NOT NULL,
    user_id bigserial NOT NULL,
    role document_access_roles NOT NULL,
    PRIMARY KEY(document_id, user_id)
);