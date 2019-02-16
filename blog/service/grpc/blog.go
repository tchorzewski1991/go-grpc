package service

import (
	"github.com/tchorzewski1991/go-grpc/blog/repository"
	blogRepository "github.com/tchorzewski1991/go-grpc/blog/repository/blog"
	"github.com/tchorzewski1991/go-grpc/blog/blogpb"
	"golang.org/x/net/context"
	"google.golang.org/grpc/status"
	"google.golang.org/grpc/codes"
	"github.com/tchorzewski1991/go-grpc/blog/models"
	"log"
)

func NewBlogService(repo repository.BlogRepo) *BlogService {
	return &BlogService{repo: repo}
}

type BlogService struct {
	repo repository.BlogRepo
}

func handleError(err repository.BlogRepoError) error {
	switch err.ErrorCode() {
		case blogRepository.InvalidFormatCode:
			return status.Errorf(codes.NotFound, err.Error())
		case blogRepository.NotFoundCode:
			return status.Errorf(codes.NotFound, err.Error())
		default:
			return status.Errorf(codes.Internal, err.Error())
	}
}

func (s *BlogService) Read(ctx context.Context, req *blogpb.ReadBlogRequest) (
	*blogpb.ReadBlogResponse, error) {

	log.Printf("gRPC | Method: READ | Request: %v", req)

	requestID := req.GetBlogId()

	dbBlog, err := s.repo.Read(ctx, requestID)
	if err != nil {
		log.Printf("gRPC | Method: READ | Error: %v", err)
		return nil, handleError(err)
	}

	return &blogpb.ReadBlogResponse{
		Blog: &blogpb.Blog{
			Id:       dbBlog.ID.Hex(),
			AuthorId: dbBlog.AuthorID,
			Title:    dbBlog.Title,
			Content:  dbBlog.Content,
		},
	}, nil
}

func (s *BlogService) Create(ctx context.Context, req *blogpb.CreateBlogRequest) (
	*blogpb.CreateBlogResponse, error) {

	log.Printf("gRPC | Method: CREATE | Request: %v", req)

	payload := req.GetBlog()

	blogEntry := &models.Blog{
		AuthorID: payload.GetAuthorId(),
		Title:    payload.GetTitle(),
		Content:  payload.GetContent(),
	}

	dbBlog, err := s.repo.Create(context.Background(), blogEntry)
	if err != nil {
		log.Printf("gRPC | Method: CREATE | Error: %v", err)
		return nil, handleError(err)
	}

	return &blogpb.CreateBlogResponse{
		Blog: &blogpb.Blog{
			Id:       dbBlog.ID.Hex(),
			AuthorId: dbBlog.AuthorID,
			Title:    dbBlog.Title,
			Content:  dbBlog.Content,
		},
	}, nil
}

func (s *BlogService) List(req *blogpb.ListBlogRequest, stream blogpb.BlogService_ListServer) error {

	log.Println("gRPC | Method: LIST")

	dbBlogs, err := s.repo.GetAll(context.Background())
	if err != nil {
		log.Printf("gRPC | Method: LIST | Error: %v", err)
		return handleError(err)
	}

	var blogs []*blogpb.Blog

	for _, dbBlog := range dbBlogs {
		blogs = append(blogs, &blogpb.Blog{
			Id:       dbBlog.ID.Hex(),
			Title:    dbBlog.Title,
			Content:  dbBlog.Content,
			AuthorId: dbBlog.AuthorID,
		})
	}

	stream.Send(&blogpb.ListBlogResponse{Blogs: blogs})

	return nil
}
