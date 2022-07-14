package heizou

import "github.com/genshinsim/gcsim/pkg/core"

//For 5s after Shikanoin Heizou takes the field, his Normal Attack SPD is increased by 15%.
//He also gains 1 Declension stack for Heartstopper Strike. This effect can be triggered once every 10s.
func (c *char) c1() {
	// Add hook that monitors for crit hits. Mirrors existing favonius code
	// No log value saved as stat mod already shows up in debug view
	c.Core.Events.Subscribe(core.OnCharacterSwap, func(args ...interface{}) bool {
		if c.c1icd > c.Core.F {
			return false
		}
		prev := args[0].(int)
		next := args[1].(int)
		if next == c.Index && prev != c.Index {
			val := make([]float64, core.EndStatType)
			val[core.AtkSpd] = 0.15
			c.AddPreDamageMod(core.PreDamageMod{
				Key:    "heizou-c1",
				Expiry: c.Core.F + 240,
				Amount: func(atk *core.AttackEvent, t core.Target) ([]float64, bool) {
					if atk.Info.AttackTag != core.AttackTagNormal {
						return nil, false
					}
					return val, true
				},
			})
			c.addDecStack()
			c.c1icd = c.Core.F + 600
		}

		return false
	}, "heizou enter")

}

//The first Windmuster Iris explosion in each Windmuster Kick will regenerate 9 Elemental Energy for Shikanoin Heizou.
//Every subsequent explosion in that Windmuster Kick will each regenerate an additional 1.5 Energy for Heizou.
//One Windmuster Kick can regenerate a total of 13.5 Energy for Heizou in this manner.
func (c *char) c4(i int) {
	energy := 0.0
	switch i {
	case 1:
		energy += 9.5
	case 2, 3:
		energy += 1.5
	case 4:
		energy += 1.0
	}
	c.AddEnergy("heizou c4", energy)
}

//Each Declension stack will increase the CRIT Rate of the Heartstopper Strike unleashed by 4%.
//When Heizou possesses Conviction, this Heartstopper Strike's CRIT DMG is increased by 32%.

func (c *char) c6(snap *core.Snapshot) {
	snap.Stats[core.CR] += float64(c.decStack) * 0.04
	if c.decStack == 4 {
		snap.Stats[core.CD] += 0.32
	}
}
