package keqing

import (
	"fmt"

	"github.com/genshinsim/gcsim/pkg/core"
	"github.com/genshinsim/gcsim/pkg/coretype"
)

func (c *char) Attack(p map[string]int) (int, int) {
	//apply attack speed
	f, a := c.ActionFrames(core.ActionAttack, p)

	ai := core.AttackInfo{
		ActorIndex: c.Index,
		Abil:       fmt.Sprintf("Normal %v", c.NormalCounter),
		AttackTag:  coretype.AttackTagNormal,
		ICDTag:     core.ICDTagNormalAttack,
		ICDGroup:   core.ICDGroupDefault,
		Element:    core.Physical,
		Durability: 25,
	}

	for i, mult := range attack[c.NormalCounter] {
		ai.Mult = mult[c.TalentLvlAttack()]
		c.Core.Combat.QueueAttack(ai, core.NewDefCircHit(0.1, false, coretype.TargettableEnemy), delay[c.NormalCounter][i], delay[c.NormalCounter][i])
	}

	if c.Base.Cons == 6 {
		c.activateC6("attack")
	}

	c.AdvanceNormalIndex()
	return f, a
}

func (c *char) ChargeAttack(p map[string]int) (int, int) {

	f, a := c.ActionFrames(core.ActionCharge, p)

	ai := core.AttackInfo{
		ActorIndex: c.Index,
		AttackTag:  coretype.AttackTagExtra,
		ICDTag:     core.ICDTagNormalAttack,
		ICDGroup:   core.ICDGroupDefault,
		Element:    core.Physical,
		Durability: 25,
	}

	for i, mult := range charge {
		ai.Mult = mult[c.TalentLvlAttack()]
		ai.Abil = fmt.Sprintf("Charge %v", i)
		c.Core.Combat.QueueAttack(ai, core.NewDefCircHit(1, false, coretype.TargettableEnemy), f-i*10-5, f-i*10-5)
	}

	if c.Tags["e"] == 1 {
		//2 hits
		ai := core.AttackInfo{
			ActorIndex: c.Index,
			AttackTag:  core.AttackTagElementalArt,
			ICDTag:     core.ICDTagElementalArt,
			ICDGroup:   core.ICDGroupDefault,
			Element:    core.Electro,
			Durability: 50,
			Mult:       skillCA[c.TalentLvlSkill()],
		}
		for i := 0; i < 2; i++ {
			c.Core.Combat.QueueAttack(ai, core.NewDefCircHit(1, false, coretype.TargettableEnemy), f, f)
		}

		//place on cooldown
		c.Tags["e"] = 0
		// c.CD[def.SkillCD] = c.eStartFrame + 100
		c.SetCD(core.ActionSkill, c.eStartFrame+450-c.Core.Frame)

		// TODO: Particle timing?
		if c.Core.Rand.Float64() < .5 {
			c.QueueParticle("keqing", 2, core.Electro, 100)
		} else {
			c.QueueParticle("keqing", 3, core.Electro, 100)
		}
	}

	if c.Base.Cons == 6 {
		c.activateC6("charge")
	}

	return f, a
}

func (c *char) Skill(p map[string]int) (int, int) {
	if c.Tags["e"] == 1 {
		return c.skillNext(p)
	}
	return c.skillFirst(p)
}

func (c *char) skillFirst(p map[string]int) (int, int) {

	f, a := c.ActionFrames(core.ActionSkill, p)

	ai := core.AttackInfo{
		Abil:       "Stellar Restoration",
		ActorIndex: c.Index,
		AttackTag:  core.AttackTagElementalArt,
		ICDTag:     core.ICDTagNone,
		ICDGroup:   core.ICDGroupDefault,
		Element:    core.Electro,
		Durability: 25,
		Mult:       skill[c.TalentLvlSkill()],
	}

	c.Core.Combat.QueueAttack(ai, core.NewDefCircHit(1, false, coretype.TargettableEnemy), f, f)

	c.Tags["e"] = 1
	c.eStartFrame = c.Core.Frame

	//place on cd after certain frames if started is still true
	//looks like the E thing lasts 5 seconds
	c.AddTask(func() {
		if c.Tags["e"] == 1 {
			c.Tags["e"] = 0
			// c.CD[def.SkillCD] = c.eStartFrame + 100
			c.SetCD(core.ActionSkill, c.eStartFrame+450-c.Core.Frame) //TODO: cooldown if not triggered, 7.5s
		}
	}, "keqing-skill-cd", c.Core.Frame+300) //TODO: check this

	if c.Base.Cons == 6 {
		c.activateC6("skill")
	}

	return f, a
}

