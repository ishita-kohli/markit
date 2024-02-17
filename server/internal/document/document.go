package document

import (
	"context"
	"time"
)

type Document struct {
	LeanDocument
	Roles []DocumentAccess `json:"roles"`
}

type PermissionLevel int

const (
	OWNER PermissionLevel = iota
	EDITOR
	VIEWER
	NOACCESS
)

func ReadPermissionLevelFromString(payload string) PermissionLevel {
	if payload == "owner" {
		return OWNER
	}
	if payload == "editor" {
		return EDITOR
	}
	if payload == "viewer" {
		return VIEWER
	}
	return NOACCESS
}

type LeanDocument struct {
	ID        int64     `json:"id"`
	Title     string    `json:"title"`
	Body      string    `json:"body"`
	CreatedAt time.Time `json:"created_at"`
}
type DocumentAccess struct {
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

type UpdateDocumentReq struct {
	UserID     int64
	DocumentID int64
	Body       string `json:"text"`
}
type ShareDocumentReq struct {
	CurrentUserID int64
	ShareUserID   int64 `json:"share_with"`
	DocumentID    int64
	Role          string `json:"role"`
}

type Repository interface {
	CreateDocument(ctx context.Context, req *CreateDocumentReq) (int64, error)
	GetDocumentById(ctx context.Context, id int64) (*Document, error)
	Listdocuments(c context.Context, req *DocumentlistReq) ([](*LeanDocument), error)
	CheckAccess(ctx context.Context, userId int64, documentId int64) (PermissionLevel, error)
	UpdateDocument(ctx context.Context, documentId int64, body string) error
	AddAccess(ctx context.Context, documentId int64, userId int64, role PermissionLevel) error
}

type Service interface {
	CreateDocument(c context.Context, req *CreateDocumentReq) (*Document, error)
	Listdocuments(c context.Context, req *DocumentlistReq) ([](*LeanDocument), error)
	Getdocument(c context.Context, req *GetDocumentByIDReq) (*Document, error)
	UpdateDocument(c context.Context, req *UpdateDocumentReq) error
	ShareDocument(c context.Context, req *ShareDocumentReq) error
}
