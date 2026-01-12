CREATE TABLE roles
(
    id   smallserial PRIMARY KEY,
    name varchar(20) NOT NULL UNIQUE
);