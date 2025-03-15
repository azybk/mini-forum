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

CREATE TABLE IF NOT EXISTS comments(
    id int AUTO_INCREMENT PRIMARY key,
    post_id int not NULL,
    user_id bigint not null,
    comment_content text not null,
    created_at timestamp default current_timestamp,
    updated_at timestamp default current_timestamp,
    created_by text NOT NULL,
    updated_by text NOT NULL,
    CONSTRAINT fk_post_id_comments FOREIGN KEY (post_id) REFERENCES posts(id),
    CONSTRAINT fk_user_id_comments FOREIGN KEY (user_id) REFERENCES users(id)
);