package main

import (
	"fmt"
	"net"
	"sort"
	"strconv"
)

func worker(ports chan int, results chan int, ip string) {
	for p := range ports {
		address := fmt.Sprintf("%s:%d", ip, p)
		conn, err := net.Dial("tcp", address)
		if err != nil {
			results <- 0
			continue
		}
		conn.Close()
		results <- p

	}
}

//main相当于goroutine
func Scanner(ip string) string {
	//chan缓冲100个数据worker，不必等待里面有内容
	ports := make(chan int, 100)
	//main函数用的goroutine不需要缓冲
	results := make(chan int)
	var openport []int
	for i := 0; i < cap(ports); i++ {
		go worker(ports, results, ip)
	}
	//分配工作必须在，接收数据之前运行，所以必须开启一个goroutine
	go func() {
		for i := 1; i < 1024; i++ {
			ports <- i
		}
	}()

	for i := 1; i < 1024; i++ {
		//从通道接收数据
		port := <-results
		if port != 0 {
			openport = append(openport, port)
		}

	}
	close(ports)
	close(results)
	sort.Ints(openport)
	i := ""
	for _, port := range openport {
		i += strconv.Itoa(port) + "-"
		fmt.Printf("%d 端口打开了\n", port)
	}
	return i
}
