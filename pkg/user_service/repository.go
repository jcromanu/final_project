package userservice

import (
	"context"
	"database/sql"

	"github.com/go-kit/kit/log"

	"github.com/jcromanu/final_project/errors"
	"github.com/jcromanu/final_project/pkg/entities"
)

const (
	create_user_sql = `INSERT INTO USERS(name,age,pwd_hash,additional_information,parent) 
	VALUES(?,?,?,?,')`
)

type UserRepository interface {
	CreateUser(context.Context, entities.User) (int32, error)
	/*GetUser(context.Context, int) (entities.User, error)
	UpdateUser(context.Context, entities.User) (entities.User, error)
	DeleteUser(context.Context, int) (bool, error)
	Authenticate(context.Context, string, string) (bool, error)*/
}

type userRepository struct {
	db  *sql.DB
	log log.Logger
}

func NewUserRepository(db *sql.DB, log log.Logger) *userRepository {
	return &userRepository{
		db:  db,
		log: log,
	}
}

func (r *userRepository) CreateUser(ctx context.Context, usr entities.User) (int32, error) {
	var res sql.Result
	res, err := r.db.ExecContext(ctx, create_user_sql, usr.Name, usr.Pwd_hash, usr.Additional_information, usr.Parent)
	if err != nil {
		return 0, errors.NewDatabaseError()
	}
	id, _ := res.LastInsertId()
	return int32(id), nil
}
