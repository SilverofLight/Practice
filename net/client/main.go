package main
import (
	"fmt"
	"net"
	"bufio"
	"os"
)

func main(){
	//打印
	fmt.Println("客户端启动。。。")
	//调用Dial函数，需要指定tcp协议，ip和端口号
	conn, err := net.Dial("tcp","127.0.0.1:8888")
	if err != nil {
		fmt.Println("连接失败,err: ",err)
		return
	}
	fmt.Println("连接成功 .conn:", conn)

	//向服务端发送数据，然后退出
	reader := bufio.NewReader(os.Stdin)//os.Srdin代表终端的标准输入

	//读取一行输入信息
	str,err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("输入失败，err: ",err)
	}else {
		n,err := conn.Write([]byte(str))
		if err != nil {
			fmt.Println("连接失败，err:", err)
		}else {
			fmt.Printf("发送成功，发送了%d字节的数据，并退出",n)
		}

	}


}