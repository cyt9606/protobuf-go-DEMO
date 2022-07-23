package main

import (
	"encoding/json"
	"fmt"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
	"io/ioutil"
	"os"
	"protobuf/example"
)

func main(){

	proto_1_1:=&example.Result{
		Url: "www.com",
		Title: "bbbaaa",
	}
	proto_1:=&example.Person{
		Name: "david",
		Age: 18,
		From: "China",
		Result: proto_1_1,
	}

	//proto序列化
	byte1,err:=proto.Marshal(proto_1)
	if err!=nil{panic(err)}
	fmt.Println(byte1)
	proto_2:=new(example.Person)
	proto.Unmarshal(byte1,proto_2)	//反序列化
	fmt.Println(proto.Equal(proto_1,proto_2))

	//proto转json
	json_byte,err:=protojson.Marshal(proto_1)
	fmt.Println(string(json_byte))
	json_1:=new(example.Person)
	json.Unmarshal(json_byte,json_1)
	fmt.Println(json_1.Name)

	//json转proto
	jsonfile,err:=os.Open("example.json")
	if err!=nil{panic(err)}
	json_byte2,err:=ioutil.ReadAll(jsonfile)
	if err!=nil{panic(err)}
	proto_3:=new(example.Person)
	protojson.Unmarshal(json_byte2,proto_3)
	fmt.Println(protojson.Format(proto_3))
}