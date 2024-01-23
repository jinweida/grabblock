package common

import (
	"fmt"
	"github.com/rs/xid"
)

func GenXid() string{
	id := xid.New()
	fmt.Printf("github.com/rs/xid:           %s\n", id.String())
	return id.String()
}
