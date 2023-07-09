package user

import (
	"context"
	"database/sql"
)

type DBTX interface {
	ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error)
	PrepareContext(context.Context, string) (*sql.Stmt, error)
	QueryContext(context.Context, string, ...interface{}) (*sql.Rows, error)
	QueryRowContext(context.Context, string, ...interface{}) *sql.Row
}

type repository struct {
	db DBTX
}

func NewRepository(db DBTX) Repository {
	return &repository{db: db}
}

func (r *repository) CreateDocument(ctx context.Context, document *Document) (int64, error) {
	var lastInsertId int64
	query := `WITH (INSERT INTO documents(title) VALUES($1) RETURNING id) AS ins INSERT INTO document_access(document_id, user_id, role) VALUES(
		SELECT id, $2, 'owner' FROM ins 
	)RETURNING document_id`
	err := r.db.QueryRowContext(ctx, query, document.Title).Scan(&lastInsertId)
	if err != nil {
		return 0, err
	}
	return lastInsertId, nil
}

func (r *repository) GetDocumentbydocumentid(ctx context.Context, document_id string) (*Document, error) {
	d := Document{}
	query := "SELECT id,title,body,created_at FROM documents WHERE id = $1"
	err := r.db.QueryRowContext(ctx, query, document_id).Scan(&d, &d.ID, &d.Title, &d.CreatedAt)
	if err != nil {
		return &Document{}, nil
	}

	return &d, nil
}
