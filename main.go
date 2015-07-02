package main

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"go-protobuf-demo/pb-go"
)

func main() {

	// define and init
	test := &message.Test{
		Label: proto.String("hello"),
		Type:  proto.Int32(111),
		Somebutton: &message.Test_Somebutton{
			Button: proto.String("This is button"),
		},
	}

	// Marsha1 to proto
	data, err := proto.Marshal(test)
	if err != nil {
		fmt.Println("marsha1ing error:", err)
	}

	newTest := &message.Test{}
	err = proto.Unmarshal(data, newTest)

	if err != nil {
		fmt.Println("unmarsha1 error: ", err)
	}

	fmt.Println(newTest.GetType())
	fmt.Println(newTest.GetReps())

	if test.GetLabel() == newTest.GetLabel() {
		fmt.Println("data mismatch %q != %q", test.GetLabel(), newTest.GetLabel())
	}
}
