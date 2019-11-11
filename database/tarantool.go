package database

import (
	"fmt"
	"github.com/tarantool/go-tarantool"
	"github.com/tarantool/go-tarantool/queue"
	"log"
	"os"
	"sync"
	"tarantoolTest/models"
	"time"
)

var TrConns []*TarantoolConnect

type TarantoolConnect struct {
	sync.RWMutex
	Host   string
	Port   string
	User   string
	Pass   string
	client *tarantool.Connection
	queue  queue.Queue
}

var cfg = queue.Cfg{
	Temporary:   true,
	IfNotExists: true,
	Kind:        queue.FIFO,
	Opts: queue.Opts{
		Ttl:   10 * time.Second,
		Ttr:   5 * time.Second,
		Delay: 3 * time.Second,
		Pri:   1,
	},
}

func TarantoolInit() {
	index := 1
	for {
		host := os.Getenv(fmt.Sprintf("TARANTOOL_HOST_%d", index))
		if len(host) == 0 {
			break
		}
		tr := TarantoolConnect{
			Host: host,
			Port: os.Getenv(fmt.Sprintf("TARANTOOL_PORT_%d", index)),
			User: os.Getenv(fmt.Sprintf("TARANTOOL_USER_%d", index)),
			Pass: os.Getenv(fmt.Sprintf("TARANTOOL_PASS_%d", index)),
		}
		TrConns = append(TrConns, &tr)
		index++
	}
}

func (conn TarantoolConnect) Connect() *tarantool.Connection {
	conn.Lock()
	defer conn.Unlock()

	if conn.client == nil {
		var err error
		conn.client, err = tarantool.Connect(
			fmt.Sprintf(`%s:%s`, conn.Host, conn.Port),
			tarantool.Opts{
				Timeout:       time.Minute,
				Reconnect:     time.Second,
				MaxReconnects: 5,
				User:          conn.User,
				Pass:          conn.Pass,
			})

		if err != nil {
			log.Fatal(err.Error())
		}
		return conn.client
	}

	return conn.client
}

func (conn TarantoolConnect) Disconnect() {
	conn.client.Close()
}

func InsertOrUpdate(conn *tarantool.Connection, table string, item interface{}) {
	switch table {
	case "fulltraffic":
		item = item.(models.FullTraffic)
	case "clicks":
		item = item.(models.Click)
	case "breaks":
		item = item.(models.Breaking)
	case "postbacks":
		item = item.(models.PostBack)
	}

	_, err := conn.Call17("InsertOrUpdate", []interface{}{item, table})
	if err != nil {
		fmt.Println(err.Error())
	}
}

func DropTraffList(conn *tarantool.Connection, list []uint, table string) {
	_, err := conn.Call17("DropTraffList", []interface{}{list, table})
	if err != nil {
		fmt.Println(err.Error())
	}
}

func GetTrafficIVCodes(conn *tarantool.Connection) ([]models.FullTraffic, []uint) {
	var tr []models.FullTraffic
	var ivcodes []uint
	err := conn.SelectTyped("fulltraffic", "secondary", 0, 4294967295, tarantool.IterLe, []interface{}{int(time.Now().Add(-5 * time.Minute).Unix())}, &tr)
	if err != nil {
		fmt.Println(err.Error())
	}

	for _, v := range tr {
		ivcodes = append(ivcodes, v.IVCode)
	}
	return tr, ivcodes
}

func GetClicksIVCodes(conn *tarantool.Connection) ([]models.Click, []uint) {
	var tr []models.Click
	var ivcodes []uint
	err := conn.SelectTyped("clicks", "secondary", 0, 4294967295, tarantool.IterLe, []interface{}{int(time.Now().Add(-5 * time.Minute).Unix())}, &tr)
	if err != nil {
		fmt.Println(err.Error())
	}

	for _, v := range tr {
		ivcodes = append(ivcodes, v.IVCode)
	}
	return tr, ivcodes
}

func GetBreaksIVCodes(conn *tarantool.Connection) ([]models.Breaking, []uint) {
	var tr []models.Breaking
	var ivcodes []uint
	err := conn.SelectTyped("breaks", "secondary", 0, 4294967295, tarantool.IterLe, []interface{}{int(time.Now().Add(-5 * time.Minute).Unix())}, &tr)
	if err != nil {
		fmt.Println(err.Error())
	}

	for _, v := range tr {
		ivcodes = append(ivcodes, v.IVCode)
	}
	return tr, ivcodes
}

func GetPostBacksIVCodes(conn *tarantool.Connection) ([]models.PostBack, []uint) {
	var tr []models.PostBack
	var ivcodes []uint
	err := conn.SelectTyped("post_backs", "secondary", 0, 4294967295, tarantool.IterLe, []interface{}{int(time.Now().Add(-5 * time.Minute).Unix())}, &tr)
	if err != nil {
		fmt.Println(err.Error())
	}

	for _, v := range tr {
		ivcodes = append(ivcodes, v.IVCode)
	}
	return tr, ivcodes
}

//func (conn TarantoolConnect) PutPbInQueue(name string, pb models.PostBack) bool{
//	var connect *tarantool.Connection
//	connect = conn.client
//	if conn.client == nil {
//		connect = conn.Connect()
//	}
//	if conn.queue == nil {
//		conn.queue = queue.New(connect, name)
//		if err := conn.queue.Create(cfg); err != nil {
//			log.Fatalf("queue create: %s", err)
//			return false
//		}
//	}
//	_, err := conn.queue.Put(&pb)
//	if err != nil {
//		log.Fatalf("put typed task: %s", err)
//	}
//	return true
//}
//
//func (conn TarantoolConnect) TakePbFromQueue(name string) (models.PostBack,error){
//	var pb models.PostBack
//	_, err := conn.queue.TakeTyped(&pb) //blocking operation
//	if err != nil {
//		log.Fatalf("take take typed: %s", err)
//		return models.PostBack{}, err
//	}
//	return pb, nil
//}