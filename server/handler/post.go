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
	tmpHR := &model.CreateMemoResponse{Message: "OK"}

	err := json.NewEncoder(w).Encode(tmpHR)
	if err != nil {
		log.Println(err)
	}
}
