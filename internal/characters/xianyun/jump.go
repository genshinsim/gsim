package xianyun

import (
	"github.com/genshinsim/gcsim/internal/frames"
	"github.com/genshinsim/gcsim/pkg/core/action"
	"github.com/genshinsim/gcsim/pkg/core/player"
)

var burstJumpFrames []int

func init() {
	burstJumpFrames = frames.InitAbilSlice(58)
	burstJumpFrames[action.ActionHighPlunge] = 6
	burstJumpFrames[action.ActionLowPlunge] = 5
}

func (c *char) Jump(p map[string]int) (action.Info, error) {
	if c.StatusIsActive(player.XianyunAirborneBuff) {
		c.Core.Player.SetAirborne(player.AirborneXianyun)
		return action.Info{
			Frames:          frames.NewAbilFunc(burstJumpFrames),
			AnimationLength: burstJumpFrames[action.InvalidAction],
			CanQueueAfter:   burstJumpFrames[action.ActionLowPlunge], // earliest cancel
			State:           action.JumpState,
		}, nil
	}
	return c.Character.Jump(p)
}
