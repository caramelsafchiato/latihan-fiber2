CREATE TABLE IF NOT EXISTS students (
    id          SERIAL          PRIMARY KEY,
    username    VARCHAR(50)     NOT NULL,
    email       VARCHAR(255)    NOT NULL,
    password    VARCHAR(255)    NOT NULL,
    is_active   BOOLEAN         NOT NULL DEFAULT TRUE,
    created_at  TIMESTAMPTZ     NOT NULL DEFAULT NOW()
);

-- Keunikan username tanpa membedakan huruf besar dan kecil.
-- Inilah yang menggantikan pemeriksaan manual di pertemuan 2.
CREATE UNIQUE INDEX IF NOT EXISTS studentds_username_lower_key
ON students (LOWER(username));

CREATE INDEX IF NOT EXISTS students_email_lower_idx
ON students (LOWER(email));