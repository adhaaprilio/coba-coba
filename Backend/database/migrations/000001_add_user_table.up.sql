CREATE SEQUENCE user_id_seq;

CREATE TABLE IF NOT EXISTS users (
    user_id  VARCHAR(255) PRIMARY KEY DEFAULT nextval('user_id_seq'),
    name VARCHAR(255) NOT NULL,
    username VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL,
    password VARCHAR(255) NOT NULL
);