CREATE TABLE users
(
    id            serial PRIMARY KEY,
    username      varchar(40)  NOT NULL UNIQUE,
    email         varchar(255) NOT NULL UNIQUE,
    password      varchar(128) NOT NULL,
    phone_country varchar(5),
    phone_number  varchar(20),
    role_id       smallint     NOT NULL,
    avatar        text,
    created_at    timestamptz  NOT NULL DEFAULT now(),
    updated_at    timestamptz  NOT NULL DEFAULT now(),
    deleted_at    timestamptz
);