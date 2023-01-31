CREATE TABLE IF NOT EXISTS users(
    id uuid,
    name varchar(255),
    username varchar(255),
    email varchar(255),
    password varchar(255),
    constraint pk_users PRIMARY KEY (id, email) --
);

CREATE TABLE IF NOT EXISTS wallet (
    wallet_id uuid PRIMARY KEY,
    user_id uuid,
    amount float
)