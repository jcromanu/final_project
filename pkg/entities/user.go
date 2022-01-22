package entities

type User struct {
	Id                     int32
	Pwd_hash               string
	Name                   string
	Age                    int32
	Additional_information string
	Parent                 []string
}

type Message struct {
	Message string
	Code    int32
}
