package service

import (
	"context"
	"template/storage"

	pb "template/genproto"
	l "template/pkg/logger"

	"github.com/jmoiron/sqlx"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type UserService struct {
	storage storage.IStorage
	logger  l.Logger
}

func NewUserService(db *sqlx.DB, log l.Logger) *UserService {
	return &UserService{
		storage: storage.NewStorage(db),
		logger:  log,
	}
}

func (s *UserService) Create(ctx context.Context, req *pb.UserRequest) (*pb.User, error) {
	user, err := s.storage.User().Create(req)
	if err != nil {
		s.logger.Error("error insert", l.Any("error insert user", err))
		return &pb.User{}, status.Error(codes.Internal, "something went wrong, please check user create func")
	}

	return user, nil
}

func (s *UserService) Delete(ctx context.Context, req *pb.Id) (*pb.Empty, error) {
	_, err := s.storage.User().Delete(int(req.Id))
	if err != nil {
		s.logger.Error("error insert", l.Any("error insert user", err))
		return &pb.Empty{}, err
	}

	return &pb.Empty{}, nil
}

func (s *UserService) Update(ctx context.Context, req *pb.User) (*pb.User, error) {
	user, err := s.storage.User().Update(req)
	if err != nil {
		s.logger.Error("error while updating user", l.Any("error insert user", err))
		return &pb.User{}, err
	}
	return user, nil
}

func (s *UserService) Get(ctx context.Context, req *pb.Id) (*pb.User, error) {
	user, err := s.storage.User().Get(req)
	if err != nil {
		s.logger.Error("error while updating user", l.Any("error insert user", err))
		return &pb.User{}, err
	}
	return user, nil
}
