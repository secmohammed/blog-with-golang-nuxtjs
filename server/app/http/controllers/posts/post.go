package posts

import (
    "encoding/json"
    "go-auth-with-crud-api/server/app/http/requests"
    "go-auth-with-crud-api/server/app/models"
    "go-auth-with-crud-api/server/utils"
    "net/http"

    "github.com/gorilla/mux"
)

// Show is used to show the post.
func Show(w http.ResponseWriter, r *http.Request) {
    title := mux.Vars(r)["post"]
    post, err := models.FindByTitle(title)
    if err != nil {
        utils.Respond(w, utils.Message(false, err.Error()))
        return

    }
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(post)
}

//Destroy function is used to destroy the post by its title.
func Destroy(w http.ResponseWriter, r *http.Request) {
    title := mux.Vars(r)["post"]
    post := &models.Post{
        Title: title,
    }
    err := models.Delete(&post, "title", title)
    if err != nil {
        utils.Respond(w, utils.Message(false, err.Error()))

        return
    }
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(post)
}

// Index function is used to index posts.
func Index(w http.ResponseWriter, r *http.Request) {
    var posts models.Posts
    posts, err := models.GetAllPosts()
    if err != nil {
        utils.Respond(w, utils.Message(false, err.Error()))
        return
    }
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(posts)

}

// Store function is used to create a post.
func Store(w http.ResponseWriter, r *http.Request) {
    post := &models.Post{}
    err := json.NewDecoder(r.Body).Decode(&post) //decode the request body into struct and failed if any error occur
    if err != nil {
        utils.Respond(w, utils.Message(false, "Invalid request"))
        return
    }
    post.UserID = r.Context().Value(utils.ContextKeyAuthToken).(uint)

    messages, status := requests.IsSubmittedStorePostFormValid(post)
    if !status {
        w.WriteHeader(http.StatusUnprocessableEntity)
        utils.Respond(w, messages)
        return
    }
    err = models.Create(&post)
    if err != nil {
        utils.Respond(w, utils.Message(false, "Couldn't create post"))
        return
    }
    response := utils.Message(true, "Post has been created")
    response["post"] = post
    utils.Respond(w, response)

}
