package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"

	"google.golang.org/protobuf/proto"
	"movieexample.com/gen"
	model "movieexample.com/metadata/pkg"
)

var metadata = &model.Metadata{
	ID:          "123",
	Title:       "Movie 123",
	Description: "Movie desc 123",
	Director:    "Mohammad Adnan",
}

var genMetadata = &gen.Metadata{
	Id:          "123",
	Title:       "Movie 123",
	Description: "Movie desc 123",
	Director:    "Mohammad Adnan",
}

func main() {
	jsonBytes, err := serializeToJson(metadata)
	if err != nil {
		panic(err)
	}

	xmlBytes, err := serializeToXml(metadata)
	if err != nil {
		panic(err)
	}

	protoBytes, err := serializeToProto(genMetadata)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Json size: %d \n", len(jsonBytes))
	fmt.Printf("XML size: %d \n", len(xmlBytes))
	fmt.Printf("Proto size: %d \n", len(protoBytes))
}

func serializeToJson(m *model.Metadata) ([]byte, error) {
	return json.Marshal(m)
}

func serializeToXml(m *model.Metadata) ([]byte, error) {
	return xml.Marshal(m)
}

func serializeToProto(m *gen.Metadata) ([]byte, error) {
	return proto.Marshal(m)
}
