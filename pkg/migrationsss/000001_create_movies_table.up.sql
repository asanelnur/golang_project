CREATE TABLE IF NOT EXISTS categories (
    id          bigserial PRIMARY KEY,
    created_at  timestamp(0) with time zone NOT NULL DEFAULT NOW(),
    updated_at  timestamp(0) with time zone NOT NULL DEFAULT NOW(),
    title       text                        NOT NULL,
    description text                        NOT NULL
);

CREATE TABLE IF NOT EXISTS authors (
    id              bigserial PRIMARY KEY,
    created_at      timestamp(0) with time zone NOT NULL DEFAULT NOW(),
    updated_at      timestamp(0) with time zone NOT NULL DEFAULT NOW(),
    name           text                        NOT NULL,
    info           text,
    age            int
);

CREATE TABLE IF NOT EXISTS books (
    id              bigserial PRIMARY KEY,
    created_at      timestamp(0) with time zone NOT NULL DEFAULT NOW(),
    updated_at      timestamp(0) with time zone NOT NULL DEFAULT NOW(),
    title           text                        NOT NULL,
    description     text,
    price           int,
    category        bigserial,
    author          bigserial,
    FOREIGN KEY (category)
        REFERENCES categories(id),
    FOREIGN KEY (author)
        REFERENCES authors(id)
); 