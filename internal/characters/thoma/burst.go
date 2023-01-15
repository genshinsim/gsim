package thoma

import (
	"github.com/genshinsim/gcsim/internal/frames"
	"github.com/genshinsim/gcsim/pkg/core/action"
	"github.com/genshinsim/gcsim/pkg/core/attributes"
	"github.com/genshinsim/gcsim/pkg/core/combat"
)

var burstFrames []int

const (
	burstKey     = "thoma-q"
	burstICDKey  = "thoma-q-icd"
	burstHitmark = 40
)

func init() {
	burstFrames = frames.InitAbilSlice(58)
	burstFrames[action.ActionAttack] = 57
	burstFrames[action.ActionSkill] = 56
	burstFrames[action.ActionDash] = 57
	burstFrames[action.ActionSwap] = 56
}

// Burst attack damage queue generator
func (c *char) Burst(p map[string]int) action.ActionInfo {
	ai := combat.AttackInfo{
		ActorIndex: c.Index,
		Abil:       "Crimson Ooyoroi",
		AttackTag:  combat.AttackTagElementalBurst,
		ICDTag:     combat.ICDTagNone,
		ICDGroup:   combat.ICDGroupDefault,
		StrikeType: combat.StrikeTypeDefault,
		Element:    attributes.Pyro,
		Durability: 50,
		Mult:       burst[c.TalentLvlBurst()],
	}

	// damage component not final
	c.Core.QueueAttack(
		ai,
		combat.NewCircleHit(c.Core.Combat.Player(), 4),
		burstHitmark,
		burstHitmark,
	)

	d := 15
	if c.Base.Cons >= 2 {
		d = 18
	}

	c.AddStatus(burstKey, d*60, true)

	// C4: restore 15 energy
	if c.Base.Cons >= 4 {
		c.Core.Tasks.Add(func() {
			c.AddEnergy("thoma-c4", 15)
		}, 8)
	}

	cd := 20
	if c.Base.Cons >= 1 {
		cd = 17 // the CD reduction activates when a character protected by Thoma's shield is hit. Since it is almost impossible for this not to activate, we set the duration to 17 for sim purposes.
	}
	c.SetCD(action.ActionBurst, cd*60)
	c.ConsumeEnergy(7)

	return action.ActionInfo{
		Frames:          frames.NewAbilFunc(burstFrames),
		AnimationLength: burstFrames[action.InvalidAction],
		CanQueueAfter:   burstFrames[action.ActionSkill],
		State:           action.BurstState,
	}
}

func (c *char) summonFieryCollapse() {
	ai := combat.AttackInfo{
		ActorIndex: c.Index,
		Abil:       "Fiery Collapse",
		AttackTag:  combat.AttackTagElementalBurst,
		ICDTag:     combat.ICDTagElementalBurst,
		ICDGroup:   combat.ICDGroupDefault,
		StrikeType: combat.StrikeTypeDefault,
		Element:    attributes.Pyro,
		Durability: 25,
		Mult:       burstproc[c.TalentLvlBurst()],
		FlatDmg:    0.022 * c.MaxHP(),
	}
	done := false
	shieldCb := func(_ combat.AttackCB) {
		if done {
			return
		}
		shieldamt := (burstshieldpp[c.TalentLvlBurst()]*c.MaxHP() + burstshieldflat[c.TalentLvlBurst()])
		c.genShield("Thoma Burst", shieldamt, true)
		done = true
	}
	c.Core.QueueAttack(ai, combat.NewCircleHit(c.Core.Combat.Player(), 4.59), 0, 11, shieldCb)
}
