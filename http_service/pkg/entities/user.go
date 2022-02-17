package entities

type User struct {
	Id                    int32    `json:"user_id"`
	PwdHash               string   `json:"pwd_hash,omitempty" validate:"required"`
	Name                  string   `json:"name" validate:"required"`
	Age                   int32    `json:"age" validate:"required"`
	AdditionalInformation string   `json:"additional_information" validate:"required"`
	Parent                []string `json:"parent"`
	Email                 string   `json:"email" validate:"required"`
}

type Message struct {
	Message string
	Code    int32
}
