package consts

import (
	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/hypersdk/consts"
)

const (
	Name = "hypertimestamp"
)

var ID ids.ID

func init() {
	b := make([]byte, consts.IDLen)
	copy(b, []byte(Name))
	vmId, err := ids.ToID(b)
	if err != nil {
		panic(err)
	}
	ID = vmId
}
