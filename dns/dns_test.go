package dns

import (
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
)

func TestCNAME(t *testing.T) {
	src := "blog.bysir.top."
	dst,err := CNAME(src)
	assert.Nil(t,err,"CNAME")
	t.Log(dst)
}


func TestCNAME2(t *testing.T) {
	code := "hk00048"
	a,err := strconv.ParseInt(code[2:],10,64)
	assert.Nil(t,err,"")
	t.Log(a)
}