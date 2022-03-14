package royal

import (
	"github.com/genshinsim/gcsim/internal/weapons/common"
	"github.com/genshinsim/gcsim/pkg/core"
	"github.com/genshinsim/gcsim/pkg/coretype"
)

func init() {
	core.RegisterWeaponFunc("royalgreatsword", weapon)
}

func weapon(char coretype.Character, c *core.Core, r int, param map[string]int) string {
	common.Royal(char, c, r, param)
	return "royalgreatsword"
}
