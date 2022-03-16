package main

import (
	"fmt"
	"html/template"
	"net/http"
	"regexp"
	"strconv"
	"strings"
)

//端口状态
type ipPortInfo struct {
	ip   string
	port []int
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		//解析模板
		t, _ := template.ParseFiles("layout.html", "Scanner.html")
		//指定执行的模板
		t.ExecuteTemplate(writer, "layout", "")

	})
	http.HandleFunc("/process", func(writer http.ResponseWriter, request *http.Request) {
		request.ParseForm()
		data := request.PostForm
		ip := data.Get("ip")
		portRange := data.Get("portRange")
		m, _ := regexp.MatchString("((2(5[0-5]|[0-4]\\d))|[0-1]?\\d{1,2})(\\.((2(5[0-5]|[0-4]\\d))|[0-1]?\\d{1,2})){3}", ip)
		pRm, _ := regexp.MatchString("[0-9]?\\d{1,4}-[0-9]?\\d{1,4}", portRange)
		if m {
			//ip正确
			if pRm {
				//分割端口
				arr := strings.Split(portRange, "-")
				fmt.Println(arr[0])
				firstPort, _ := strconv.Atoi(arr[0])
				lastPort, _ := strconv.Atoi(arr[1])
				if firstPort > 65535 || firstPort < 0 || lastPort > 65535 || lastPort < 0 || firstPort > lastPort {
					fmt.Fprintln(writer, portRange+"端口范围错误")
				} else {
					/*for _, port := range Scanner(ip) {
						fmt.Fprintln(writer, string(port)+"端口打开了")
					}*/

					t, _ := template.ParseFiles("layout.html", "Scanner.html")

					//指定执行的模板
					//writer.Write([]byte(ip + "已开启的端口"))
					insertDB(ip, Scanner(ip))
					t.ExecuteTemplate(writer, "layout", strings.Split(selectDB(ip), "-"))

				}
			} else {
				fmt.Fprintln(writer, portRange+"端口范围错误")
			}
		} else {
			fmt.Fprintln(writer, ip+"不是一个正确的ip")
		}

	})
	http.Handle("/img/", http.StripPrefix("/img/", http.FileServer(http.Dir("img/"))))
	server.ListenAndServe()
}
