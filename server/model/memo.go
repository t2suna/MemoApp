package model

import "time"

type (
	// A TODO expresses ...
	TODO struct {
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
	ReadTODORequest struct {
		PrevID int64 `json:"prev_id"`
		//Size   int64 `json:"size"`
	}
	// A ReadTODOResponse expresses ...
	ReadTODOResponse struct {
		TODOs []*TODO `json:"todos"`
	}

	// A UpdateTODORequest expresses ...
	UpdateTODORequest struct {
		ID          int64  `json:"id"`
		Subject     string `json:"subject"`
		Description string `json:"description"`
	}
	// A UpdateTODOResponse expresses ...
	UpdateTODOResponse struct {
		TODO TODO `json:"todo"`
	}

	// A DeleteTODORequest expresses ...
	DeleteTODORequest struct {
		IDs []int64 `json:"ids"`
	}
	// A DeleteTODOResponse expresses ...
	DeleteTODOResponse struct{}
)