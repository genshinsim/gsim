package ironsting

import (
	"fmt"

	"github.com/genshinsim/gcsim/pkg/core"
	"github.com/genshinsim/gcsim/pkg/coretype"
)

func init() {
	core.RegisterWeaponFunc("iron sting", weapon)
	core.RegisterWeaponFunc("ironsting", weapon)
}

//After using an Elemental Skill, increases Normal and Charged Attack DMG by 8% for 12s. Max 2 stacks.
func weapon(char coretype.Character, c *core.Core, r int, param map[string]int) string {

	expiry := 0
	atk := 0.045 + 0.015*float64(r)
	stacks := 0
	icd := 0

	c.Subscribe(coretype.OnDamage, func(args ...interface{}) bool {

		atk := args[1].(*coretype.AttackEvent)

		if atk.Info.ActorIndex != char.Index() {
			return false
		}
		if atk.Info.Element == core.Physical {
			return false
		}
		if icd > c.Frame {
			return false
		}
		icd = c.Frame + 60
		if expiry < c.Frame {
			stacks = 0
		}
		stacks++
		if stacks > 2 {
			stacks = 2
		}
		expiry = c.Frame + 360
		return false
	}, fmt.Sprintf("ironsting-%v", char.Name()))

	val := make([]float64, core.EndStatType)
	char.AddMod(coretype.CharStatMod{
		Key:    "ironsting",
		Expiry: -1,
		Amount: func() ([]float64, bool) {
			if expiry < c.Frame {
				stacks = 0
				return nil, false
			}
			val[core.DmgP] = atk * float64(stacks)
			return val, true
		},
	})
	return "ironsting"
}
