package config

import "github.com/soundsnick/arzamas/models"

// SeedTestPosts creates test posts
func SeedTestPosts() {
	db := models.GetDB()
	user := db.Find(&models.User{Email: "soundsnick@gmail.com"}).Value.(*models.User)
	db.Create(&models.Post{Title: "Test post", Content: "This is first test post, Lorem ipsum dolor sit amet", UserID: user.ID, User: *user})
}

// SeedTestUser creates test user
func SeedTestUser() {
	db := models.GetDB()
	db.Delete(&models.User{Email: "test.1313213231323.user@gmail.com"})
	db.Create(&models.User{Name: "Test", LastName: "User", Email: "test.1313213231323.user@gmail.com", Password: "123123123qwerty"})
}
