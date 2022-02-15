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
	createUserSQL = `INSERT INTO USER(username,age,pwd_hash,additional_information,parent,email) 
	VALUES(?,?,?,?,?,?)`
	getUserSQLById    = `SELECT u.username ,  u.additional_information , u.age , u.parent , u.email FROM USER u  WHERE u.id = ? `
	getUserSQLByEmail = `SELECT u.id FROM USER u  WHERE u.email = ? `
	updateUserSQL     = `UPDATE USER set username = ? , age = ? , pwd_hash = ? , additional_information = ? , parent = ? , email = ? where id = ?  `
	deleteUserSQL     = `DELETE FROM USER WHERE id = ? `
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
	var id string
	stmt, err := r.db.PrepareContext(ctx, getUserSQLByEmail)
	if err != nil {
		return 0, errors.NewInternalError()
	}
	err = stmt.QueryRowContext(ctx, usr.Email).Scan(&id)
	if err == sql.ErrNoRows {
		stmt, err = r.db.PrepareContext(ctx, createUserSQL)
		if err != nil {
			return 0, errors.NewInternalError()
		}
		res, err = stmt.ExecContext(ctx, usr.Name, usr.Age, usr.PwdHash, usr.AdditionalInformation, strings.Join(usr.Parent, ","), usr.Email)
		if err != nil {
			level.Error(r.log).Log(err.Error())
			return 0, errors.NewInternalError()
		}
		id, err := res.LastInsertId()
		if err != nil {
			return 0, errors.NewInternalError()
		}
		return int32(id), nil
	}
	return 0, errors.NewUserExistsError()
}

func (r *userRepository) GetUser(ctx context.Context, id int32) (entities.User, error) {
	var parent *string
	usr := entities.User{}
	stmt, err := r.db.PrepareContext(ctx, getUserSQLById)
	if err != nil {
		return entities.User{}, errors.NewInternalError()
	}
	err = stmt.QueryRow(id).Scan(&usr.Name, &usr.AdditionalInformation, &usr.Age, &parent, &usr.Email)
	if err == sql.ErrNoRows {
		level.Error(r.log).Log(err.Error())
		return entities.User{}, errors.NewUserNotFoundError()
	}
	if err != nil {
		level.Error(r.log).Log(err.Error())
		return entities.User{}, errors.NewInternalError()
	}
	usr.Id = id
	usr.Parent = strings.Split(*parent, ",")
	return usr, nil
}

func (r *userRepository) UpdateUser(ctx context.Context, usr entities.User) error {
	var id int32
	stmt, err := r.db.PrepareContext(ctx, getUserSQLById)
	if err != nil {
		return errors.NewInternalError()
	}
	if err := stmt.QueryRow(usr.Id).Scan(); err == sql.ErrNoRows {
		level.Error(r.log).Log(err.Error())
		return errors.NewUserNotFoundError()
	}
	stmt, err = r.db.PrepareContext(ctx, getUserSQLByEmail)
	if err != nil {
		level.Error(r.log).Log(err.Error())
		return errors.NewInternalError()
	}
	err = stmt.QueryRowContext(ctx, usr.Email).Scan(&id)
	if err != sql.ErrNoRows && id != usr.Id {
		level.Error(r.log).Log("registered email")
		return errors.NewEmailRegisteredError()
	}
	stmt, err = r.db.PrepareContext(ctx, updateUserSQL)
	if err != nil {
		level.Error(r.log).Log(err.Error())
		return errors.NewInternalError()
	}
	_, err = stmt.ExecContext(ctx, usr.Name, usr.Age, usr.PwdHash, usr.AdditionalInformation, strings.Join(usr.Parent, ","), usr.Email, usr.Id)
	if err != nil {
		level.Error(r.log).Log(err.Error())
		return errors.NewInternalError()
	}
	return nil
}

func (r *userRepository) DeleteUser(ctx context.Context, id int32) error {
	stmt, err := r.db.PrepareContext(ctx, getUserSQLById)
	if err != nil {
		return errors.NewInternalError()
	}
	if err := stmt.QueryRow(id).Scan(); err == sql.ErrNoRows {
		level.Error(r.log).Log("Error deleting user ")
		return errors.NewUserNotFoundError()
	}

	stmt, err = r.db.PrepareContext(ctx, deleteUserSQL)
	if err != nil {
		return errors.NewInternalError()
	}
	_, err = stmt.ExecContext(ctx, id)
	if err != nil {
		level.Error(r.log).Log("Error deleting user " + err.Error())
		return errors.NewInternalError()
	}
	return nil
}
