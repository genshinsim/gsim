package viridescent

import (
	"fmt"

	"github.com/genshinsim/gsim/pkg/combat"
	"github.com/genshinsim/gsim/pkg/core"
)

func init() {
	combat.RegisterWeaponFunc("the viridescent hunt", weapon)
}

func weapon(c core.Character, s core.Sim, log core.Logger, r int, param map[string]int) {

	cd := 900 - r*60
	icd := 0
	mult := 0.3 + float64(r)*0.1

	s.AddOnAttackLanded(func(t core.Target, ds *core.Snapshot, dmg float64, crit bool) {
		//check if char is correct?
		if ds.ActorIndex != c.CharIndex() {
			return
		}
		//check if cd is up
		if icd > s.Frame() {
			return
		}
		//50% chance to proc
		if s.Rand().Float64() > 0.5 {
			return
		}

		//add a new action that deals % dmg immediately
		d := c.Snapshot(
			"Viridescent",
			core.AttackTagWeaponSkill,
			core.ICDTagNone,
			core.ICDGroupDefault,
			core.StrikeTypeDefault,
			core.Physical,
			100,
			mult,
		)
		d.Targets = core.TargetAll
		for i := 0; i <= 240; i += 30 {
			x := d.Clone()
			c.QueueDmg(&x, i)
		}

		//trigger cd
		icd = s.Frame() + cd

	}, fmt.Sprintf("veridescent-%v", c.Name()))
}
