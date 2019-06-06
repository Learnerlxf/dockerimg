package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"sync"
	"websocket"
)

/////////////////////////http/////////////////////////
func HandHello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("wecome connector"))
}

func HandGetUserList(w http.ResponseWriter, r *http.Request) {
	resp, _ := json.Marshal(userMap)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(resp))
}

func HandTranspond(w http.ResponseWriter, bodyByte []byte) {
	//bodyByte, err := ioutil.ReadAll(r.Body)

	//if err != nil {
	//	w.Write([]byte("读取body失败"))
	//	return
	//}
	type Msg struct {
		Uid     uint64 `json:"uid"`
		Content string `json:"content"`
	}
	var msg Msg
	err := json.Unmarshal(bodyByte, &msg)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	if c, ok := userMap[msg.Uid]; ok {
		c.WriteMessage(1, []byte(msg.Content))
		w.Write([]byte("转发成功"))
		return
	}
	w.Write([]byte("转发失败，没有该用户"))
}

/////////////////////////websocket/////////////////////////
var addr = flag.String("addr", ":9901", "http service address")
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
} // use default options
type ReqMsg struct {
	Id  uint64 `json:"id"`
	Msg string `json:"msg"`
}
type ResMsg struct {
	Id  uint64 `json:"id"`
	Msg string `json:"msg"`
}

var userMap map[uint64]*websocket.Conn
var userMapRwmutex sync.RWMutex

func AddUser(uid uint64, c *websocket.Conn) {
	userMapRwmutex.Lock()
	defer userMapRwmutex.Unlock()
	userMap[uid] = c
}

func RmUser(uid uint64) {
	userMapRwmutex.Lock()
	defer userMapRwmutex.Unlock()
	delete(userMap, uid)
}

func onWsConnect(w http.ResponseWriter, r *http.Request) {
	bodyByte, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	type Msg struct {
		Type    uint64 `json:"type"`
		Uid     uint64 `json:"uid"`
		Content string `json:"content"`
	}
	var msg Msg
	err = json.Unmarshal(bodyByte, &msg)
	if err != nil {
		if string(bodyByte) == "http" {
			w.Write([]byte("hello http"))
			return
		}
		goto upgreadeWs
	}
	switch msg.Type {
	case 1:
		HandHello(w, r)
		return
	case 2:
		HandGetUserList(w, r)
		return
	case 3:
		HandTranspond(w, bodyByte)
		return
	}

upgreadeWs:

	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}
	defer c.Close()

	var uid uint64
	for {
	reBegin:
		mt, message, err := c.ReadMessage()
		if err != nil {
			RmUser(uid)
			log.Println("read:", err)
			break
		}
		var reqMsg ReqMsg
		var resMsg ResMsg
		err = json.Unmarshal(message, &reqMsg)
		if err != nil {
			log.Println(err)
			goto reBegin
		}
		resMsg.Id = reqMsg.Id
		switch reqMsg.Id {
		case 1:
			uid = Login(c, &reqMsg)
			resMsg.Msg = strconv.Itoa(int(uid))
		case 2:

		}
		resByte, _ := json.Marshal(resMsg)
		err = c.WriteMessage(mt, resByte)
		if err != nil {
			log.Println("write:", err)
			break
		}
	}
}
func Login(c *websocket.Conn, msg *ReqMsg) uint64 {
	uid, err := strconv.Atoi(msg.Msg)
	if err != nil {
		return 0
	}
	AddUser(uint64(uid), c)
	return uint64(uid)
}

func WsServer() {
	flag.Parse()
	log.SetFlags(0)
	http.HandleFunc("/ws", onWsConnect)
	http.HandleFunc("/ws/hello", HandHello)
	log.Fatal(http.ListenAndServe(*addr, nil))
}

func main() {
	userMap = make(map[uint64]*websocket.Conn, 999)
	go func() {
		WsServer()
	}()
	select {}
}

