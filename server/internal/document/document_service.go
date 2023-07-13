package document

import (
	"context"
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
