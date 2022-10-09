package postgres

import (
	"fmt"
	pb "template/genproto"

	"github.com/jmoiron/sqlx"
)

type userRepo struct {
	db *sqlx.DB
}

func NewUserRepo(db *sqlx.DB) *userRepo {
	return &userRepo{db: db}
}

func (r *userRepo) Create(user *pb.UserRequest) (*pb.User, error) {
	userResp := pb.User{}
	err := r.db.QueryRow(`insert into 
	users (name, last_name) 
	values($1, $2) 
	returning id, name, last_name`, user.Name, user.LastName).
		Scan(&userResp.Id, &userResp.Name, &userResp.LastName)
	if err != nil {
		return &pb.User{}, err
	}
	return &userResp, nil
}

func (r *userRepo) Delete(id int) (*pb.Empty, error) {
	_, err := r.db.Exec(`DELETE FROM users WHERE id=$1`, id)
	if err != nil {
		return &pb.Empty{}, err
	}

	return &pb.Empty{}, nil
}
func (r *userRepo) Update(req *pb.User) (*pb.User, error) {
	userResp := pb.User{}
	err := r.db.QueryRow(`update users 
	set name=$1, last_name=$2 
	where id = $3
	returning id, name, last_name`, req.Name, req.LastName, req.Id).
		Scan(&userResp.Id, &userResp.Name, &userResp.LastName)
	if err != nil {
		fmt.Println("error while update in postgres/user.go", err)
		return &pb.User{}, err
	}
	return &userResp, nil
}

func (r *userRepo) Get(req *pb.Id) (*pb.User, error) {
	user := &pb.User{}
	err := r.db.QueryRow(`select id,name,last_name
	from users where id=$1`, req.Id).Scan(&user.Id, &user.Name, &user.LastName)
	if err != nil {
		return &pb.User{}, err
	}

	return user, nil
}
