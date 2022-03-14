package amber

import (
	"github.com/genshinsim/gcsim/pkg/core"
	"github.com/genshinsim/gcsim/pkg/coretype"
)

type bunny struct {
	ae  coretype.AttackEvent
	src int
}

//TODO: forbidden bunny cryo swirl tech
func (c *char) makeBunny() {
	b := bunny{}
	b.src = c.Core.Frame
	ai := core.AttackInfo{
		Abil:       "Baron Bunny",
		ActorIndex: c.Index,
		AttackTag:  core.AttackTagElementalArt,
		ICDTag:     core.ICDTagNone,
		ICDGroup:   core.ICDGroupDefault,
		StrikeType: core.StrikeTypeBlunt,
		Element:    core.Pyro,
		Durability: 50,
		Mult:       bunnyExplode[c.TalentLvlSkill()],
	}
	snap := c.Snapshot(&ai)
	b.ae = coretype.AttackEvent{
		Info:        ai,
		Pattern:     core.NewDefCircHit(2, false, coretype.TargettableEnemy),
		SourceFrame: c.Core.Frame,
		Snapshot:    snap,
	}

	c.bunnies = append(c.bunnies, b)

	//ondeath explodes
	//duration is 8.2 sec
	c.AddTask(func() {
		c.explode(b.src)
	}, "bunny", 492)
}

func (c *char) explode(src int) {
	n := 0
	c.coretype.Log.NewEvent("amber exploding bunny", coretype.LogCharacterEvent, c.Index, "src", src)
	for _, v := range c.bunnies {
		if v.src == src {
			c.Core.Combat.QueueAttackEvent(&v.ae, 1)
			//4 orbs
			c.QueueParticle("amber", 4, core.Pyro, 100)
		} else {
			c.bunnies[n] = v
			n++
		}
	}

	c.bunnies = c.bunnies[:n]
}

func (c *char) manualExplode() {
	//do nothing if there are no bunnies
	if len(c.bunnies) == 0 {
		return
	}
	//only explode the first bunny
	if len(c.bunnies) > 0 {
		c.bunnies[0].ae.Info.Mult += 2
		c.Core.Combat.QueueAttackEvent(&c.bunnies[0].ae, 1)
		c.QueueParticle("amber", 4, core.Pyro, 100)
	}
	c.bunnies = c.bunnies[1:]
}

func (c *char) overloadExplode() {
	//explode all bunnies on overload

	c.Core.Subscribe(coretype.OnDamage, func(args ...interface{}) bool {

		atk := args[1].(*coretype.AttackEvent)
		if len(c.bunnies) == 0 {
			return false
		}
		//TODO: only amber trigger?
		if atk.Info.ActorIndex != c.Index {
			return false
		}

		if atk.Info.AttackTag != core.AttackTagOverloadDamage {
			return false
		}

		for _, v := range c.bunnies {
			//every bunny gets bonus multiplikers
			v.ae.Info.Mult += 2
			c.Core.Combat.QueueAttackEvent(&v.ae, 1)
			c.QueueParticle("amber", 4, core.Pyro, 100)
		}
		c.bunnies = make([]bunny, 0, 2)

		return false
	}, "amber-bunny-explode-overload")

}
