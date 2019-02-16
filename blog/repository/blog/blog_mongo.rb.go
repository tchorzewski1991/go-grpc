package blog

// This is MongoDB implementation of BlogRepo interface. It contains logic
// and error handling closely tied to that particular database system.
// File name contains suffix to make clear distinction what implementation
// it contains. There might be other BlogRepo implementations in the future
// like BoltDB, or even more common Postgres.

import (
	"github.com/mongodb/mongo-go-driver/mongo"
	"github.com/tchorzewski1991/go-grpc/blog/models"
	"context"
	"github.com/mongodb/mongo-go-driver/bson/primitive"
	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/tchorzewski1991/go-grpc/blog/repository"
)

const (
	InvalidFormatCode = iota
	NotFoundCode
	InternalCode
)

type InvalidFormatError struct { errorMessage string }
type NotFoundError      struct { errorMessage string }
type InternalError      struct { errorMessage string }

func(e *InvalidFormatError) Error() string { return e.errorMessage }
func(e *NotFoundError)      Error() string { return e.errorMessage }
func(e *InternalError)      Error() string { return e.errorMessage }

func(e *InvalidFormatError) ErrorCode() int { return InvalidFormatCode }
func(e *NotFoundError)      ErrorCode() int { return NotFoundCode }
func(e *InternalError)      ErrorCode() int { return InternalCode }

func NewMongoDbBlogRepo(db *mongo.Database) *mongoDbBlogRepo {
	return &mongoDbBlogRepo{Coll: db.Collection("blogs")}
}

type mongoDbBlogRepo struct {
	Coll *mongo.Collection
}

func (r *mongoDbBlogRepo) Read(ctx context.Context, ID string) (
	*models.Blog, repository.BlogRepoError) {

	oid, err := primitive.ObjectIDFromHex(ID)
	if err != nil { return nil, &InvalidFormatError{"ID format is invalid"} }

	blog   := &models.Blog{}
	filter := bson.D{{"_id", oid}}

	err = r.Coll.FindOne(context.Background(), filter).Decode(blog)
	if err != nil { return nil, &NotFoundError{"Blog is not found"} }

	return blog, nil
}

func (r *mongoDbBlogRepo) Create(ctx context.Context, blogEntry *models.Blog) (
	*models.Blog, repository.BlogRepoError) {

	insertion, err := r.Coll.InsertOne(context.Background(), blogEntry)
	if err != nil { return nil, &InternalError{"Internal server error"} }

	oid, ok := insertion.InsertedID.(primitive.ObjectID)
	if !ok { return nil, &InternalError{"Cannot convert to OID"} }

	blogEntry.ID = oid

	return blogEntry, nil
}

func (r *mongoDbBlogRepo) GetAll(ctx context.Context) (
	[]*models.Blog, repository.BlogRepoError) {

	cursor, err := r.Coll.Find(context.Background(), bson.D{})
	defer cursor.Close(context.Background())
	if err != nil { return nil, &InternalError{"Internal server error"} }

	var blogs []*models.Blog

	for cursor.Next(context.TODO()) {
		item := &models.Blog{}

		err = cursor.Decode(item)
		if err != nil { return nil, &InternalError{"Internal server error"} }

		blogs = append(blogs, item)
	}

	if c := cursor.Err(); c != nil {
		return nil, &InternalError{"Internal server error"}
	}

	return blogs, nil
}