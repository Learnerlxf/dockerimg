package main

import (
	"fmt"
	"net/http"
	"os"
	"runtime"
	"strconv"
	"time"
)


/////////////////////////http/////////////////////////
func HandHello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("wecome shop"))

	fmt.Printf("client ip: %v \n", r.Header.Get("X-Forwarded-For"))
}

func HandHostInfo(w http.ResponseWriter, r *http.Request) {
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

func HttpServer() {
	http.HandleFunc("/", HandHello)
	http.HandleFunc("/hostInfo", HandHostInfo)
	http.ListenAndServe(":8000", nil)
	fmt.Println("service time end")
}



func main() {
	go func() {
		HttpServer()
	}()
	select {}
}
