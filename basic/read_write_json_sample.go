package basic

import (
	"log"
	"my-protobuf/protogen/basic"
	"os"

	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

func WriteProtoToJSON(msg proto.Message, fname string) {
	b, err := protojson.Marshal(msg)

	if err != nil {
		log.Fatalln("Can not marshal message", err)
	}

	if err := os.WriteFile(fname, b, 0644); err != nil {
		log.Fatalln("Can not write to file", err)
	}
}

func ReadProtoFromJSON(fname string, dest proto.Message) {
	in, err := os.ReadFile(fname)

	if err != nil {
		log.Fatalln("Can not read file", err)
	}

	if err := protojson.Unmarshal(in, dest); err != nil {
		log.Fatalln("Can not unmarshal", err)
	}
}

func WriteToJSONSample() {
	u := dummyUser()
	WriteProtoToJSON(&u, "superman_file.json")
}

func ReadFromJSONSample() {
	dest := basic.User{}

	ReadProtoFromJSON("superman_file.json", &dest)

	log.Println(&dest)
}
