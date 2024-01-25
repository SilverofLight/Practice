package main
import (
	"fmt"
	"net"
)

func main(){
	fmt.Println("服务端启动。。。")
	//进行监听,指定tcp协议，ip和端口号PORT
	listen, err := net.Listen("tcp","127.0.0.1:8888")
	if err != nil {
		fmt.Println("监听失败，err: ", err)
		return
	}
	//等待连接
	for {
		conn,err2 := listen.Accept()
		if err2 != nil {
			fmt.Println("等待连接失败，err: ", err)
		}else {
			fmt.Printf("等待连接成功，conn: %v, 接受到的客户端信息： %v\n", conn, conn.RemoteAddr().String())
		}
		//用协成处理请求
		go process(conn)//不同客户端的请求，连接conn不同
	}
	
}

func process(conn net.Conn){
	//连接用完要关闭
	defer conn.Close()
	
	for {
		//创建切片，将读取的数据放入切片
		buf := make([]byte,1024)

		//从conn读取数据
		n,err := conn.Read(buf)
		if err != nil {
			fmt.Println("err: ",err)
			return
		}
		//打印在服务器
		fmt.Println(string(buf[0:n]))
	}
}