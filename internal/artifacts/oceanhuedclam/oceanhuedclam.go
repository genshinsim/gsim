package oceanhuedclam

import (
	"fmt"

	"github.com/genshinsim/gcsim/pkg/core"
)

func init() {
	core.RegisterSetFunc("oceanhuedclam", New)
	core.RegisterSetFunc("ocean hued clam", New)
	core.RegisterSetFunc("ocean-hued clam", New)
}

// 2-Piece Bonus: Healing Bonus +15%
// 4-Piece Bonus: When the character equipping this artifact set heals a character in the party, a Sea-Dyed Foam will appear for 3 seconds, accumulating the amount of HP recovered from healing (including overflow healing). At the end of the duration, the Sea-Dyed Foam will explode, dealing DMG to nearby opponents based on 90% of the accumulated healing.
// (This DMG is calculated similarly to Reactions such as Electro-Charged, and Superconduct, but it is not affected by Elemental Mastery, Character Levels, or Reaction DMG Bonuses).
// 	Only one Sea-Dyed Foam can be produced every 3.5 seconds. Each Sea-Dyed Foam can accumulate up to 30,000 HP (including overflow healing). There can be no more than one Sea-Dyed Foam active at any given time.
// 	This effect can still be triggered even when the character who is using this artifact set is not on the field.
func New(c core.Character, s *core.Core, count int) {
	if count >= 2 {
		m := make([]float64, core.EndStatType)
		m[core.Heal] = 0.15
		c.AddMod(core.CharStatMod{
			Key: "ohc-2pc",
			Amount: func(a core.AttackTag) ([]float64, bool) {
				return m, true
			},
			Expiry: -1,
		})
	}
	if count >= 4 {
		bubbleHealStacks := 0.0
		bubbleDurationExpiry := 0
		bubbleICDExpiry := 0

		s.Events.Subscribe(core.OnInitialize, func(args ...interface{}) bool {
			s.Flags.OHCActiveChar = -1
			return true
		}, "OHC-init")

		// On Heal subscription to start accumulating the healing
		s.Events.Subscribe(core.OnHeal, func(args ...interface{}) bool {
			src := args[0].(int)
			healAmt := args[2].(float64)

			if src != c.CharIndex() {
				return false
			}
			// OHC must either be inactive or this equipped character has to have an OHC bubble active
			if !((s.Flags.OHCActiveChar == -1) || (s.Flags.OHCActiveChar == c.CharIndex())) {
				return false
			}

			bubbleHealStacks += healAmt
			if bubbleHealStacks >= 30000 {
				bubbleHealStacks = 30000
			}

			// Activate bubble if this character's bubble is off CD, and add the bubble pop task
			if bubbleICDExpiry < s.F {
				bubbleDurationExpiry = s.F + 3*60
				bubbleICDExpiry = s.F + 3.5*60

				s.Flags.OHCActiveChar = c.CharIndex()

				// Bubble pop task
				c.AddTask(func() {
					// Bubble is physical damage. This is handled in the reaction damage function, so it is not affected by physical dmg bonus/enemy defense
					d := c.Snapshot(
						"OHC Damage",
						core.AttackTagNone,
						core.ICDTagNone,
						core.ICDGroupDefault,
						core.StrikeTypeDefault,
						core.Physical,
						0,
						0,
					)
					d.Targets = core.TargetAll
					d.IsOHCDamage = true
					d.FlatDmg = bubbleHealStacks * .9
					c.QueueDmg(&d, 0)

					// Reset
					s.Flags.OHCActiveChar = -1
					bubbleHealStacks = 0
				}, "ohc-bubble-pop", 3*60)

				s.Log.Debugw("ohc bubble activated", "frame", s.F, "event", core.LogArtifactEvent, "char", c.CharIndex(), "bubble_pops_at", bubbleDurationExpiry, "ohc_icd_expiry", bubbleICDExpiry)
			}

			s.Log.Debugw("ohc bubble accumulation", "frame", s.F, "event", core.LogArtifactEvent, "char", c.CharIndex(), "bubble_pops_at", bubbleDurationExpiry, "bubble_total", bubbleHealStacks)

			return false
		}, fmt.Sprintf("ohc-4pc-heal-accumulation"))
	}
}
