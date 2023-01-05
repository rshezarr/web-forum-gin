-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS users (
	id SERIAL PRIMARY KEY,
	email VARCHAR(255) UNIQUE NOT NULL,
	username VARCHAR(255) UNIQUE NOT NULL,
	password VARCHAR(255) NOT NULL,
	posts INTEGER DEFAULT 0
);

CREATE TABLE IF NOT EXISTS posts (
	id SERIAL PRIMARY KEY,
	user_id INTEGER NOT NULL,
	title VARCHAR(100) NOT NULL,
	content TEXT NOT NULL,
	creation_time TIMESTAMP NOT NULL,
	likes INTEGER DEFAULT 0,
	dislikes INTEGER DEFAULT 0,
	FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS commentaries (
    id SERIAL PRIMARY KEY,
	user_id INTEGER NOT NULL,
	post_id INTEGER NOT NULL,
	content TEXT NOT NULL,
	likes INTEGER DEFAULT 0,
	dislikes INTEGER DEFAULT 0,
	FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
	FOREIGN KEY (post_id) REFERENCES posts(id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS votes (
    user_id INTEGER NOT NULL,
	content_type VARCHAR(10) NOT NULL,
  	content_id INTEGER NOT NULL,
  	is_like BOOLEAN NOT NULL DEFAULT FALSE,
  	is_dislike BOOLEAN NOT NULL DEFAULT FALSE,
  	PRIMARY KEY (user_id, content_type, content_id)
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
