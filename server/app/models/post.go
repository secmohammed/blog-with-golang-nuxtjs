package models

// Post struct type.
type Post struct {
    BaseGorm
    Title  string `json:"title"`
    Body   string `json:"body"`
    UserID uint   `json:"user_id"`
    User   User   `json:"user"`
}

// Posts are collection of post type.
type Posts []Post

// GetAllPosts is used to fetch all of the posts.
func GetAllPosts() (Posts, error) {
    var posts Posts
    err := db.Preloads("User").Find(&posts).Error

    return posts, err
}

// FindByTitle is used to retrieve the post by its id.
func FindByTitle(title string) (*Post, error) {
    var post Post

    err := db.Where("title = ?", title).First(&post).Error
    return &post, err
}
