-- name: CreateDocument :one
INSERT INTO documents(title)
VALUES($1)
RETURNING id;

-- name: SetPermission :one
INSERT INTO document_access(document_id, user_id, role)
VALUES($1, $2, $3)
RETURNING document_id;

-- name: GetDocumentById :one
SELECT *
FROM documents
WHERE id = $1;

-- name: GetPermissionsForDocumentId :many
SELECT *
FROM document_access
WHERE document_id = $1;