CREATE SCHEMA IF NOT EXISTS delegacia_facil;

CREATE TABLE IF NOT EXISTS delegacia_facil.users (
                                                    id UUID PRIMARY KEY,
                                                    first_name VARCHAR(255) NOT NULL UNIQUE,
                                                    last_name VARCHAR(255) NOT NULL UNIQUE,
                                                    email VARCHAR(255) NOT NULL UNIQUE,
                                                    password_hash VARCHAR(255) NOT NULL,
                                                    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
                                                    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE TABLE delegacia_facil.delegacia (
                           id UUID PRIMARY KEY,
                           nome VARCHAR(255) NOT NULL,
                           endereco VARCHAR(255),
                           horario24h BOOLEAN,
                           latitude DOUBLE PRECISION,
                           longitude DOUBLE PRECISION
);