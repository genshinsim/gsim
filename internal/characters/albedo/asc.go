package albedo

import (
	"github.com/genshinsim/gcsim/pkg/core/attributes"
	"github.com/genshinsim/gcsim/pkg/core/combat"
	"github.com/genshinsim/gcsim/pkg/core/player/character"
	"github.com/genshinsim/gcsim/pkg/enemy"
	"github.com/genshinsim/gcsim/pkg/modifier"
)

// Transient Blossoms generated by Abiogenesis: Solar
// Isotoma deal 25% more DMG to opponents whose HP is below 50%.
func (c *char) a1() {
	if !c.Core.Combat.DamageMode {
		return
	}
	m := make([]float64, attributes.EndStatType)
	m[attributes.DmgP] = 0.25
	c.AddAttackMod(character.AttackMod{
		Base: modifier.NewBase("albedo-a1", -1),
		Amount: func(atk *combat.AttackEvent, t combat.Target) ([]float64, bool) {
			if atk.Info.AttackTag != combat.AttackTagElementalArt {
				return nil, false
			}
			// Can't be triggered by itself when refreshing
			if atk.Info.Abil == "Abiogenesis: Solar Isotoma" {
				return nil, false
			}
			if e, ok := t.(*enemy.Enemy); !(ok && e.HP()/e.MaxHP() < .5) {
				return nil, false
			}
			return m, true
		},
	})
}
