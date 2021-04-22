package ip_location

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/wangtuanjie/ip17mon"
	"net"
	"strconv"
	"strings"
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
// ip 转成int类型存储 节省空间 同时如果ip为索引时查询效率也会高一些
func ParseIp(ip string) (uint32, error) {
	objIp := net.ParseIP(ip)
	if objIp == nil {
		return 0, fmt.Errorf("IP 解析错误[%s]", ip)
	}
	var iIp uint32
	iIp |= uint32(objIp[12]) << 24
	iIp |= uint32(objIp[13]) << 16
	iIp |= uint32(objIp[14]) << 8
	iIp |= uint32(objIp[15])
	return iIp, nil
}
func UintIPV4String(ip uint32) (string, error) {
	ipObj := make(net.IP, net.IPv4len)
	ipObj[0] = byte(ip >> 24)
	ipObj[1] = byte(ip >> 16)
	ipObj[2] = byte(ip >> 8)
	ipObj[3] = byte(ip)
	return ipObj.String(), nil
}

// JWT 解析
func GetUuidByToken(src string) string {
	res := strings.Split(src, ".")
	if len(res) > 2 {
		decodeBytes, _ := base64.StdEncoding.DecodeString(res[1])
		//{"session":"bb89d391e84f4f7f9bb3f15daf3c6354","source":"web","uuid":363711293513183232}
		type result struct {
			//Session string `json:"session"`
			//Source string `json:"source"`
			Uuid int64 `json:"uuid"`
		}
		r := &result{}
		json.Unmarshal(decodeBytes, r)
		return strconv.Itoa(int(r.Uuid))
	} else {
		return ""
	}
}
