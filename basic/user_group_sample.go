package basic

import (
	"log"
	"my-protobuf/protogen/basic"

	"google.golang.org/protobuf/encoding/protojson"
)

func BasicUserGroup() {
	batman := basic.User{
		Id:       97,
		Username: "batman",
		IsActive: true,
		Password: []byte("batmanpassword"),
		Gender:   basic.Gender_GENDER_MALE,
	}

	nightwing := basic.User{
		Id:       96,
		Username: "nightwing",
		IsActive: true,
		Password: []byte("nightwingpassword"),
		Gender:   basic.Gender_GENDER_MALE,
	}

	robin := basic.User{
		Id:       95,
		Username: "robin",
		IsActive: true,
		Password: []byte("robinpassword"),
		Gender:   basic.Gender_GENDER_MALE,
	}

	batFamilay := basic.UserGroup{
		GroupId:     999,
		GroupName:   "Bat Familay",
		Users:       []*basic.User{&batman, &nightwing, &robin},
		Description: "Thhe classic bat family",
	}

	jsonBytes, _ := protojson.Marshal(&batFamilay)
	log.Println(string(jsonBytes))
}
