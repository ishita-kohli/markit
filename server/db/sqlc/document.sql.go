// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.22.0
// source: document.sql

package sqlc

import (
	"context"
	"time"
)

const createDocument = `-- name: CreateDocument :one
INSERT INTO documents(title)
VALUES($1)
RETURNING id
`

func (q *Queries) CreateDocument(ctx context.Context, title string) (int64, error) {
	row := q.db.QueryRowContext(ctx, createDocument, title)
	var id int64
	err := row.Scan(&id)
	return id, err
}

const getAccessLevelForDocumentByUser = `-- name: GetAccessLevelForDocumentByUser :one
SELECT document_id, user_id, role 
FROM document_access
WHERE document_id = $1
    AND user_id = $2
`

type GetAccessLevelForDocumentByUserParams struct {
	DocumentID int64
	UserID     int64
}

func (q *Queries) GetAccessLevelForDocumentByUser(ctx context.Context, arg GetAccessLevelForDocumentByUserParams) (DocumentAccess, error) {
	row := q.db.QueryRowContext(ctx, getAccessLevelForDocumentByUser, arg.DocumentID, arg.UserID)
	var i DocumentAccess
	err := row.Scan(&i.DocumentID, &i.UserID, &i.Role)
	return i, err
}

const getDocumentById = `-- name: GetDocumentById :one
SELECT id, title, body, created_at, updated_at
FROM documents
WHERE id = $1
`

func (q *Queries) GetDocumentById(ctx context.Context, id int64) (Document, error) {
	row := q.db.QueryRowContext(ctx, getDocumentById, id)
	var i Document
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.Body,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getDocumentListByUser = `-- name: GetDocumentListByUser :many
SELECT d.id,
    d.title,
    d.body,
    d.created_at
FROM document_access AS a
    INNER JOIN documents AS d ON a.document_id = d.id
WHERE a.user_id = $1
    AND a.role = $2
`

type GetDocumentListByUserParams struct {
	UserID int64
	Role   DocumentAccessRoles
}

type GetDocumentListByUserRow struct {
	ID        int64
	Title     string
	Body      string
	CreatedAt time.Time
}

func (q *Queries) GetDocumentListByUser(ctx context.Context, arg GetDocumentListByUserParams) ([]GetDocumentListByUserRow, error) {
	rows, err := q.db.QueryContext(ctx, getDocumentListByUser, arg.UserID, arg.Role)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetDocumentListByUserRow
	for rows.Next() {
		var i GetDocumentListByUserRow
		if err := rows.Scan(
			&i.ID,
			&i.Title,
			&i.Body,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getPermissionsForDocumentId = `-- name: GetPermissionsForDocumentId :many
SELECT document_id, user_id, role
FROM document_access
WHERE document_id = $1
`

func (q *Queries) GetPermissionsForDocumentId(ctx context.Context, documentID int64) ([]DocumentAccess, error) {
	rows, err := q.db.QueryContext(ctx, getPermissionsForDocumentId, documentID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []DocumentAccess
	for rows.Next() {
		var i DocumentAccess
		if err := rows.Scan(&i.DocumentID, &i.UserID, &i.Role); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const setPermission = `-- name: SetPermission :one
INSERT INTO document_access(document_id, user_id, role)
VALUES($1, $2, $3)
ON CONFLICT (document_id,user_id)
DO UPDATE SET role = EXCLUDED.role
RETURNING document_id
`

type SetPermissionParams struct {
	DocumentID int64
	UserID     int64
	Role       DocumentAccessRoles
}

func (q *Queries) SetPermission(ctx context.Context, arg SetPermissionParams) (int64, error) {
	row := q.db.QueryRowContext(ctx, setPermission, arg.DocumentID, arg.UserID, arg.Role)
	var document_id int64
	err := row.Scan(&document_id)
	return document_id, err
}

const updateDocumentText = `-- name: UpdateDocumentText :exec
UPDATE documents
SET body = $2
WHERE id = $1
`

type UpdateDocumentTextParams struct {
	ID   int64
	Body string
}

func (q *Queries) UpdateDocumentText(ctx context.Context, arg UpdateDocumentTextParams) error {
	_, err := q.db.ExecContext(ctx, updateDocumentText, arg.ID, arg.Body)
	return err
}
