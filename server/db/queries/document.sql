-- name: CreateDocument :one
INSERT INTO documents(title)
VALUES($1)
RETURNING id;

-- name: UpdateDocumentText :exec
UPDATE documents
SET body = $2
WHERE id = $1;

-- name: SetPermission :one
INSERT INTO document_access(document_id, user_id, role)
VALUES($1, $2, $3)
ON CONFLICT (document_id,user_id)
DO UPDATE SET role = EXCLUDED.role
RETURNING document_id;

-- name: GetDocumentById :one
SELECT *
FROM documents
WHERE id = $1;

-- name: GetPermissionsForDocumentId :many
SELECT *
FROM document_access
WHERE document_id = $1;

-- name: GetAccessLevelForDocumentByUser :one
SELECT * 
FROM document_access
WHERE document_id = $1
    AND user_id = $2;

-- name: GetDocumentListByUser :many
SELECT d.id,
    d.title,
    d.body,
    d.created_at
FROM document_access AS a
    INNER JOIN documents AS d ON a.document_id = d.id
WHERE a.user_id = $1
    AND a.role = $2;

