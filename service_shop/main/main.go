package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"runtime"
	"strconv"
	"time"
	"websocket"
)

var reqNum int
var reqNumChan chan int

/////////////////////////http/////////////////////////
func HandHello(w http.ResponseWriter, r *http.Request) {
	reqNumChan <- 1
	WriteReqLogNum()
	w.Write([]byte("wecome shop"))
}

func HandHostInfo(w http.ResponseWriter, r *http.Request) {
	reqNumChan <- 1
	WriteReqLogNum()
	resp := GetHostInfo()
	timeout, err := strconv.Atoi(r.FormValue("timeout"))
	if err == nil {
		time.Sleep(time.Second * time.Duration(timeout))
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(resp))
}

func GetHostInfo() string {
	hostName, _ := os.Hostname()
	version := runtime.Version()
	goos := runtime.GOOS
	goarch := runtime.GOARCH
	numCpu := runtime.NumCPU()
	egid := os.Getegid()
	euid := os.Geteuid()
	gid := os.Getgid()
	pid := os.Getpid()
	ppid := os.Getppid()
	resp := fmt.Sprintf("hostName=%s,version=%s,goos=%s,goarch=%s,numCpu=%d,egid=%d,euid=%d,gid=%d,pid=%d,ppid=%d,time=%s \n",
		hostName, version, goos, goarch, numCpu, egid, euid, gid, pid, ppid, time.Now().Format("2006-01-02 15:04:05"))

	resp = fmt.Sprintf("%s MY_NODE_NAME=%s, MY_POD_NAME=%s, MY_POD_NAMESPACE=%s, MY_POD_ID=%s, MY_POD_SERVICE_ACCOUNT=%s \n",
		resp, os.Getenv("MY_NODE_NAME"), os.Getenv("MY_POD_NAME"), os.Getenv("MY_POD_NAMESPACE"), os.Getenv("MY_POD_ID"), os.Getenv("MY_POD_SERVICE_ACCOUNT"))

	return resp
}

func WriteReqLogNum() {
	if reqNum%1000 == 0 {
		log.Printf("req times = %d \n", reqNum)
	}
}

func HttpServer() {
	http.HandleFunc("/", HandHello)
	http.HandleFunc("/hostInfo", HandHostInfo)
	http.ListenAndServe(":8000", nil)
	fmt.Println("service time end")
}

/////////////////////////websocket/////////////////////////
var addr = flag.String("addr", ":8080", "http service address")
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
} // use default options

func onWsConnect(w http.ResponseWriter, r *http.Request) {
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}
	defer c.Close()
	for {
		mt, message, err := c.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			break
		}
		reqNumChan <- 1
		WriteReqLogNum()
		err = c.WriteMessage(mt, []byte(GetHostInfo()+"  recive"+string(message)))
		if err != nil {
			log.Println("write:", err)
			break
		}
	}
}

func WsServer() {
	flag.Parse()
	log.SetFlags(0)
	http.HandleFunc("/ws", onWsConnect)
	log.Fatal(http.ListenAndServe(*addr, nil))
}

func main() {
	reqNumChan = make(chan int, 10)
	go func() {
		for {
			<-reqNumChan
			reqNum += 1
		}
	}()
	go func() {
		WsServer()
	}()
	go func() {
		HttpServer()
	}()
	select {}
}
