package kaeya

import "github.com/genshinsim/gcsim/pkg/core"

func (c *char) ActionFrames(a core.ActionType, p map[string]int) (int, int) {
	switch a {
	case core.ActionAttack:
		f := 0
		switch c.NormalCounter {
		//TODO: need to add atkspd mod
		case 0:
			f = 14 //frames from keqing lib
		case 1:
			f = 41 - 14
		case 2:
			f = 72 - 41
		case 3:
			f = 128 - 72
		case 4:
			f = 176 - 128
		}
		f = int(float64(f) / (1 + c.Stat(core.AtkSpd)))
		return f, f
	case core.ActionCharge:
		return 87, 87
	case core.ActionSkill:
		return 58, 58 //could be 52 if going into Q
	case core.ActionBurst:
		return 78, 78
	default:
		c.coretype.Log.NewEventBuildMsg(core.LogActionEvent, c.Index, "unknown action (invalid frames): ", a.String())
		return 0, 0
	}
}
