package zhongli

import (
	"github.com/genshinsim/gcsim/internal/tmpl/character"
	"github.com/genshinsim/gcsim/pkg/core"
)

type char struct {
	*character.Tmpl
	maxStele   int
	steleCount int
	energyICD  int
}

func init() {
	core.RegisterCharFunc(core.Zhongli, NewChar)
}

func NewChar(s *core.Core, p core.CharacterProfile) (core.Character, error) {
	c := char{}
	t, err := character.NewTemplateChar(s, p)
	if err != nil {
		return nil, err
	}
	c.Tmpl = t
	c.Base.Element = core.Geo
	c.Energy = 40
	c.EnergyMax = 40
	c.Weapon.Class = core.WeaponClassSpear
	c.BurstCon = 5
	c.SkillCon = 3
	c.NormalHitNum = 6

	c.maxStele = 1
	if c.Base.Cons >= 1 {
		c.maxStele = 2
	}

	c.a2()

	return &c, nil
}

func (c *char) ActionStam(a core.ActionType, p map[string]int) float64 {
	switch a {
	case core.ActionDash:
		return 18
	case core.ActionCharge:
		return 25
	default:
		c.Core.Log.NewEvent("ActionStam not implemented", core.LogActionEvent, c.Index, "action", a.String())
		return 0
	}

}

func (c *char) a2() {
	c.Core.Shields.AddBonus(func() float64 {
		if c.Tags["shielded"] == 0 {
			return 0
		}
		return float64(c.Tags["a2"]) * 0.05
	})
}
