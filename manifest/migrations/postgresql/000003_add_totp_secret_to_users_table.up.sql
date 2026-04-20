ALTER TABLE users
    ADD COLUMN totp_secret varchar(64) NOT NULL DEFAULT '';