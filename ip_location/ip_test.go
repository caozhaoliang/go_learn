package ip_location

import (
	"fmt"
	"github.com/wangtuanjie/ip17mon"
	"testing"
)

func init() {
	ip17mon.Init("./17monipdbNew.dat")
}

func TestIP(t *testing.T) {
	loc, err := ip17mon.Find("116.228.111.18")
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	fmt.Println(loc)
}
