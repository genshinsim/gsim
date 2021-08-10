package blackcliff

import (
	"fmt"

	"github.com/genshinsim/gsim/pkg/combat"
	"github.com/genshinsim/gsim/pkg/core"
)

func init() {
	combat.RegisterWeaponFunc("blackcliff warbow", weapon)
	combat.RegisterWeaponFunc("blackcliff slasher", weapon)
	combat.RegisterWeaponFunc("blackcliff agate", weapon)
	combat.RegisterWeaponFunc("blackcliff pole", weapon)
	combat.RegisterWeaponFunc("blackcliff longsword", weapon)
}

//After defeating an enemy, ATK is increased by 12/15/18/21/24% for 30s.
//This effect has a maximum of 3 stacks, and the duration of each stack is independent of the others.
func weapon(c core.Character, s core.Sim, log core.Logger, r int, param map[string]int) {

	atk := 0.09 + float64(r)*0.03
	index := 0
	stacks := []int{-1, -1, -1}

	m := make([]float64, core.EndStatType)
	c.AddMod(core.CharStatMod{
		Key: "blackcliff",
		Amount: func(a core.AttackTag) ([]float64, bool) {
			count := 0
			for _, v := range stacks {
				if v > s.Frame() {
					count++
				}
			}
			m[core.ATKP] = atk * float64(count)
			return m, true
		},
		Expiry: -1,
	})

	s.AddOnTargetDefeated(func(t core.Target) {
		stacks[index] = s.Frame() + 1800
		index++
		if index == 3 {
			index = 0
		}
	}, fmt.Sprintf("blackcliff-%v", c.Name()))
}
