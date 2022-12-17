package wanderer

import (
	"fmt"

	"github.com/genshinsim/gcsim/internal/frames"
	"github.com/genshinsim/gcsim/pkg/core/action"
	"github.com/genshinsim/gcsim/pkg/core/attributes"
	"github.com/genshinsim/gcsim/pkg/core/combat"
)

var (
	attackFramesNormal   [][]int
	attackFramesE   [][]int
	attackHitmarksNormal = [][]int{{11}, {6}, {32, 41}}
	attackHitmarksE = [][]int{{15}, {3}, {32, 40}}
	attackRadius   = []float64{1.8, 1.8, 2.2}
)

const normalHitNum = 3

func init() {
	//TODO: Release = Hitmark? (No Travel Time)
	attackFramesNormal = make([][]int, normalHitNum)

	attackFramesNormal[0] = frames.InitNormalCancelSlice(attackHitmarksNormal[0][0], 12)
	attackFramesNormal[0][action.ActionAttack] = 25
	attackFramesNormal[0][action.ActionCharge] = 25
	attackFramesNormal[0][action.ActionWalk] = 35

	attackFramesNormal[1] = frames.InitNormalCancelSlice(attackHitmarksNormal[1][0], 5)
	attackFramesNormal[1][action.ActionAttack] = 18
	attackFramesNormal[1][action.ActionCharge] = 27
	attackFramesNormal[1][action.ActionWalk] = 39

	attackFramesNormal[2] = frames.InitNormalCancelSlice(attackHitmarksNormal[2][0], 33)
	attackFramesNormal[2][action.ActionAttack] = 64
	attackFramesNormal[2][action.ActionCharge] = 50
	attackFramesNormal[2][action.ActionWalk] = 39


	attackFramesE = make([][]int, normalHitNum)

	attackFramesE[0] = frames.InitNormalCancelSlice(attackHitmarksE[0][0], 15)
	attackFramesE[0][action.ActionAttack] = 30
	attackFramesE[0][action.ActionCharge] = 31
	attackFramesE[0][action.ActionWalk] = 43

	attackFramesE[1] = frames.InitNormalCancelSlice(attackHitmarksE[1][0], 5)
	attackFramesE[1][action.ActionAttack] = 17
	attackFramesE[1][action.ActionCharge] = 23
	attackFramesE[1][action.ActionWalk] = 34

	attackFramesE[2] = frames.InitNormalCancelSlice(attackHitmarksE[2][0], 32)
	attackFramesE[2][action.ActionAttack] = 54
	attackFramesE[2][action.ActionCharge] = 53
	attackFramesE[2][action.ActionWalk] = 34
}

func (c *char) Attack(p map[string]int) action.ActionInfo {
	delay := c.checkForSkillEnd()

	relevantHitmarks := attackHitmarksNormal
	relevantFrames := attackFramesNormal

	if c.StatusIsActive(skillKey) {
		relevantHitmarks = attackHitmarksE
		relevantFrames = attackFramesE
	}

	for i := 0; i < hits[c.NormalCounter]; i++ {
		ai := combat.AttackInfo{
			ActorIndex: c.Index,
			Abil:       fmt.Sprintf("Normal %v", c.NormalCounter),
			AttackTag:  combat.AttackTagNormal,
			ICDTag:     combat.ICDTagNormalAttack,
			ICDGroup:   combat.ICDGroupDefault,
			StrikeType: combat.StrikeTypeDefault,
			Element:    attributes.Anemo,
			Durability: 25,
			Mult:       attack[c.NormalCounter][i][c.TalentLvlAttack()],
		}
		radius := attackRadius[c.NormalCounter]

		c.Core.QueueAttack(
			ai,
			combat.NewCircleHit(c.Core.Combat.Player(), radius),
			delay,
			delay+relevantHitmarks[c.NormalCounter][i],
		)
	}

	defer c.AdvanceNormalIndex()

	return action.ActionInfo{
		Frames:          func(next action.Action) int { return delay +
			frames.AtkSpdAdjust(relevantFrames[c.NormalCounter][next], c.Stat(attributes.AtkSpd)) },
		AnimationLength: delay + relevantFrames[c.NormalCounter][action.InvalidAction],
		CanQueueAfter:   delay + relevantHitmarks[c.NormalCounter][len(relevantHitmarks[c.NormalCounter])-1],
		State:           action.NormalAttackState,
	}

}
