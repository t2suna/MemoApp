package model

import "time"

type (
	// A TODO expresses ...
	Memo struct {
		ID          int64     `json:"id"`
		Subject     string    `json:"subject"`
		Description string    `json:"description"`
		CreatedAt   time.Time `json:"created_at"`
		UpdatedAt   time.Time `json:"updated_at"`
	}
	// A CreateTODORequest expresses ...
	CreateMemoRequest struct {
		Subject     string `json:"subject"`
		Description string `json:"description"`
	}
	// A CreateTODOResponse expresses ...
	CreateMemoResponse struct {
		Message string `json:"message"`
	}

	// A ReadTODORequest expresses ...
	ReadMemoRequest struct {
		PrevID int64 `json:"prev_id"`
		//Size   int64 `json:"size"`
	}
	// A ReadTODOResponse expresses ...
	ReadMemoResponse struct {
		Memos []*Memo `json:"memos"`
	}

	// A UpdateTODORequest expresses ...
	UpdateMemoRequest struct {
		ID          int64  `json:"id"`
		Subject     string `json:"subject"`
		Description string `json:"description"`
	}
	// A UpdateTODOResponse expresses ...
	UpdateMemoResponse struct {
		TODO Memo `json:"memo"`
	}

	// A DeleteTODORequest expresses ...
	DeleteMemoRequest struct {
		IDs []int64 `json:"ids"`
	}
	// A DeleteTODOResponse expresses ...
	DeleteMemoResponse struct{}
)
