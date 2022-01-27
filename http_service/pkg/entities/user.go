package entities

type User struct {
	Id                     int32    `json:"user_id"`
	Pwd_hash               string   `json:"pwd_hash"`
	Name                   string   `json:"name"`
	Age                    int32    `json:"age"`
	Additional_information string   `json:"additional_information"`
	Parent                 []string `json:"parent"`
}

type Message struct {
	Message string
	Code    int32
}
