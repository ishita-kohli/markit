package user

import "context"

type Document struct {
	ID        int64            `json:"id"`
	Title     string           `json:"title"`
	Body      string           `json:"body"`
	Roles     []Documentaccess `json:"roles"`
	CreatedAt string           `json:"created_at"`
}
type Documentaccess struct {
	UserID string `json:"user_id"`
	Role   string `json:"role"`
}

type CreateDocumentReq struct {
	Title string `json:"title"`
}

type CreateDocumentRes struct {
	ID        int64  `json:"id"`
	Title     string `json:"title"`
	Body      string `json:"body"`
	CreatedAt string `json:"created_at"`
}

type Repository interface {
	CreateDocument(ctx context.Context, user *Document) (int64, error)
}

type Service interface {
	CreateDocument(c context.Context, req *CreateDocumentReq) (*CreateDocumentRes, error)
}
