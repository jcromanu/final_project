package userservice

import (
	"context"
	"database/sql"
	"strings"

	"github.com/go-kit/log"
	"github.com/go-kit/log/level"
	"github.com/jcromanu/final_project/user_service/errors"
	"github.com/jcromanu/final_project/user_service/pkg/entities"
)

const (
	create_user_sql = `INSERT INTO USER(username,age,pwd_hash,additional_information,parent) 
	VALUES(?,?,?,?,?)`
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
	res, err := r.db.ExecContext(ctx, create_user_sql, usr.Name, usr.Age, usr.Pwd_hash, usr.Additional_information, strings.Join(usr.Parent, ","))
	if err != nil {
		level.Error(r.log).Log("Error creating user" + err.Error())
		return 0, errors.NewDatabaseError()
	}
	id, _ := res.LastInsertId()
	return int32(id), nil
}
