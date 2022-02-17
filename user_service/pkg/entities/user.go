package entities

type User struct {
	Id                    int32
	PwdHash               string `validate:"required"`
	Name                  string `validate:"required"`
	Age                   int32  `validate:"required"`
	AdditionalInformation string `validate:"required"`
	Parent                []string
	Email                 string `validate:"required"`
}

type Message struct {
	Message string
	Code    int32
}
