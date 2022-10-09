package repo

import (
	pb "template/genproto"
)

type UserStorageI interface {
	Create(*pb.UserRequest) (*pb.User, error)
	Delete(int) (*pb.Empty, error)
	Update(*pb.User) (*pb.User, error)
	Get(*pb.Id) (*pb.User, error)
}
