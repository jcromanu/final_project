package entities

type User struct {
	Id                    int32    `bson:"id"`
	PwdHash               string   `validate:"required" bson:"pwd_hash"`
	Name                  string   `validate:"required" bson:"name"`
	Age                   int32    `validate:"required" bson:"age"`
	AdditionalInformation string   `validate:"required" bson:"additional_information"`
	Parent                []string `bson:"parent"`
	Email                 string   `validate:"required" bson:"email"`
}

type Message struct {
	Message string
	Code    int32
}
