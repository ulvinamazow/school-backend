CREATE TABLE IF NOT EXISTS users {
    id serial PRIMARY KEY NOT NULL,
    name character varying(45) NOT NULL,
    surname character varying(45) NOT NULL,
    username character varying(45) NOT NULL UNIQUE,
    email   character varying(45) NOT NULL UNIQUE,
    role_id INTEGER REFERENCES roles(id),
};
CREATE TABLE IF NOT EXISTS roles {
    id serial PRIMARY KEY NOT NULL,
    role_name character varying(45) NOT NULL UNIQUE
};