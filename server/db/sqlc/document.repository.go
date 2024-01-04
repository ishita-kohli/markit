package sqlc

import (
	"context"
	"database/sql"
	"server/internal/document"
)

type repository struct {
	q  *Queries
	db *sql.DB
}

func (r *repository) CreateDocument(ctx context.Context, req *document.CreateDocumentReq) (int64, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}
	defer tx.Rollback()

	qtx := r.q.WithTx(tx)

	docId, err := qtx.CreateDocument(ctx, req.Title)
	if err != nil {
		return 0, err
	}

	perm := SetPermissionParams{
		DocumentID: docId,
		UserID:     req.UserID,
		Role:       DocumentAccessRolesOwner,
	}
	if _, err := qtx.SetPermission(ctx, perm); err != nil {
		return 0, err
	}

	if err = tx.Commit(); err != nil {
		return 0, err
	}

	return docId, nil
}

func (r *repository) GetDocumentById(ctx context.Context, id int64) (*document.Document, error) {
	dbDocument, err := r.q.GetDocumentById(ctx, id)
	if err != nil {
		return nil, err
	}

	dbDocumentAccess, err := r.q.GetPermissionsForDocumentId(ctx, id)
	if err != nil {
		return nil, err
	}

	result := &document.Document{
		LeanDocument: document.LeanDocument{
			ID:        dbDocument.ID,
			Title:     dbDocument.Title,
			Body:      dbDocument.Body,
			CreatedAt: dbDocument.CreatedAt,
		},
		Roles: make([]document.DocumentAccess, 0),
	}

	for i := range dbDocumentAccess {
		result.Roles = append(result.Roles, document.DocumentAccess{
			UserID: dbDocumentAccess[i].UserID,
			Role:   string(dbDocumentAccess[i].Role),
		})
	}

	return result, nil
}
func (r *repository) Listdocuments(c context.Context, req *document.DocumentlistReq) ([](*document.LeanDocument), error) {
	dbDocument, err := r.q.GetDocumentListByUser(c, GetDocumentListByUserParams{
		UserID: req.UserID,
		Role:   DocumentAccessRoles(req.Role),
	})
	if err != nil {
		return nil, err
	}

	result := make([](*document.LeanDocument), 0)

	for i := range dbDocument {
		result = append(result, &document.LeanDocument{
			ID:        dbDocument[i].ID,
			Title:     dbDocument[i].Title,
			Body:      dbDocument[i].Body,
			CreatedAt: dbDocument[i].CreatedAt,
		})
	}

	return result, nil
}

func (r *repository) CheckAccess(ctx context.Context, userId int64, documentId int64) (document.PermissionLevel, error) {
	access, err := r.q.GetAccessLevelForDocumentByUser(ctx, GetAccessLevelForDocumentByUserParams{
		DocumentID: documentId,
		UserID:     userId,
	})
	if err != nil {
		return document.NOACCESS, err
	}

	switch access.Role {
	case DocumentAccessRolesOwner:
		return document.OWNER, nil
	case DocumentAccessRolesEditor:
		return document.EDITOR, nil
	case DocumentAccessRolesViewer:
		return document.VIEWER, nil
	default:
		return document.NOACCESS, nil
	}
}
func (r *repository) UpdateDocument(ctx context.Context, documentId int64, body string) error {
	return r.q.UpdateDocumentText(ctx, UpdateDocumentTextParams{
		ID:   documentId,
		Body: body,
	})
}

func NewDocumentRepository(q *Queries, db *sql.DB) document.Repository {
	return &repository{q: q, db: db}
}
