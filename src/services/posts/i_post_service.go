package service

import (
	repo "github.com/red-gold/telar-core/data"
	dto "github.com/red-gold/telar-web/src/domain/posts"
	uuid "github.com/satori/go.uuid"
)

type PostService interface {
	SavePost(post *dto.Post) error
	FindOnePost(filter interface{}) (*dto.Post, error)
	FindPostList(filter interface{}, limit int64, skip int64, sort map[string]int) ([]dto.Post, error)
	QueryPost(search string, ownerUserId *uuid.UUID, postTypeId *int, sortBy string, page int64) ([]dto.Post, error)
	FindById(objectId uuid.UUID) (*dto.Post, error)
	FindByOwnerUserId(ownerUserId uuid.UUID) ([]dto.Post, error)
	UpdatePost(filter interface{}, data interface{}, opts ...*repo.UpdateOptions) error
	UpdatePostById(data *dto.Post) error
	DeletePost(filter interface{}) error
	DeletePostByOwner(ownerUserId uuid.UUID, postId uuid.UUID) error
	DeleteManyPost(filter interface{}) error
	CreatePostIndex(indexes map[string]interface{}) error
	IncrementScoreCount(objectId uuid.UUID, ownerUserId uuid.UUID) error
	DecrementScoreCount(objectId uuid.UUID, ownerUserId uuid.UUID) error
	Increment(objectId uuid.UUID, field string, value int) error
	IncrementCommentCount(objectId uuid.UUID) error
	DecerementCommentCount(objectId uuid.UUID) error
}
