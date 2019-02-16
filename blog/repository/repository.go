package repository

// Repository represents all contracts for any database related operations.
// High level 'contract' definition refers to the formalized form of
// interaction with the system's component. Business logic shouldn't
// be placed here.

import (
	"context"
	"github.com/tchorzewski1991/go-grpc/blog/models"
)

type BlogRepo interface {
	Read(ctx context.Context, ID string) (*models.Blog, BlogRepoError)
	Create(ctx context.Context, blog *models.Blog) (*models.Blog, BlogRepoError)
	GetAll(ctx context.Context) ([]*models.Blog, BlogRepoError)
}

type BlogRepoError interface {
	ErrorCode() int
	error
}
