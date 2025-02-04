package basic

import (
	"log"
	"my-protobuf/protogen/basic"

	"google.golang.org/protobuf/encoding/protojson"
)

func BasicUser() {
	addr := basic.Address{
		Street:  "Daily Planet",
		City:    "Metropolis",
		Country: "US",
		Coordinate: &basic.Address_Coordinate{
			Latitude:  4070797893425118,
			Longitude: -74.01163838107261,
		},
	}

	u := basic.User{
		Id:       99,
		Username: "superman",
		IsActive: true,
		Password: []byte("supermanpassword"),
		// Emails:   []string{"superman@movie.com", "superman@dc.com"},
		// Gender:   basic.Gender_GENDER_MALE,
		Address: &addr,
	}

	jsonBytes, _ := protojson.Marshal(&u)
	log.Println(string(jsonBytes))
}

func ProtoToJsonUser() {
	p := basic.User{
		Id:       98,
		Username: "wonderwoman",
		IsActive: true,
		Password: []byte("wonderwomanpassword"),
		Emails:   []string{"wonderwoman@movie.com", "wonderwoman@dc.com"},
		Gender:   basic.Gender_GENDER_FEMALE,
	}

	jsonBytes, _ := protojson.Marshal(&p)
	log.Println(string(jsonBytes))
}

func JsonToProtoUser() {
	json := `
	{
		"id":98, 
		"username":"wonderwoman", 
		"is_active":true, 
		"password":"d29uZGVyd29tYW5wYXNzd29yZA==", 
		"emails":[
			"wonderwoman@movie.com", 
			"wonderwoman@dc.com"
		], 
		"gender":"GENDER_FEMALE"}
	`
	var p basic.User

	err := protojson.Unmarshal([]byte(json), &p)

	if err != nil {
		log.Fatalln("Got error : ", err)
	}

	log.Println(&p)
}
