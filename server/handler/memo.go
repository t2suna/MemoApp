package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/t2suna/memo_server/model"
)

type MemoHandler struct{}

func NewMemoHandler() *MemoHandler {
	return &MemoHandler{}
}

func (m *MemoHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		var memoRequest model.CreateMemoRequest
		if err := json.NewDecoder(r.Body).Decode(&memoRequest); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
	} else if r.Method == http.MethodGet {
		tmpHR := &model.CreateMemoResponse{Message: "Hello World!"}

		err := json.NewEncoder(w).Encode(tmpHR)
		if err != nil {
			log.Println(err)
		}
	}

}
