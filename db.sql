CREATE TABLE IF NOT EXISTS posts(
    id int AUTO_INCREMENT PRIMARY KEY,
    user_id int NOT NULL,
    post_title VARCHAR(250) NOT NULL,
    post_content text NOT NULL,
    post_hashtags text NOT NULL,
    created_at timestamp default current_timestamp,
    updated_at timestamp default current_timestamp,
    created_by text NOT NULL,
    updated_by text NOT NULL
);

ALTER TABLE posts MODIFY user_id bigint;
ALTER TABLE posts ADD CONSTRAINT fk_user_id_posts FOREIGN KEY(user_id)  REFERENCES users(id);