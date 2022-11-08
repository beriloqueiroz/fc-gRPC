package service

import (
	"github.com/beriloqueiroz/fc-gRPC/internal/database"
	"github.com/beriloqueiroz/fc-gRPC/internal/pb"
)

type CategoryService struct {
	pb.UnimplementedCategoryServiceServer
	CategoryDB database.Category
}
