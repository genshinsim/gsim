package diona

import (
	"fmt"

	"github.com/genshinsim/gcsim/internal/tmpl/shield"
	"github.com/genshinsim/gcsim/pkg/core"
	"github.com/genshinsim/gcsim/pkg/coretype"
)

func (c *char) Attack(p map[string]int) (int, int) {
	travel, ok := p["travel"]
	if !ok {
		travel = 10
	}

	f, a := c.ActionFrames(core.ActionAttack, p)
	ai := core.AttackInfo{
		ActorIndex: c.Index,
		Abil:       fmt.Sprintf("Normal %v", c.NormalCounter),
		AttackTag:  coretype.AttackTagNormal,
		ICDTag:     core.ICDTagNone,
		ICDGroup:   core.ICDGroupDefault,
		StrikeType: core.StrikeTypePierce,
		Element:    core.Physical,
		Durability: 25,
		Mult:       auto[c.NormalCounter][c.TalentLvlAttack()],
	}
	c.Core.Combat.QueueAttack(ai, core.NewDefSingleTarget(1, coretype.TargettableEnemy), f, travel+f)

	c.AdvanceNormalIndex()

	return f, a
}

func (c *char) Aimed(p map[string]int) (int, int) {
	travel, ok := p["travel"]
	if !ok {
		travel = 10
	}
	weakspot, ok := p["weakspot"]

	f, a := c.ActionFrames(core.ActionAim, p)
	ai := core.AttackInfo{
		ActorIndex:   c.Index,
		Abil:         "Aim (Charged)",
		AttackTag:    coretype.AttackTagExtra,
		ICDTag:       core.ICDTagExtraAttack,
		ICDGroup:     core.ICDGroupDefault,
		StrikeType:   core.StrikeTypePierce,
		Element:      coretype.Cryo,
		Durability:   25,
		Mult:         aim[c.TalentLvlAttack()],
		HitWeakPoint: weakspot == 1,
	}
	// d.AnimationFrames = f

	c.Core.Combat.QueueAttack(ai, core.NewDefSingleTarget(1, coretype.TargettableEnemy), f, travel+f)

	return f, a
}

func (c *char) Skill(p map[string]int) (int, int) {
	travel, ok := p["travel"]
	if !ok {
		travel = 10
	}
	f, a := c.ActionFrames(core.ActionSkill, p)

	// 2 paws
	var bonus float64 = 1
	cd := 360 + f
	pawCount := 2

	if p["hold"] == 1 {
		//5 paws, 75% absorption bonus
		bonus = 1.75
		cd = 900 + f
		pawCount = 5
	}

	shd := (pawShieldPer[c.TalentLvlSkill()]*c.MaxHP() + pawShieldFlat[c.TalentLvlSkill()]) * bonus
	if c.Base.Cons >= 2 {
		shd = shd * 1.15
	}
	ai := core.AttackInfo{
		ActorIndex: c.Index,
		Abil:       "Icy Paw",
		AttackTag:  core.AttackTagElementalArt,
		ICDTag:     core.ICDTagElementalArt,
		ICDGroup:   core.ICDGroupDefault,
		StrikeType: core.StrikeTypeDefault,
		Element:    coretype.Cryo,
		Durability: 25,
		Mult:       paw[c.TalentLvlSkill()],
	}
	count := 0

	for i := 0; i < pawCount; i++ {
		c.Core.Combat.QueueAttack(ai, core.NewDefSingleTarget(1, coretype.TargettableEnemy), 0, travel+f-5+i)
		if c.Core.Rand.Float64() < 0.8 {
			count++
		}
	}

	//particles
	c.QueueParticle("Diona", count, coretype.Cryo, 90) //90s travel time

	//add shield
	c.AddTask(func() {
		c.Core.Shields.Add(&shield.Tmpl{
			Src:        c.Core.Frame,
			ShieldType: core.ShieldDionaSkill,
			Name:       "Diona Skill",
			HP:         shd,
			Ele:        coretype.Cryo,
			Expires:    c.Core.Frame + pawDur[c.TalentLvlSkill()], //15 sec
		})
	}, "Diona-Paw-Shield", f)

	c.SetCD(core.ActionSkill, cd)
	return f, a
}

func (c *char) Burst(p map[string]int) (int, int) {

	f, a := c.ActionFrames(core.ActionBurst, p)

	//initial hit
	ai := core.AttackInfo{
		ActorIndex: c.Index,
		Abil:       "Signature Mix (Initial)",
		AttackTag:  core.AttackTagElementalBurst,
		ICDTag:     core.ICDTagElementalBurst,
		ICDGroup:   core.ICDGroupDefault,
		StrikeType: core.StrikeTypeDefault,
		Element:    coretype.Cryo,
		Durability: 25,
		Mult:       burst[c.TalentLvlBurst()],
	}
	c.Core.Combat.QueueAttack(ai, core.NewDefCircHit(5, false, coretype.TargettableEnemy), 0, f-10)

	ai.Abil = "Signature Mix (Tick)"
	ai.Mult = burstDot[c.TalentLvlBurst()]
	snap := c.Snapshot(&ai)
	hpplus := snap.Stats[core.Heal]
	maxhp := c.MaxHP()
	heal := burstHealPer[c.TalentLvlBurst()]*maxhp + burstHealFlat[c.TalentLvlBurst()]

	//ticks every 2s, first tick at t=1s, then t=3,5,7,9,11, lasts for 12.5
	for i := 0; i < 6; i++ {
		c.AddTask(func() {
			c.Core.Combat.QueueAttackWithSnap(ai, snap, core.NewDefCircHit(5, false, core.TargettableEnemy), 0)
			// c.Core.Log.NewEvent("diona healing", core.LogCharacterEvent, c.Index, "+heal", hpplus, "max hp", maxhp, "heal amount", heal)
			c.Core.Health.Heal(core.HealInfo{
				Caller:  c.Index,
				Target:  c.Core.ActiveChar,
				Message: "Drunken Mist",
				Src:     heal,
				Bonus:   hpplus,
			})
		}, "Diona Burst (DOT)", 60+i*120)
	}

	//apparently lasts for 12.5
	c.Core.AddStatus("dionaburst", f+750) //TODO not sure when field starts, is it at animation end? prob when it lands...

	//c1
	if c.Base.Cons >= 1 {
		//15 energy after ends, flat not affected by ER
		c.AddTask(func() {
			c.Energy += 15
			if c.Energy > c.EnergyMax {
				c.Energy = c.EnergyMax
			}
			c.coretype.Log.NewEvent("diona c1 regen 15 energy", coretype.LogEnergyEvent, c.Index, "new energy", c.Energy)
		}, "Diona C1", f+750)
	}

	if c.Base.Cons == 6 {
		c.AddTask(func() {
			for _, char := range c.Core.Chars {
				this := char
				val := make([]float64, core.EndStatType)
				val[core.EM] = 200
				this.AddMod(coretype.CharStatMod{
					Key:    "diona-c6",
					Expiry: c.Core.Frame + 750,
					Amount: func() ([]float64, bool) {
						return val, this.HP()/this.MaxHP() > 0.5
					},
				})
			}
		}, "c6-em-share", f)
	}

	c.SetCDWithDelay(core.ActionBurst, 1200, 49)
	c.ConsumeEnergy(49)
	return f, a
}
