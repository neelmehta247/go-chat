package main

import (
	"fmt"

	"github.com/neelmehta247/go-chat/proto"
)

func main() {
	msg := &proto.Hello{
		Message: "hello",
	}

	fmt.Println(msg.Message)
}
