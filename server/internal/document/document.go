package document

import (
	"context"
	"time"
)

type Document struct {
	ID        int64            `json:"id"`
	Title     string           `json:"title"`
	Body      string           `json:"body"`
	Roles     []Documentaccess `json:"roles"`
	CreatedAt time.Time        `json:"created_at"`
}
type Documentaccess struct {
	UserID int64  `json:"user_id"`
	Role   string `json:"role"`
}

type CreateDocumentReq struct {
	Title string `json:"title"`
}

type Repository interface {
	CreateDocument(ctx context.Context, req *CreateDocumentReq, userId int64) (int64, error)
	GetDocumentById(ctx context.Context, id int64) (*Document, error)
}

type Service interface {
	CreateDocument(c context.Context, req *CreateDocumentReq) (*Document, error)
}
