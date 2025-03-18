package posts

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/azybk/mini-forum/internal/model/posts"
	"github.com/gin-gonic/gin"
)

func (h *Handler) UpsertUserActivity(c *gin.Context) {
	ctx := c.Request.Context()

	var request posts.UserActivityRequest
	if err:=c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	postIDstr := c.Param("postID")
	postID, err := strconv.ParseInt(postIDstr, 10, 64)
	
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": errors.New("invalid Post ID").Error(),
		})
		return
	}

	userID := c.GetInt64("userID")
	err = h.postSvc.UpsertUserActivity(ctx, postID, userID, request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.Status(http.StatusOK)
}