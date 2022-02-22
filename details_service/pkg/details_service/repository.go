package userservice

import (
	"context"

	"github.com/go-kit/log"
	"github.com/go-kit/log/level"
	"github.com/jcromanu/final_project/user_service/errors"
	"github.com/jcromanu/final_project/user_service/pkg/entities"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type userRepository struct {
	db  *mongo.Database
	log log.Logger
}

func NewUserRepository(db *mongo.Database, log log.Logger) *userRepository {
	return &userRepository{
		db:  db,
		log: log,
	}
}

func (r *userRepository) CreateUser(ctx context.Context, usr entities.User) (int32, error) {
	type seqx struct {
		Id int32 `bson:"sequence_value"`
	}

	seq := &seqx{}
	collection := r.db.Collection("user_collection")
	if cursor := collection.FindOne(ctx, bson.M{"email": usr.Email}); cursor.Err() != mongo.ErrNoDocuments {
		level.Error(r.log).Log(cursor.Err())
		return 0, errors.NewUserExistsError()
	}
	sequence := r.db.Collection("sequence")
	_, err := sequence.UpdateByID(ctx, "userid", bson.M{"$inc": bson.D{{"sequence_value", 1}}})
	if err != nil {
		level.Error(r.log).Log(err.Error())
		return 0, errors.NewInternalError()
	}
	err = sequence.FindOne(ctx, bson.M{"_id": "userid"}).Decode(&seq)
	if err == mongo.ErrNoDocuments {
		level.Error(r.log).Log(err.Error())
		return 0, errors.NewInternalError()
	}

	level.Info(r.log).Log(seq)
	level.Info(r.log).Log(seq.Id)
	usr.Id = seq.Id

	_, err = collection.InsertOne(ctx, usr)
	if err != nil {
		level.Error(r.log).Log(err.Error())
		return 0, errors.NewInternalError()
	}

	return usr.Id, nil
}

func (r *userRepository) GetUser(ctx context.Context, id int32) (entities.User, error) {
	var usr entities.User
	collection := r.db.Collection("user_collection")
	if err := collection.FindOne(ctx, bson.M{"id": id}).Decode(&usr); err == mongo.ErrNoDocuments {
		level.Error(r.log).Log(err)
		return entities.User{}, errors.NewUserNotFoundError()
	}
	usr.PwdHash = ""
	return usr, nil
}

func (r *userRepository) UpdateUser(ctx context.Context, usr entities.User) error {
	decusr := entities.User{}
	collection := r.db.Collection("user_collection")
	if err := collection.FindOne(ctx, bson.M{"id": usr.Id}).Decode(&decusr); err == mongo.ErrNoDocuments {
		level.Error(r.log).Log(err)
		return errors.NewUserNotFoundError()
	}
	_, err := collection.ReplaceOne(ctx, bson.M{"id": usr.Id}, usr)
	if err != nil {
		level.Error(r.log).Log(err)
		return errors.NewInternalError()
	}
	return nil
}

func (r *userRepository) DeleteUser(ctx context.Context, id int32) error {
	collection := r.db.Collection("user_collection")
	if err := collection.FindOne(ctx, bson.M{"id": id}); err.Err() == mongo.ErrNoDocuments {
		level.Error(r.log).Log(err)
		return errors.NewUserNotFoundError()
	}
	_, err := collection.DeleteOne(ctx, bson.M{"id": id})
	if err != nil {
		level.Error(r.log).Log(err)
		return errors.NewInternalError()
	}
	return nil
}

/*
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
		level.Error(r.log).Log(errors.NewEmailRegisteredError().Error())
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
}*/
