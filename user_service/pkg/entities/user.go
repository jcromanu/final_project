package entities

type User struct {
	Id                     int32
	Pwd_hash               string `validate:"required"`
	Name                   string `validate:"required"`
	Age                    int32  `validate:"required"`
	Additional_information string `validate:"required"`
	Parent                 []string
}

type Message struct {
	Message string
	Code    int32
}
