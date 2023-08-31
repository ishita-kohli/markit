package document

import (
	"context"
	"time"
)

type Document struct {
	LeanDocument
	Roles []Documentaccess `json:"roles"`
}

type DocumentAccess int

const (
	OWNER DocumentAccess = iota
	EDITOR
	VIEWER
	NOACCESS
)

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
type GetDocumentByIDReq struct {
	UserID     int64
	DocumentID int64
}

type Repository interface {
	CreateDocument(ctx context.Context, req *CreateDocumentReq) (int64, error)
	GetDocumentById(ctx context.Context, id int64) (*Document, error)
	Listdocuments(c context.Context, req *DocumentlistReq) ([](*LeanDocument), error)
	CheckAccess(ctx context.Context, userId int64, documentId int64) (DocumentAccess, error)
}

type Service interface {
	CreateDocument(c context.Context, req *CreateDocumentReq) (*Document, error)
	Listdocuments(c context.Context, req *DocumentlistReq) ([](*LeanDocument), error)
	Getdocument(c context.Context, req *GetDocumentByIDReq) (*Document, error)
}
