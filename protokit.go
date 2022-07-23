package main

import(
	"google.golang.org/protobuf/encoding/prototext"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/encoding/protojson"
)

func ProtoToJson(m proto.Message) ([]byte,error){
	return protojson.Marshal(m)
}

func JsonToProto(b []byte,r proto.Message)(error){
	return protojson.Unmarshal(b,r)
}

func  ProtoToByte(m proto.Message)([]byte,error){
	return  proto.Marshal(m)
}

func ByteToProto(b []byte,m proto.Message)(err error){
	return proto.Unmarshal(b,m)
}

func ProtoToText(m proto.Message)string{
	return prototext.Format(m)
}

