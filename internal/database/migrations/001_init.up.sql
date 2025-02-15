CREATE TABLE IF NOT EXISTS users (
	id              serial4      NOT NULL,
	username        varchar(255) UNIQUE NOT NULL,
	password_hash   varchar(255) NOT NULL,

	CONSTRAINT users_pkey PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS item (
	id serial4                    NOT NULL,
	"name"          varchar(20)   NOT NULL,
	amount          int4          NOT NULL,

    CONSTRAINT item_pkey PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS orders (
	id serial4                    NOT NULL,
	user_id         int4          NOT NULL,
	item_id         int4          NOT NULL,
	CONSTRAINT orders_pkey PRIMARY KEY (id)
);


CREATE TABLE NOT EXISTS  coinhistory (
	id              serial4       NOT NULL,
	fromuser_id     int4          NOT NULL,
	touser_id       int4          NOT NULL,
	amount          int4          NOT NULL,
	CONSTRAINT coinhistory_pkey PRIMARY KEY (id)
);

CREATE TABLE NOT EXISTS wallet (
	id              serial4       NOT NULL,
	user_id         int4          NOT NULL,
	amount          int4          NOT NULL,
	CONSTRAINT wallet_pkey PRIMARY KEY (id)
);