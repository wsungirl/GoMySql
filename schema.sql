CREATE TABLE users (
	id bigint NOT NULL AUTO_INCREMENT,
	name varchar(128) NOT NULL UNIQUE,
	password_hash TEXT NOT NULL,
	PRIMARY KEY (id)
);

CREATE TABLE dbs (
	id BIGINT NOT NULL AUTO_INCREMENT,
	user_id bigint NOT NULL,
	db_name varchar(128) NOT NULL,
	PRIMARY KEY (id),
  FOREIGN KEY (user_id) REFERENCES users(id),
  UNIQUE (user_id, db_name)
);

CREATE TABLE permissions (
	db_id BIGINT NOT NULL,
  user_id BIGINT NOT NULL,
	action VARCHAR(16) NOT NULL,
	entity VARCHAR(16) NOT NULL,
  FOREIGN KEY (db_id) REFERENCES dbs(id),
  FOREIGN KEY (user_id) REFERENCES users(id)
);

CREATE TABLE history (
	db_id bigint NOT NULL,
  user_id BIGINT NOT NULL,
	action varchar(16) NOT NULL,
	entity TEXT NOT NULL,
  FOREIGN KEY (db_id) REFERENCES dbs(id),
  FOREIGN KEY (user_id) REFERENCES users(id)
);

CREATE TABLE sessions (
	user_id bigint NOT NULL,
	access_token TEXT NOT NULL,
	FOREIGN KEY (user_id) REFERENCES users(id)
);

