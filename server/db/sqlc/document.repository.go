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
		Leandocument: document.Leandocument{
			ID:        dbDocument.ID,
			Title:     dbDocument.Title,
			Body:      dbDocument.Body,
			CreatedAt: dbDocument.CreatedAt,
		},
		Roles: make([]document.Documentaccess, 0),
	}

	for i := range dbDocumentAccess {
		result.Roles = append(result.Roles, document.Documentaccess{
			UserID: dbDocumentAccess[i].UserID,
			Role:   string(dbDocumentAccess[i].Role),
		})
	}

	return result, nil
}

func NewDocumentRepository(q *Queries, db *sql.DB) document.Repository {
	return &repository{q: q, db: db}
}
