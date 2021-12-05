package qiqi

import (
	"github.com/genshinsim/gcsim/pkg/character"
	"github.com/genshinsim/gcsim/pkg/core"
)

func init() {
	core.RegisterCharFunc("qiqi", NewChar)
}

type char struct {
	*character.Tmpl
	talismanExpiry    []int
	talismanICDExpiry []int
	c4ICDExpiry       int
	skillLastUsed     int
	skillHealSnapshot core.Snapshot // Required as both on hit procs and continuous healing need to use this
}

// TODO: Not implemented - C6 (revival mechanic, not suitable for sim)
// C4 - Enemy Atk reduction, not useful in this sim version
func NewChar(s *core.Core, p core.CharacterProfile) (core.Character, error) {
	c := char{}
	t, err := character.NewTemplateChar(s, p)
	if err != nil {
		return nil, err
	}
	c.Tmpl = t
	c.Energy = 80
	c.EnergyMax = 80
	c.Weapon.Class = core.WeaponClassSword
	c.NormalHitNum = 5
	c.BurstCon = 3
	c.SkillCon = 5

	c.skillLastUsed = 0

	c.talismanHealHook()
	c.onNACAHitHook()
	c.a1()

	return &c, nil
}

// Ensures the set of targets are initialized properly
func (c *char) Init(index int) {
	c.Tmpl.Init(index)

	c.talismanExpiry = make([]int, len(c.Core.Targets))
	c.talismanICDExpiry = make([]int, len(c.Core.Targets))
}

func (c *char) talismanHealHook() {
	c.Core.Events.Subscribe(core.OnDamage, func(args ...interface{}) bool {
		t := args[0].(core.Target)
		ds := args[1].(*core.Snapshot)

		if c.talismanExpiry[t.Index()] < c.Core.F {
			return false
		}
		if c.talismanICDExpiry[t.Index()] >= c.Core.F {
			return false
		}

		healAmt := c.healDynamic(burstHealPer, burstHealFlat, c.TalentLvlBurst())
		c.Core.Health.HealIndex(c.Index, ds.ActorIndex, healAmt)
		c.talismanICDExpiry[t.Index()] = c.Core.F + 60

		c.Core.Log.Debugw("Qiqi Talisman Healing", "frame", c.Core.F, "event", core.LogCharacterEvent, "char", c.Index, "target", t.Index(), "healed_char", ds.ActorIndex, "talisman_expiry", c.talismanExpiry[t.Index()], "talisman_healing_icd", c.talismanICDExpiry[t.Index()], "healed_amt", healAmt)

		return false
	}, "talisman-heal-hook")
}

// Handles C2, A4, and skill NA/CA on hit hooks
// Additionally handles burst Talisman hook - can't be done another way since Talisman is applied before the burst damage is dealt
func (c *char) onNACAHitHook() {
	c.Core.Events.Subscribe(core.OnAttackWillLand, func(args ...interface{}) bool {
		t := args[0].(core.Target)
		ds := args[1].(*core.Snapshot)

		if ds.ActorIndex != c.Index {
			return false
		}

		// Talisman is applied before the damage is dealt
		if ds.Abil == "Fortune-Preserving Talisman" {
			c.talismanExpiry[t.Index()] = c.Core.F + 15*60
		}

		// All of the below only occur on Qiqi NA/CA hits
		if !((ds.AttackTag == core.AttackTagNormal) || (ds.AttackTag == core.AttackTagExtra)) {
			return false
		}

		// C2
		// Qiqi’s Normal and Charge Attack DMG against opponents affected by Cryo is increased by 15%.
		if (c.Base.Cons >= 2) && (t.AuraContains(core.Cryo)) {
			ds.Stats[core.DmgP] += .15
		}

		// A4
		// When Qiqi hits opponents with her Normal and Charged Attacks, she has a 50% chance to apply a Fortune-Preserving Talisman to them for 6s. This effect can only occur once every 30s.
		if (c.c4ICDExpiry <= c.Core.F) && (c.Rand.Float64() < 0.5) {
			// Don't want to overwrite a longer burst duration talisman with a shorter duration one
			// TODO: Unclear how the interaction works if there is already a talisman on enemy
			// TODO: Being generous for now and not putting it on CD if there is a conflict
			if c.talismanExpiry[t.Index()] < c.Core.F+360 {
				c.talismanExpiry[t.Index()] = c.Core.F + 360
				c.c4ICDExpiry = c.Core.F + 30*60

				c.Core.Log.Debugw("Qiqi A4 Adding Talisman", "frame", c.Core.F, "event", core.LogCharacterEvent, "char", c.Index, "target", t.Index(), "talisman_expiry", c.talismanExpiry[t.Index()], "c4_icd_expiry", c.c4ICDExpiry)
			}
		}

		// Qiqi NA/CA healing proc in skill duration
		if c.Core.Status.Duration("qiqiskill") > 0 {
			c.Core.Health.HealAll(c.Index, c.healSnapshot(&c.skillHealSnapshot, skillHealOnHitPer, skillHealOnHitFlat, c.TalentLvlSkill()))
		}

		return false
	}, "qiqi-onhit-naca-hook")
}

// Implements event hook and incoming healing bonus function for A1
// TODO: Could possibly change this so the AddIncHealBonus occurs at start, then event subscription occurs upon using Qiqi skill?
// TODO: Likely more efficient to not maintain event subscription always, but grouping the two for clarity currently
// When a character under the effects of Adeptus Art: Herald of Frost triggers an Elemental Reaction, their Incoming Healing Bonus is increased by 20% for 8s.
func (c *char) a1() {

	// Add to incoming healing bonus array
	c.Core.Health.AddIncHealBonus(func(healedCharIndex int) float64 {
		healedCharName := c.Core.Chars[healedCharIndex].Name()

		if c.Core.Status.Duration("qiqia1"+healedCharName) > 0 {
			return .2
		}
		return 0
	})

	a1Hook := func(args ...interface{}) bool {
		if c.Core.Status.Duration("qiqiskill") == 0 {
			return false
		}
		ds := args[1].(*core.Snapshot)

		// Active char is the only one under the effects of Qiqi skill
		if ds.ActorIndex != c.Core.ActiveChar {
			return false
		}

		c.Core.Status.AddStatus("qiqia1"+c.Core.Chars[c.Core.ActiveChar].Name(), 8*60)

		return false
	}
	c.Core.Events.Subscribe(core.OnTransReaction, a1Hook, "qiqi-a1")
	c.Core.Events.Subscribe(core.OnAmpReaction, a1Hook, "qiqi-a1")
}

func (c *char) ActionStam(a core.ActionType, p map[string]int) float64 {
	switch a {
	case core.ActionDash:
		return 18
	case core.ActionCharge:
		return 20
	default:
		c.Core.Log.Warnw("ActionStam not implemented", "character", c.Base.Name)
		return 0
	}
}
