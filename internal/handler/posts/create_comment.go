package posts

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/azybk/mini-forum/internal/model/posts"
	"github.com/gin-gonic/gin"
)

func (h *Handler) CreateComment(c *gin.Context) {
	ctx := c.Request.Context()

	var request posts.CreateCommentRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	postIDStr := c.Param("postID")
	postID, err := strconv.ParseInt(postIDStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": errors.New("invalid Post ID").Error(),
		})
		return
	}

	userID := c.GetInt64("userID")

	err = h.postSvc.CreateComment(ctx, postID, userID, request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": errors.New("Post ID not found").Error(),
		})
		return
	}

	c.Status(http.StatusCreated)
}
