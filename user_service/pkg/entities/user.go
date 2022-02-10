package entities

type User struct {
	Id                    int32
	PwdHash               string `validate:"required"`
	Name                  string `validate:"required"`
	Age                   int32  `validate:"required"`
	AdditionalInformation string `validate:"required"`
	Parent                []string
}

type Message struct {
	Message string
	Code    int32
}
