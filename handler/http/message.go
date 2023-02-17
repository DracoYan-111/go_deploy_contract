package handler

import (
	"GoContractDeployment/navigation"
	repository "GoContractDeployment/repository"
	post "GoContractDeployment/repository/post"
	"encoding/json"
	"github.com/go-chi/chi"
	"net/http"
	"strconv"
)

// NewPostHandler ...
func NewPostHandler(db *navigation.DB) *Post {
	return &Post{
		// 加载到接口的实例
		repo: post.NewSQLPostRepo(db.SQL),
	}
}

// Post 返回所有的接口
type Post struct {
	repo repository.PostRepo
}

// Fetch all post data
func (p *Post) Fetch(w http.ResponseWriter, r *http.Request) {
	num, _ := strconv.Atoi(chi.URLParam(r, "id"))

	payload := p.repo.Fetch(r.Context(), int64(num))

	respondwithJSON(w, http.StatusOK, payload)
}

// 添加任务
func (p *Post) AddJob(w http.ResponseWriter, r *http.Request) {
	//p.repo.AddJob()
}

// respondwithJSON write json response format
func respondwithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
