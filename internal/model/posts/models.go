package posts

import "time"

type (
	CreatePostRequest struct {
		PostTitle    string   `json:"postTitle"`
		PostContent  string   `json:"postContent"`
		PostHashtags []string `json:"postHashtags"`
	}

	CreateCommentRequest struct {
		CommentContent string `json:"commentContent"`
	}

	UserActivityRequest struct {
		IsLiked bool `json:"isLiked"`
	}
)

type (
	PostModel struct {
		ID           int64     `db:"id"`
		UserID       int64     `db:"user_id"`
		PostTitle    string    `db:"post_title"`
		PostContent  string    `db:"post_content"`
		PostHashtags string    `db:"post_hashtags"`
		CreatedAt    time.Time `db:"created_at"`
		UpdatedAt    time.Time `db:"updated_at"`
		CreatedBy    string    `db:"created_by"`
		UpdatedBy    string    `db:"updated_by"`
	}

	CommentModel struct {
		ID             int64     `db:"id"`
		PostID         int64     `db:"post_id"`
		UserID         int64     `db:"user_id"`
		CommentContent string    `db:"comment_content"`
		CreatedAt      time.Time `db:"created_at"`
		UpdatedAt      time.Time `db:"updated_at"`
		CreatedBy      string    `db:"created_by"`
		UpdatedBy      string    `db:"updated_by"`
	}

	UserActivityModel struct {
		ID int64 `db:"id"`
		PostID int64 `db:"post_id"`
		UserID int64 `db:"user_id"`
		IsLiked bool `db:"is_liked"`
		CreatedAt      time.Time `db:"created_at"`
		UpdatedAt      time.Time `db:"updated_at"`
		CreatedBy      string    `db:"created_by"`
		UpdatedBy      string    `db:"updated_by"`
	}
)
