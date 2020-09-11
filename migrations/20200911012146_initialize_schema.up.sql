CREATE TABLE IF NOT EXISTS users(
    id serial PRIMARY KEY,
    hash VARCHAR (50) UNIQUE NOT NULL,
    nickname VARCHAR (50) UNIQUE NOT NULL,
    password VARCHAR (50) NOT NULL,
    email VARCHAR (100) UNIQUE NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
    deleted_at TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE UNIQUE INDEX users_hash_idx ON users(hash);
CREATE UNIQUE INDEX users_nickname_idx ON users(nickname);
CREATE UNIQUE INDEX users_email_idx ON users(email);

CREATE TABLE IF NOT EXISTS currencies(
    id serial PRIMARY KEY,
    hash VARCHAR (50) UNIQUE NOT NULL,
    name VARCHAR (50) UNIQUE NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
    deleted_at TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE UNIQUE INDEX currencies_hash_idx ON currencies(hash);

CREATE TABLE IF NOT EXISTS balances(
    id serial PRIMARY KEY,
    user_id INTEGER,
    currency_id INTEGER,
    balance FLOAT,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
    deleted_at TIMESTAMP NOT NULL DEFAULT NOW(),
    CONSTRAINT fk_user FOREIGN KEY(user_id) REFERENCES users(id),
    CONSTRAINT fk_currency FOREIGN KEY(currency_id) REFERENCES currencies(id)
);
