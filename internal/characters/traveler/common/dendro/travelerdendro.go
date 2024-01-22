package dendro

import (
	"github.com/genshinsim/gcsim/internal/characters/traveler/common"
	tmpl "github.com/genshinsim/gcsim/internal/template/character"
	"github.com/genshinsim/gcsim/pkg/core"
	"github.com/genshinsim/gcsim/pkg/core/attributes"
	"github.com/genshinsim/gcsim/pkg/core/geometry"
	"github.com/genshinsim/gcsim/pkg/core/info"
	"github.com/genshinsim/gcsim/pkg/core/player/character"
)

type char struct {
	*tmpl.Character
	burstPos                   geometry.Point
	burstRadius                float64
	burstOverflowingLotuslight int
	skillC1                    bool // this variable also ensures that C1 only restores energy once per cast
	burstTransfig              attributes.Element
	gender                     int
}

func NewChar(gender int) core.NewCharacterFunc {
	return func(s *core.Core, w *character.CharWrapper, p info.CharacterProfile) error {
		c := char{
			gender: gender,
		}
		c.Character = tmpl.NewWithWrapper(s, w)

		c.Base.Atk += common.TravelerBaseAtkIncrease(p)
		c.Base.Element = attributes.Dendro
		c.EnergyMax = 80
		c.BurstCon = 5
		c.SkillCon = 3
		c.NormalHitNum = normalHitNum

		w.Character = &c

		return nil
	}
}

func (c *char) Init() error {
	c.a1Init()
	c.a4Init()

	if c.Base.Cons >= 6 {
		c.c6Init()
	}

	c.skillC1 = false
	return nil
}
