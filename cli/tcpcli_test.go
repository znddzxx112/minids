package client

import (
	"bufio"
	"fmt"
	"os"
	"testing"
)

func TestTcpcliProcess(t *testing.T)  {

}

func TestOsstdin(t *testing.T)  {
	fmt.Print(">")
	rd := bufio.NewReader(os.Stdin)
	s1, _ := rd.ReadString('\n')
	fmt.Println(s1)
}