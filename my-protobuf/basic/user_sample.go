package basic

import (
	"log"
	"my-protobuf/protogen/basic"
)

func BasicUser() {
	u := basic.User{
		Id:       99,
		Username: "superman",
		IsActive: true,
		Password: []byte("supermanpassword"),
	}

	log.Println(&u)
}
