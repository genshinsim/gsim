package fischl

import (
	tmpl "github.com/genshinsim/gcsim/internal/template/character"
	"github.com/genshinsim/gcsim/pkg/core"
	"github.com/genshinsim/gcsim/pkg/core/action"
	"github.com/genshinsim/gcsim/pkg/core/combat"
	"github.com/genshinsim/gcsim/pkg/core/geometry"
	"github.com/genshinsim/gcsim/pkg/core/info"
	"github.com/genshinsim/gcsim/pkg/core/keys"
	"github.com/genshinsim/gcsim/pkg/core/player/character"
	"github.com/genshinsim/gcsim/pkg/model"
)

func init() {
	core.RegisterCharFunc(keys.Fischl, NewChar)
}

type char struct {
	*tmpl.Character
	// field use for calculating oz damage
	ozPos           geometry.Point
	ozSnapshot      combat.AttackEvent
	ozSource        int  // keep tracks of source of oz aka resets
	ozActive        bool // purely used for gscl conditional purposes
	ozActiveUntil   int  // used for oz ticks, a4, c1 and c6
	ozTickSrc       int  // used for oz recast attacks
	ozTravel        int
	burstOzSpawnSrc int // prevent double oz spawn from burst
}

func NewChar(s *core.Core, w *character.CharWrapper, p info.CharacterProfile) error {
	c := char{}
	c.Character = tmpl.NewWithWrapper(s, w)

	c.EnergyMax = 60
	c.NormalHitNum = normalHitNum
	c.SkillCon = 3
	c.BurstCon = 5

	c.ozSource = -1
	c.ozActive = false
	c.ozActiveUntil = -1
	c.ozTickSrc = -1

	c.ozTravel = 10
	travel, ok := p.Params["oz_travel"]
	if ok {
		c.ozTravel = travel
	}

	w.Character = &c

	return nil
}

func (c *char) Init() error {
	c.a4()
	if c.Base.Cons >= 6 {
		c.c6()
	}
	return nil
}

func (c *char) Condition(fields []string) (any, error) {
	switch fields[0] {
	case "oz":
		return c.ozActive, nil
	case "oz-source":
		return c.ozSource, nil
	case "oz-duration":
		duration := c.ozActiveUntil - c.Core.F
		if duration < 0 {
			duration = 0
		}
		return duration, nil
	default:
		return c.Character.Condition(fields)
	}
}

func (c *char) ActionReady(a action.Action, p map[string]int) (bool, action.Failure) {
	// check if it is possible to recast oz
	if a == action.ActionSkill && p["recast"] != 0 && c.ozActive {
		return !c.StatusIsActive(skillRecastCDKey), action.SkillCD
	}
	// check if cast skill with oz on-field
	if a == action.ActionSkill && c.ozActive {
		return false, action.NoFailure
	}
	return c.Character.ActionReady(a, p)
}

func (c *char) AnimationStartDelay(k model.AnimationDelayKey) int {
	if k == model.AnimationXingqiuN0StartDelay {
		return 9
	}
	return c.AnimationStartDelay(k)
}
