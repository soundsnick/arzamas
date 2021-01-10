package handlers

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/soundsnick/arzamas/core"
	"github.com/soundsnick/arzamas/post"
	"github.com/soundsnick/arzamas/session"
)

// PostIndex returns all posts
func PostIndex(c *gin.Context) {
	posts := post.GetAll()
	c.JSON(200, gin.H{
		"data": posts,
	})
}

// PostSearch searches posts by query
func PostSearch(c *gin.Context) {
	query := c.Query("query")
	if query == "" {
		c.JSON(422, gin.H{
			"error": "'query' required",
		})
	} else {
		posts := post.GetByTitle(query)
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
			"data": post.GetByUserID(userID),
		})
	}
}

// PostCreate creates post
func PostCreate(c *gin.Context) {
	form := post.CreationForm{
		Title:   c.Query("title"),
		Content: c.Query("content"),
		Cover:   c.Query("cover"),
		Token:   c.Query("token"),
	}

	// Check validation
	validatedField, validateErr := post.ValidateCreationForm(form)
	if validateErr != nil {
		c.JSON(422, gin.H{
			"error": validateErr,
			"field": validatedField,
		})
	} else {
		user := session.GetUserByToken(form.Token)
		if user.ID == 0 {
			c.JSON(422, gin.H{
				"error": "unauthorised: wrong token",
			})
		} else {
			post := post.Post{
				Title:   form.Title,
				Content: form.Content,
				Cover:   form.Cover,
				UserID:  user.ID,
				User:    user,
			}
			core.GetDB().Create(&post)
			c.JSON(200, gin.H{
				"data": post,
			})
		}
	}
}

// PostRead returns post by id
func PostRead(c *gin.Context) {
	ID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	post := post.GetByID(ID)
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

// PostUpdate updates post
func PostUpdate(c *gin.Context) {
	ID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	postFound := post.GetByID(ID)
	form := post.CreationForm{
		Title:   c.Query("title"),
		Content: c.Query("content"),
		Cover:   c.Query("cover"),
		Token:   c.Query("token"),
	}
	if err != nil || postFound.ID == 0 {
		c.JSON(422, gin.H{
			"error": "wrong id",
		})
	} else {
		// Check validation
		validatedField, validateErr := post.ValidateUpdateForm(&form, postFound)
		if validateErr != nil {
			c.JSON(422, gin.H{
				"error": validateErr,
				"field": validatedField,
			})
		} else {
			user := session.GetUserByToken(form.Token)

			if user.ID == 0 {
				c.JSON(422, gin.H{
					"error": "unauthorised: wrong token",
				})
			} else {
				postFound.Title = form.Title
				postFound.Content = form.Content
				postFound.Cover = form.Cover
				core.GetDB().Save(&postFound)
				c.JSON(200, gin.H{
					"data": form,
				})
			}
		}
	}
}

// PostDelete deletes post
func PostDelete(c *gin.Context) {
	ID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(422, gin.H{
			"error": "wrong id",
		})
	} else {
		post.DeleteByID(ID)
		c.JSON(200, gin.H{
			"message": "deleted",
		})
	}
}
