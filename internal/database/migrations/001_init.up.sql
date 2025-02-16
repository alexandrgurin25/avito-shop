CREATE TABLE IF NOT EXISTS users (
    id              serial4      NOT NULL,
    username        varchar(255) UNIQUE NOT NULL,
    password_hash   varchar(255) NOT NULL,

    CONSTRAINT users_pkey PRIMARY KEY (id)
);

CREATE UNIQUE INDEX IF NOT EXISTS username_index
ON users (username);

CREATE TABLE IF NOT EXISTS item (
    id              serial4      NOT NULL,
    name            varchar(20)   NOT NULL,
    amount          int4          NOT NULL,

    CONSTRAINT item_pkey PRIMARY KEY (id)
);

CREATE UNIQUE INDEX IF NOT EXISTS name_index
ON item (name);

CREATE TABLE IF NOT EXISTS orders (
    id              serial4      NOT NULL,
    user_id         int4          NOT NULL,
    item_id         int4          NOT NULL,
    
    CONSTRAINT orders_pkey PRIMARY KEY (id),
    CONSTRAINT fk_user FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE,
    CONSTRAINT fk_item FOREIGN KEY (item_id) REFERENCES item (id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS coinhistory (
    id              serial4      NOT NULL,
    fromuser_id     int4          NOT NULL,
    touser_id       int4          NOT NULL,
    amount          int4          NOT NULL,

    CONSTRAINT coinhistory_pkey PRIMARY KEY (id),
    CONSTRAINT fk_from_user FOREIGN KEY (fromuser_id) REFERENCES users (id) ON DELETE CASCADE,
    CONSTRAINT fk_to_user FOREIGN KEY (touser_id) REFERENCES users (id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS wallet (
    id              serial4      NOT NULL,
    user_id         int4          NOT NULL,
    amount          int4          NOT NULL,

    CONSTRAINT wallet_pkey PRIMARY KEY (id),
    CONSTRAINT fk_wallet_user FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE
);
