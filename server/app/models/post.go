package models

import (
    "github.com/jinzhu/gorm"
)

// Post struct type.
type Post struct {
    gorm.Model
    Title  string `json:"title"`
    Body   string `json:"body"`
    UserID uint
    User   User `gorm:"foreignKey:UserID; association_foreignKey:UserID"`
}

// Posts are collection of post type.
type Posts []Post

// GetAllPosts is used to fetch all of the posts.
func GetAllPosts() (Posts, error) {
    var posts Posts
    err := db.Find(&posts).Error
    return posts, err
}

// FindByTitle is used to retrieve the post by its id.
func FindByTitle(title string) (*Post, error) {
    var post Post
    err := db.Where("title = ?", title).First(&post).Error
    return &post, err
}
