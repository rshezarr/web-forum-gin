-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS users (
	id SERIAL PRIMARY KEY,
	email VARCHAR(255) UNIQUE NOT NULL,
	username VARCHAR(255) UNIQUE NOT NULL,
	password VARCHAR(255) NOT NULL,
	posts INTEGER DEFAULT 0,
	token TEXT DEFAULT NULL,
    expiration_time TIMESTAMP 
);

CREATE TABLE IF NOT EXISTS posts (
	id SERIAL PRIMARY KEY,
	user_id INTEGER NOT NULL,
	title VARCHAR(100),
	content TEXT,
	creation_time TIMESTAMP NOT NULL,
	likes INTEGER DEFAULT 0,
	dislikes INTEGER DEFAULT 0,
	FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS post_category (
    post_id INTEGER NOT NULL,
	category VARCHAR(255),
	FOREIGN KEY (post_id) REFERENCES posts(id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS commentaries (
    id SERIAL PRIMARY KEY,
	user_id INTEGER NOT NULL,
	post_id INTEGER NOT NULL,
	content TEXT NOT NULL,
	likes INT DEFAULT 0,
	dislikes INT DEFAULT 0,
	FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
	FOREIGN KEY (post_id) REFERENCES posts(id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS likes (
    user_id INTEGER NOT NULL,
	post_id INTEGER DEFAULT NULL,
	commentary_id INTEGER DEFAULT NULL,
	like BOOLEAN NOT NULL,
	FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
	FOREIGN KEY (post_id) REFERENCES posts(id) ON DELETE CASCADE,
	FOREIGN KEY (commentary_id) REFERENCES commentaries(id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS dislikes (
    user_id INTEGER NOT NULL,
	post_id INTEGER DEFAULT NULL,
	commentary_id INTEGER DEFAULT NULL,
	dislike BOOLEAN NOT NULL,
	FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
	FOREIGN KEY (post_id) REFERENCES posts(id) ON DELETE CASCADE,
	FOREIGN KEY (commentary_id) REFERENCES commentaries(id) ON DELETE CASCADE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE likes;

DROP TABLE dislikes;

DROP TABLE commentaries;

DROP TABLE post_category;

DROP TABLE posts;

DROP TABLE users;
-- +goose StatementEnd
