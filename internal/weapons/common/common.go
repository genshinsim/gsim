package common

import (
	"fmt"

	"github.com/genshinsim/gcsim/pkg/core"
	"github.com/genshinsim/gcsim/pkg/coretype"
)

//After defeating an enemy, ATK is increased by 12/15/18/21/24% for 30s.
//This effect has a maximum of 3 stacks, and the duration of each stack is independent of the others.
func Blackcliff(char coretype.Character, c *core.Core, r int, param map[string]int) {

	atk := 0.09 + float64(r)*0.03
	index := 0
	stacks := []int{-1, -1, -1}

	m := make([]float64, core.EndStatType)
	char.AddMod(coretype.CharStatMod{
		Key: "blackcliff",
		Amount: func() ([]float64, bool) {
			count := 0
			for _, v := range stacks {
				if v > c.Frame {
					count++
				}
			}
			m[core.ATKP] = atk * float64(count)
			return m, true
		},
		Expiry: -1,
	})

	c.Subscribe(core.OnTargetDied, func(args ...interface{}) bool {
		stacks[index] = c.Frame + 1800
		index++
		if index == 3 {
			index = 0
		}
		return false
	}, fmt.Sprintf("blackcliff-%v", char.Name()))

}

func Favonius(char coretype.Character, c *core.Core, r int, param map[string]int) {

	p := 0.50 + float64(r)*0.1
	cd := 810 - r*90
	icd := 0
	//add on crit effect
	c.Subscribe(coretype.OnDamage, func(args ...interface{}) bool {
		atk := args[1].(*coretype.AttackEvent)
		crit := args[3].(bool)
		if !crit {
			return false
		}
		if atk.Info.ActorIndex != char.Index() {
			return false
		}
		if c.ActiveChar != char.Index() {
			return false
		}
		if icd > c.Frame {
			return false
		}

		if c.Rand.Float64() > p {
			return false
		}
		c.Log.NewEvent("favonius proc'd", coretype.LogWeaponEvent, char.Index())

		char.QueueParticle("favonius-"+char.Name(), 3, core.NoElement, 80)

		icd = c.Frame + cd

		return false
	}, fmt.Sprintf("favo-%v", char.Name()))

}

func Lithic(char coretype.Character, c *core.Core, r int, param map[string]int) {

	stacks := 0
	val := make([]float64, core.EndStatType)

	c.Subscribe(core.OnInitialize, func(args ...interface{}) bool {
		for _, char := range c.Chars {
			if char.Zone() == core.ZoneLiyue {
				stacks++
			}
		}
		val[core.CR] = (0.02 + float64(r)*0.01) * float64(stacks)
		val[core.ATKP] = (0.06 + float64(r)*0.01) * float64(stacks)
		return true
	}, fmt.Sprintf("lithic-%v", char.Name()))

	char.AddMod(coretype.CharStatMod{
		Key:    "lithic",
		Expiry: -1,
		Amount: func() ([]float64, bool) {
			return val, true
		},
	})
}

func Royal(char coretype.Character, c *core.Core, r int, param map[string]int) {
	stacks := 0

	c.Subscribe(coretype.OnDamage, func(args ...interface{}) bool {
		crit := args[3].(bool)
		if crit {
			stacks = 0
		} else {
			stacks++
			if stacks > 5 {
				stacks = 5
			}
		}
		return false
	}, fmt.Sprintf("royal-%v", char.Name()))

	rate := 0.06 + float64(r)*0.02
	m := make([]float64, core.EndStatType)
	char.AddMod(coretype.CharStatMod{
		Key: "royal",
		Amount: func() ([]float64, bool) {
			m[core.CR] = float64(stacks) * rate
			return m, true
		},
		Expiry: -1,
	})

}

//After damaging an opponent with an Elemental Skill, the skill has a 40/50/60/70/80%
//chance to end its own CD. Can only occur once every 30/26/22/19/16s.
func Sacrificial(char coretype.Character, c *core.Core, r int, param map[string]int) {

	last := 0
	prob := 0.3 + float64(r)*0.1
	cd := (34 - r*4) * 60

	if r >= 4 {
		cd = (19 - (r-4)*3) * 60
	}
	//add on crit effect
	c.Subscribe(coretype.OnDamage, func(args ...interface{}) bool {
		atk := args[1].(*coretype.AttackEvent)
		if atk.Info.ActorIndex != char.Index() {
			return false
		}
		if c.ActiveChar != char.Index() {
			return false
		}
		if atk.Info.AttackTag != core.AttackTagElementalArt && atk.Info.AttackTag != core.AttackTagElementalArtHold {
			return false
		}
		if last != 0 && c.Frame-last < cd {
			return false
		}
		if char.Cooldown(core.ActionSkill) == 0 {
			return false
		}
		if c.Rand.Float64() < prob {
			char.ResetActionCooldown(core.ActionSkill)
			last = c.Frame
			c.Log.NewEvent("sacrificial proc'd", coretype.LogWeaponEvent, char.Index())
		}
		return false
	}, fmt.Sprintf("sac-%v", char.Name()))

}

//For every point of the entire party's combined maximum Energy capacity,
//the Elemental Burst DMG of the character equipping this weapon is increased by 0.12%.
//A maximum of 40% increased Elemental Burst DMG can be achieved this way.
//r1 0.12 40%
//r2 0.15 50%
//r3 0.18 60%
//r4 0.21 70%
//r5 0.24 80%
func Wavebreaker(char coretype.Character, c *core.Core, r int, param map[string]int) {

	per := 0.09 + 0.03*float64(r)
	max := 0.3 + 0.1*float64(r)

	var amt float64

	c.Subscribe(core.OnInitialize, func(args ...interface{}) bool {
		var energy float64
		//calculate total team energy
		for _, x := range c.Chars {
			energy += x.MaxEnergy()
		}

		amt = energy * per / 100
		if amt > max {
			amt = max
		}
		c.Log.NewEvent("wavebreaker dmg calc", coretype.LogWeaponEvent, char.Index(), "total", energy, "per", per, "max", max, "amt", amt)
		m := make([]float64, core.EndStatType)
		m[core.DmgP] = amt
		char.AddPreDamageMod(coretype.PreDamageMod{
			Expiry: -1,
			Key:    "wavebreaker",
			Amount: func(atk *coretype.AttackEvent, t coretype.Target) ([]float64, bool) {
				if atk.Info.AttackTag == core.AttackTagElementalBurst {
					return m, true
				}
				return nil, false
			},
		})
		return true
	}, fmt.Sprintf("wavebreaker-%v", char.Name()))

}
