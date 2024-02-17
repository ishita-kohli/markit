package document

import (
	"context"
	"fmt"
	"time"
)

type service struct {
	Repository
	timeout time.Duration
}

func NewService(repository Repository) Service {
	return &service{
		repository,
		time.Duration(2) * time.Second,
	}
}

func (s *service) CreateDocument(c context.Context, req *CreateDocumentReq) (*Document, error) {
	ctx, cancel := context.WithTimeout(c, s.timeout)
	defer cancel()

	docId, err := s.Repository.CreateDocument(ctx, req)
	if err != nil {
		return nil, err
	}

	return s.Repository.GetDocumentById(ctx, docId)
}
func (s *service) Listdocuments(c context.Context, req *DocumentlistReq) ([](*LeanDocument), error) {
	ctx, cancel := context.WithTimeout(c, s.timeout)
	defer cancel()
	return s.Repository.Listdocuments(ctx, req)

}
func (s *service) Getdocument(c context.Context, req *GetDocumentByIDReq) (*Document, error) {
	ctx, cancel := context.WithTimeout(c, s.timeout)
	defer cancel()

	accessLevel, err := s.Repository.CheckAccess(ctx, req.UserID, req.DocumentID)
	if err != nil {
		return nil, err
	}

	if accessLevel == NOACCESS {
		return nil, fmt.Errorf("You don't have access\n")
	}

	return s.Repository.GetDocumentById(ctx, req.DocumentID)
}
func (s *service) UpdateDocument(c context.Context, req *UpdateDocumentReq) error {
	ctx, cancel := context.WithTimeout(c, s.timeout)
	defer cancel()

	accessLevel, err := s.Repository.CheckAccess(ctx, req.UserID, req.DocumentID)
	if err != nil {
		return err
	}

	if accessLevel == NOACCESS || accessLevel == VIEWER {
		return fmt.Errorf("You don't have access\n")
	}

	return s.Repository.UpdateDocument(ctx, req.DocumentID, req.Body)
}
func (s *service) ShareDocument(c context.Context, req *ShareDocumentReq) error {
	ctx, cancel := context.WithTimeout(c, s.timeout)
	defer cancel()

	role := ReadPermissionLevelFromString(req.Role)

	accessLevel, err := s.Repository.CheckAccess(ctx, req.CurrentUserID, req.DocumentID)
	if err != nil {
		return err
	}

	if accessLevel == NOACCESS {
		return fmt.Errorf("You don't have access\n")
	}
	if accessLevel == VIEWER && role != VIEWER {
		return fmt.Errorf("You can only give viewer access")
	}

	return s.Repository.AddAccess(ctx, req.DocumentID, req.ShareUserID, role)
}
