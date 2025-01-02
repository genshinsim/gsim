package chasca

import (
	"errors"
	"fmt"

	"github.com/genshinsim/gcsim/internal/frames"
	"github.com/genshinsim/gcsim/pkg/core/action"
	"github.com/genshinsim/gcsim/pkg/core/attacks"
	"github.com/genshinsim/gcsim/pkg/core/attributes"
	"github.com/genshinsim/gcsim/pkg/core/combat"
	"github.com/genshinsim/gcsim/pkg/core/geometry"
)

var aimedFrames [][]int
var aimedHitmarks = []int{15, 86}

// TODO: confirm these hitmarks
var skillAimHitmarks = []int{2, 4, 6, 8, 10, 12}
var skillAimFrames []int

// per bullet E CA Charge Time = []int{29, 21, 15, 19, 19, 14}
var cumuSkillAimChargeFrames = []int{
	29,
	29 + 21,
	29 + 21 + 15,
	29 + 21 + 15 + 19,
	29 + 21 + 15 + 19 + 19,
	29 + 21 + 15 + 19 + 19 + 14,
}

func init() {
	aimedFrames = make([][]int, 2)

	// Aimed Shot
	aimedFrames[0] = frames.InitAbilSlice(26)
	aimedFrames[0][action.ActionDash] = aimedHitmarks[0]
	aimedFrames[0][action.ActionJump] = aimedHitmarks[0]

	// Fully-Charged Aimed Shot
	aimedFrames[1] = frames.InitAbilSlice(96)
	aimedFrames[1][action.ActionDash] = aimedHitmarks[1]
	aimedFrames[1][action.ActionJump] = aimedHitmarks[1]

	skillAimFrames = frames.InitAbilSlice(31)
	skillAimFrames[action.ActionAttack] = 18
	skillAimFrames[action.ActionCharge] = 18
	skillAimFrames[action.ActionBurst] = 12
	skillAimFrames[action.ActionDash] = 3
	skillAimFrames[action.ActionJump] = 19
}

func (c *char) Aimed(p map[string]int) (action.Info, error) {
	if c.nightsoulState.HasBlessing() {
		return c.aimSkillHold(p)
	}

	hold, ok := p["hold"]
	if !ok {
		hold = attacks.AimParamLv1
	}
	switch hold {
	case attacks.AimParamPhys:
	case attacks.AimParamLv1:
	default:
		return action.Info{}, fmt.Errorf("invalid hold param supplied, got %v", hold)
	}
	travel, ok := p["travel"]
	if !ok {
		travel = 10
	}
	weakspot := p["weakspot"]

	ai := combat.AttackInfo{
		ActorIndex:           c.Index,
		Abil:                 "Fully-Charged Aimed Shot",
		AttackTag:            attacks.AttackTagExtra,
		ICDTag:               attacks.ICDTagNone,
		ICDGroup:             attacks.ICDGroupDefault,
		StrikeType:           attacks.StrikeTypePierce,
		Element:              attributes.Anemo,
		Durability:           25,
		Mult:                 fullAim[c.TalentLvlAttack()],
		HitWeakPoint:         weakspot == 1,
		HitlagHaltFrames:     0.12 * 60,
		HitlagFactor:         0.01,
		HitlagOnHeadshotOnly: true,
		IsDeployable:         true,
	}
	if hold < attacks.AimParamLv1 {
		ai.Abil = "Aimed Shot"
		ai.Element = attributes.Physical
		ai.Mult = aim[c.TalentLvlAttack()]
	}
	c.Core.QueueAttack(
		ai,
		combat.NewBoxHit(
			c.Core.Combat.Player(),
			c.Core.Combat.PrimaryTarget(),
			geometry.Point{Y: -0.5},
			0.1,
			1,
		),
		aimedHitmarks[hold],
		aimedHitmarks[hold]+travel,
	)

	return action.Info{
		Frames:          frames.NewAbilFunc(aimedFrames[hold]),
		AnimationLength: aimedFrames[hold][action.InvalidAction],
		CanQueueAfter:   aimedHitmarks[hold],
		State:           action.AimState,
	}, nil
}

func (c *char) loadSkillHoldBullets() {
	c.bullets[0] = attributes.Anemo
	c.bullets[1] = attributes.Anemo
	c.bullets[2] = c.a1Conversion()

	if c.partyPHECCount < 3 {
		c.bullets[3] = attributes.Anemo
	} else {
		c.bullets[3] = c.randomElemFromBulletPool()
	}

	if c.partyPHECCount < 2 {
		c.bullets[4] = attributes.Anemo
	} else {
		c.bullets[4] = c.randomElemFromBulletPool()
	}

	if c.partyPHECCount < 1 {
		c.bullets[5] = attributes.Anemo
	} else {
		c.bullets[5] = c.randomElemFromBulletPool()
	}
}

func (c *char) resetBulletPool() {
	c.bulletPool = make([]attributes.Element, len(c.partyPHECTypes))
	copy(c.bulletPool, c.partyPHECTypes)
}

func (c *char) randomElemFromBulletPool() attributes.Element {
	if len(c.bulletPool) == 0 {
		c.resetBulletPool()
	}
	ind := c.Core.Rand.Intn(len(c.bulletPool))
	ele := c.bulletPool[ind]
	c.bulletPool[ind] = c.bulletPool[len(c.bulletPool)-1]
	c.bulletPool = c.bulletPool[:len(c.bulletPool)-1]
	return ele
}

func (c *char) aimSkillHold(p map[string]int) (action.Info, error) {
	count, ok := p["count"]
	if !ok {
		count = 6
	}
	if count > 6 {
		return action.Info{}, errors.New("count must be <= 6")
	}

	ai := combat.AttackInfo{
		ActorIndex:     c.Index,
		Abil:           "Shadowhunt Shell",
		AttackTag:      attacks.AttackTagExtra,
		AdditionalTags: []attacks.AdditionalTag{attacks.AdditionalTagNightsoul},
		ICDTag:         attacks.ICDTagChascaShining,
		ICDGroup:       attacks.ICDGroupChascaSkillShining,
		StrikeType:     attacks.StrikeTypeDefault,
		Element:        attributes.Anemo,
		Durability:     25,
		Mult:           skillShadowhunt[c.TalentLvlSkill()],
	}

	chargeDelay := cumuSkillAimChargeFrames[count]
	for i := 0; i < count; i++ {
		bulletElem := c.bullets[count-1-i] // get bullets starting from the back
		switch bulletElem {
		case attributes.Anemo:
			ai.Abil = "Shadowhunt Shell"
			ai.Element = attributes.Anemo
			ai.Mult = skillShadowhunt[c.TalentLvlSkill()]
		default:
			ai.Abil = fmt.Sprintf("Shining Shadowhunt Shell (%s)", bulletElem)
			ai.Element = bulletElem
			ai.Mult = skillShining[c.TalentLvlSkill()]
		}
		hitDelay := chargeDelay + skillAimHitmarks[i]
		c.Core.QueueAttack(ai, combat.NewSingleTargetHit(c.Core.Combat.PrimaryTarget().Key()), hitDelay, hitDelay, c.particleCB)
	}
	c.loadSkillHoldBullets()

	return action.Info{
		Frames: func(next action.Action) int {
			return skillAimFrames[next] + chargeDelay
		},
		AnimationLength: skillAimFrames[action.InvalidAction] + chargeDelay,
		CanQueueAfter:   skillAimFrames[action.ActionDash] + chargeDelay,
		State:           action.AimState,
	}, nil
}
