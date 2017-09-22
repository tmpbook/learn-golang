package main

import (
	"log"
	// 辅助库
	"github.com/golang/protobuf/proto"
	// test.pb.go 的路径
	"learn-golang/try-proto"
)

func main() { // 创建一个消息 Test
	test := &example.Test{ // 使用辅助函数设置域的值
		Label: proto.String("hello"),
		Type:  proto.Int32(17),
		Optionalgroup: &example.Test_OptionalGroup{
			RequiredField: proto.String("good bye"),
		},
	} // 进行编码
	data, err := proto.Marshal(test)
	if err != nil {
		log.Fatal("marshaling error: ", err)
	} // 进行解码
	newTest := &example.Test{}
	err = proto.Unmarshal(data, newTest)
	if err != nil {
		log.Fatal("unmarshaling error: ", err)
	} // 测试结果
	if test.GetLabel() != newTest.GetLabel() {
		log.Fatalf("data mismatch %q != %q", test.GetLabel(), newTest.GetLabel())
	}
}
