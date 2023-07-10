package document

import (
	"context"
	"time"
)

const (
	secretKey = "secret"
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

	docId, err := s.Repository.CreateDocument(ctx, req, 0) //TODO
	if err != nil {
		return nil, err
	}

	return s.Repository.GetDocumentById(ctx, docId)
}
