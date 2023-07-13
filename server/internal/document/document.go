package document

import (
	"context"
	"time"
)

type Document struct {
	LeanDocument
	Roles []Documentaccess `json:"roles"`
}
type LeanDocument struct {
	ID        int64     `json:"id"`
	Title     string    `json:"title"`
	Body      string    `json:"body"`
	CreatedAt time.Time `json:"created_at"`
}
type Documentaccess struct {
	UserID int64  `json:"user_id"`
	Role   string `json:"role"`
}
type DocumentlistReq struct {
	UserID int64
	Role   string
}

type CreateDocumentReq struct {
	Title  string `json:"title"`
	UserID int64
}

type Repository interface {
	CreateDocument(ctx context.Context, req *CreateDocumentReq) (int64, error)
	GetDocumentById(ctx context.Context, id int64) (*Document, error)
	Listdocuments(c context.Context, req *DocumentlistReq) ([](*LeanDocument), error)
}

type Service interface {
	CreateDocument(c context.Context, req *CreateDocumentReq) (*Document, error)
	Listdocuments(c context.Context, req *DocumentlistReq) ([](*LeanDocument), error)
}
