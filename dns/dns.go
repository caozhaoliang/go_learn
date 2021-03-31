package dns

import "net"

func CNAME(src string) (string,error) {
	dst,err := net.LookupCNAME(src)
	return dst,err
}
