// Code generated by "pipeline"; DO NOT EDIT.
package polarstar

import (
	_ "embed"

	"github.com/genshinsim/gcsim/pkg/model"
	"google.golang.org/protobuf/proto"
)

//go:embed data_gen.pb
var pbData []byte
var base *model.WeaponData

func init() {
	base = &model.WeaponData{}
	err := proto.Unmarshal(pbData, base)
	if err != nil {
		panic(err)
	}
}

func (x *Weapon) Data() *model.WeaponData {
	return base
}
