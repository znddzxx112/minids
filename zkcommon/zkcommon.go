package zkcommon

import (
	"fmt"
	"github.com/samuel/go-zookeeper/zk"
	"time"
)

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

func GetServerList(conn *zk.Conn, serviceName string) (list []string, err error) {
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
