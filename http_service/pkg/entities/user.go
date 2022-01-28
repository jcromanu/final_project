package entities

type User struct {
	Id                     int32    `json:"user_id"`
	Pwd_hash               string   `json:"pwd_hash" validate:"required"`
	Name                   string   `json:"name" validate:"required"`
	Age                    int32    `json:"age" validate:"required"`
	Additional_information string   `json:"additional_information" validate:"required"`
	Parent                 []string `json:"parent"`
}

type Message struct {
	Message string
	Code    int32
}
