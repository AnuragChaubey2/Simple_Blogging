// handler/posts/posts.go

package posts

import (
	"database/sql"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/AnuragChaubey2/Simple_Blogging.git/models"
	"github.com/AnuragChaubey2/Simple_Blogging.git/services/posts"
	"github.com/gorilla/mux"
)

type PostsHandler struct {
	postsService posts.PostsService
}

func NewPostsHandler(postsService posts.PostsService) *PostsHandler {
	return &PostsHandler{postsService: postsService}
}

func (h *PostsHandler) GetAllPosts(w http.ResponseWriter, r *http.Request) {
	posts, err := h.postsService.GetAllPosts(r.Context())
	if err != nil {
		http.Error(w, "Failed to retrieve blog posts", http.StatusInternalServerError)
		return
	}

	jsonResponse(w, posts, http.StatusOK)
}

func (h *PostsHandler) GetPostByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	postID, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid post ID", http.StatusBadRequest)
		return
	}

	post, err := h.postsService.GetPostByID(r.Context(), postID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			http.Error(w, "Post not found", http.StatusNotFound)
		} else {
			http.Error(w, "Failed to retrieve post", http.StatusInternalServerError)
		}
		return
	}

	jsonResponse(w, post, http.StatusOK)
}

func (h *PostsHandler) CreatePost(w http.ResponseWriter, r *http.Request) {
	var post models.Post
	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&post); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	createdID, err := h.postsService.CreatePost(r.Context(), post)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := map[string]int{"id": createdID}

	jsonResponse(w, response, http.StatusCreated)
}

func jsonResponse(w http.ResponseWriter, data interface{}, status int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		http.Error(w, "Failed to encode JSON response", http.StatusInternalServerError)
	}
}

func (h *PostsHandler) UpdatePost(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    postID, err := strconv.Atoi(vars["id"])
    if err != nil {
        http.Error(w, "Invalid post ID", http.StatusBadRequest)
        return
    }

    var updatedPost models.Post
    decoder := json.NewDecoder(r.Body)
    if err := decoder.Decode(&updatedPost); err != nil {
        http.Error(w, "Invalid request body", http.StatusBadRequest)
        return
    }

    updatedPost, err = h.postsService.UpdatePost(r.Context(), postID, updatedPost)
    if err != nil {
        if errors.Is(err, sql.ErrNoRows) {
            http.Error(w, "Post not found", http.StatusNotFound)
        } else {
            http.Error(w, "Failed to update post", http.StatusInternalServerError)
        }
        return
    }

    jsonResponse(w, updatedPost, http.StatusOK)
}

func (h *PostsHandler) DeletePost(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    postID, err := strconv.Atoi(vars["id"])
    if err != nil {
        http.Error(w, "Invalid post ID", http.StatusBadRequest)
        return
    }

    if err := h.postsService.DeletePost(r.Context(), postID); err != nil {
        if errors.Is(err, sql.ErrNoRows) {
            http.Error(w, "Post not found", http.StatusNotFound)
        } else {
            http.Error(w, "Failed to delete post", http.StatusInternalServerError)
        }
        return
    }

    successMessage := "Record with ID " + strconv.Itoa(postID) + " is successfully deleted"
    w.Write([]byte(successMessage))
}