func (c *char) skillNext(p map[string]int) (int, int) {
	f, a := c.ActionFrames(core.ActionSkill, p)

	ai := core.AttackInfo{
		Abil:       "Stellar Restoration (Slashing)",
		ActorIndex: c.Index,
		AttackTag:  core.AttackTagElementalArt,
		ICDTag:     core.ICDTagElementalArt,
		ICDGroup:   core.ICDGroupDefault,
		Element:    core.Electro,
		Durability: 50,
		Mult:       skillPress[c.TalentLvlSkill()],
	}

	c.Core.Combat.QueueAttack(ai, core.NewDefCircHit(1, false, coretype.TargettableEnemy), f, f)
	//add electro infusion

	c.Core.AddStatus("keqinginfuse", 300)

	c.AddWeaponInfuse(core.WeaponInfusion{
		Key:    "a2",
		Ele:    core.Electro,
		Tags:   []core.AttackTag{coretype.AttackTagNormal, coretype.AttackTagExtra, core.AttackTagPlunge},
		Expiry: c.Core.Frame + 300,
	})

	if c.Base.Cons >= 1 {
		//2 tick dmg at start to end
		hits, ok := p["c1"]
		if !ok {
			hits = 1 //default 1 hit
		}
		ai := core.AttackInfo{
			Abil:       "Stellar Restoration (Slashing)",
			ActorIndex: c.Index,
			AttackTag:  core.AttackTagElementalArtHold,
			ICDTag:     core.ICDTagElementalArt,
			ICDGroup:   core.ICDGroupDefault,
			Element:    core.Electro,
			Durability: 25,
			Mult:       .5,
		}
		for i := 0; i < hits; i++ {
			c.Core.Combat.QueueAttack(ai, core.NewDefCircHit(2, false, coretype.TargettableEnemy), 0, f)
		}
	}

	// TODO: Particle timing?
	if c.Core.Rand.Float64() < .5 {
		c.QueueParticle("keqing", 2, core.Electro, 100)
	} else {
		c.QueueParticle("keqing", 3, core.Electro, 100)
	}

	//place on cooldown
	c.Tags["e"] = 0
	c.SetCD(core.ActionSkill, c.eStartFrame+450-c.Core.Frame)
	return f, a
}

func (c *char) Burst(p map[string]int) (int, int) {
	f, a := c.ActionFrames(core.ActionBurst, p)
	//a4 increase crit + ER
	val := make([]float64, core.EndStatType)
	val[core.CR] = 0.15
	val[core.ER] = 0.15
	c.AddMod(coretype.CharStatMod{
		Key:    "a4",
		Amount: func() ([]float64, bool) { return val, true },
		Expiry: c.Core.Frame + 480,
	})

	//first hit 70 frame
	//first tick 74 frame
	//last tick 168
	//last hit 211

	//initial
	ai := core.AttackInfo{
		Abil:       "Starward Sword (Cast)",
		ActorIndex: c.Index,
		AttackTag:  core.AttackTagElementalBurst,
		ICDTag:     core.ICDTagElementalBurst,
		ICDGroup:   core.ICDGroupDefault,
		Element:    core.Electro,
		Durability: 25,
		Mult:       burstInitial[c.TalentLvlBurst()],
	}

	c.Core.Combat.QueueAttack(ai, core.NewDefCircHit(5, false, coretype.TargettableEnemy), 70, 70)
	//8 hits

	ai.Abil = "Starward Sword (Consecutive Slash)"
	ai.Mult = burstDot[c.TalentLvlBurst()]
	for i := 70; i < 170; i += 13 {
		c.Core.Combat.QueueAttack(ai, core.NewDefCircHit(5, false, coretype.TargettableEnemy), i, i)
	}

	//final

	ai.Abil = "Starward Sword (Last Attack)"
	ai.Mult = burstFinal[c.TalentLvlBurst()]
	c.Core.Combat.QueueAttack(ai, core.NewDefCircHit(5, false, coretype.TargettableEnemy), 211, 211)

	if c.Base.Cons == 6 {
		c.activateC6("burst")
	}

	c.ConsumeEnergy(60)
	// c.CD[def.BurstCD] = c.Core.F + 720 //12s
	c.SetCDWithDelay(core.ActionBurst, 720, 60)

	return f, a
}
