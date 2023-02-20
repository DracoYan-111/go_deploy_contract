package handler

import (
	"GoContractDeployment/models"
	"GoContractDeployment/navigation"
	repository "GoContractDeployment/repository"
	post "GoContractDeployment/repository/post"
	"encoding/json"
	"fmt"
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

type Receive struct {
	AataList string `json:"sign"`
}

// 添加任务
func (p *Post) Create(w http.ResponseWriter, r *http.Request) {
	//num:=chi.URLParam(r, "sign")
	////num, _ := strconv.Atoi(chi.URLParam(r, "sign"))
	//
	//log.Println(num)

	var requestBody Receive

	err := json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	fmt.Printf("Receiveds:", requestBody.AataList)
	//requestBody.AataList = strings.ReplaceAll(requestBody.AataList, "'", "\"")
	//fmt.Printf("Receiveds:", requestBody.AataList)

	var data []models.ReceivePost

	err = json.Unmarshal([]byte(requestBody.AataList), &data)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println()
	fmt.Printf("Receiveds:", data)

	respondwithJSON(w, http.StatusOK, data)

}

// respondwithJSON write json response format
func respondwithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

// respondwithError return error message
func respondWithError(w http.ResponseWriter, code int, msg string) {
	respondwithJSON(w, code, map[string]string{"message": msg})
}
