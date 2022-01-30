package main

import (
	"example/gthomework/work6/common"
	"fmt"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", ":8899")
	if err != nil {
		fmt.Println(err)
		return
	}
	clntFd := common.SocketUtil{conn}
	for i := 0; i < 100; i++ {
		data := fmt.Sprintf(`{"index":%d, "name":"zhangsan", "age":21, "company":"bat"}`, i+1)
		head := common.GoimPkgHeader{
			HeaderFlag:      common.HeaderFlag,
			PackageLength:   uint32(common.HeaderLength) + uint32(len([]byte(data))),
			HeaderLength:    common.HeaderLength,
			ProtocolVersion: 1,
			Operation:       1,
			SequenceId:      uint32(i),
		}
		n, err := clntFd.WritePkg(head, []byte(data))
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf("Send %d byte data : %s", n, data)
	}
}
