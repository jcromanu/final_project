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
	createUserSQL = `INSERT INTO USER(username,age,pwd_hash,additional_information,parent) 
	VALUES(?,?,?,?,?)`
	getUserSQL    = `SELECT u.username , u.pwd_hash , u.additional_information , u.age , u.parent FROM USER u  WHERE u.id = ? `
	updateUserSQL = "UPDATE USER set username = ? , age = ? , pwd_hash = ? , additional_information = ? , parent = ? where id = ?  "
	deleteUserSQL = `DELETE FROM USER WHERE id = ? `
)

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
	res, err := r.db.ExecContext(ctx, createUserSQL, usr.Name, usr.Age, usr.PwdHash, usr.AdditionalInformation, strings.Join(usr.Parent, ","))
	if err != nil {
		level.Error(r.log).Log("Error creating user" + err.Error())
		return 0, errors.NewInternalError()
	}
	id, _ := res.LastInsertId()
	return int32(id), nil
}

func (r *userRepository) GetUser(ctx context.Context, id int32) (entities.User, error) {
	var parent *string
	usr := entities.User{}
	err := r.db.QueryRow(getUserSQL, id).Scan(&usr.Name, &usr.PwdHash, &usr.AdditionalInformation, &usr.Age, &parent)
	if err == sql.ErrNoRows {
		level.Error(r.log).Log("Error retrieving user" + err.Error())
		return entities.User{}, errors.NewUserNotFoundError()
	}
	if err != nil {
		level.Error(r.log).Log("Error retrieving user" + err.Error())
		return entities.User{}, errors.NewInternalError()
	}
	usr.Id = id
	usr.Parent = strings.Split(*parent, ",")
	return usr, nil
}

func (r *userRepository) UpdateUser(ctx context.Context, usr entities.User) error {

	if err := r.db.QueryRow(getUserSQL, usr.Id).Scan(); err == sql.ErrNoRows {
		level.Error(r.log).Log("Error retrieving user ")
		return errors.NewUserNotFoundError()
	}

	_, err := r.db.ExecContext(ctx, updateUserSQL, usr.Name, usr.Age, usr.PwdHash, usr.AdditionalInformation, strings.Join(usr.Parent, ","), usr.Id)
	if err != nil {
		level.Error(r.log).Log("Error updating user " + err.Error())
		return errors.NewInternalError()
	}
	return nil
}

func (r *userRepository) DeleteUser(ctx context.Context, id int32) error {
	if err := r.db.QueryRow(getUserSQL, id).Scan(); err == sql.ErrNoRows {
		level.Error(r.log).Log("Error deleting user ")
		return errors.NewUserNotFoundError()
	}
	_, err := r.db.ExecContext(ctx, deleteUserSQL, id)
	if err != nil {
		level.Error(r.log).Log("Error deleting user " + err.Error())
		return errors.NewInternalError()
	}
	return nil
}
