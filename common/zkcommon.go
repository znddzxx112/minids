package common

import (
	"errors"
	"fmt"
	"github.com/samuel/go-zookeeper/zk"
	"math/rand"
	"time"
)

func GetServerList(serviceName string) (string, error) {
	conn, err := GetConnect([]string{"zookeeper"}, time.Second)
	if err != nil {
		panic("zookeeper connect failed")
	}
	defer conn.Close()

	serverList, getErr := GetServerListAll(conn, serviceName)
	if getErr != nil {
		panic("get serverList failed")
	}
	count := len(serverList)
	if count == 0 {
		return "", errors.New("serverList is empty")
	}

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	serverIp := serverList[r.Intn(count)]
	return serverIp, nil
}

func RegisterServiceName(serviceName string, serverIPAndPort string) *zk.Conn {
	conn, zkconnError := GetConnect([]string{"zookeeper"}, time.Second)
	if zkconnError != nil {
		panic("zookeeper connect failed")
	}

	registerErr := RegistServer(conn, serviceName, serverIPAndPort)
	if registerErr != nil {
		panic("zookeeper register server failed")
	}

	return conn
}

func GetConnect(hosts []string, timeOut time.Duration) (conn *zk.Conn, err error) {
	conn, _, err = zk.Connect(hosts, timeOut)
	if err != nil {
		fmt.Println(err)
	}
	return
}

// /**
// jiqun manage
// linshi jiedian
func RegistServer(conn *zk.Conn, serviceName string, host string) (err error) {
	// flags 0,1,2
	_, err = conn.Create("/"+serviceName+"/"+host, nil, zk.FlagEphemeral, zk.WorldACL(zk.PermAll))
	return
}

func GetServerListAll(conn *zk.Conn, serviceName string) (list []string, err error) {
	list, _, err = conn.Children("/" + serviceName)
	return
}

// **/

// /**
// peizhi manage
// yongjiu jiedian
func PutVal(conn *zk.Conn, key string, val string) error {
	_, err := conn.Create("/configManagePlatform/"+key, []byte(val), 0, zk.WorldACL(zk.PermAll))
	return err
}

func GetVal(conn *zk.Conn, key string) (string, error) {
	byte, _, err := conn.Get("/configManagePlatform/" + key)
	return string(byte[:]), err
}

// **/
