

CREATE DATABASE fineasy;

CREATE TYPE fluxType enum('expense', 'income');



-- CREATE TABLE sources(
--     id serial primary key,
--     flux fluxType not null,
--     title varchar(30) not null
-- );

-- CREATE TABLE wallets(
--     id serial primary key,
--     title varchar(30) not null
-- );

-- CREATE TABLE categories(
--     id serial primary key,
--     title varchar(30) not null
-- );

-- CREATE TABLE flows (
--     id serial primary key,
--     flux fluxType not null,
--     cash numeric(8, 2) not null,
--     source references source(id) not null,
--     dest references wallet(id) not null,
--     category references category(id) not null,
--     createdAt date not null
-- );

