package controllers

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/soundsnick/arzamas/helpers"
	"github.com/soundsnick/arzamas/models"
)

// PostIndex returns all posts
func PostIndex(c *gin.Context) {
	posts := models.GetPostsAll()
	c.JSON(200, gin.H{
		"data": posts,
	})
}

// PostGet returns post by id
func PostGet(c *gin.Context) {
	ID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	post := models.GetPostByID(ID)
	if post.ID == 0 || err != nil {
		c.JSON(200, gin.H{
			"error": "not found",
		})
	} else {
		c.JSON(200, gin.H{
			"data": post,
		})
	}
}

// PostSearch searches posts by query
func PostSearch(c *gin.Context) {
	query := c.Query("query")
	if query == "" {
		c.JSON(422, gin.H{
			"error": "'query' required",
		})
	} else {
		posts := models.GetPostsByTitle(query)
		c.JSON(200, gin.H{
			"data": posts,
		})
	}
}

// PostUser user's posts
func PostUser(c *gin.Context) {
	userID, err := strconv.ParseUint(c.Param("user_id"), 10, 64)
	if err != nil {
		c.JSON(422, gin.H{
			"error": "wrong user_id",
		})
	} else {
		c.JSON(200, gin.H{
			"data": models.GetPostsByUserID(userID),
		})
	}
}

// PostCreate creates post
func PostCreate(c *gin.Context) {
	form := helpers.PostForm{
		Title:   c.Query("title"),
		Content: c.Query("content"),
		Cover:   c.Query("cover"),
		Token:   c.Query("token"),
	}

	// Check validation
	validatedField, validateErr := helpers.ValidatePostForm(form)
	if validateErr != nil {
		c.JSON(422, gin.H{
			"error": validateErr,
			"field": validatedField,
		})
	} else {
		user := models.GetUserByToken(form.Token)
		if user.ID == 0 {
			c.JSON(422, gin.H{
				"error": "unauthorised: wrong token",
			})
		} else {
			post := models.Post{
				Title:   form.Title,
				Content: form.Content,
				Cover:   form.Cover,
				UserID:  user.ID,
				User:    user,
			}
			models.GetDB().Create(&post)
			c.JSON(200, gin.H{
				"data": post,
			})
		}
	}
}
