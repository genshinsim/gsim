package target

import (
	"github.com/genshinsim/gcsim/pkg/core"
	"github.com/genshinsim/gcsim/pkg/reactable"
)

type Tmpl struct {
	Core *core.Core
	*reactable.Reactable
	TargetType  core.TargettableType
	TargetIndex int
	HPCurrent   float64
	HPMax       float64
	Hitbox      core.Circle
	Res         map[core.EleType]float64
	Level       int

	//mods
	ResMod []core.ResistMod
	DefMod []core.DefMod

	//icd related
	icdGroupOnTimer       [core.MaxTeamPlayerCount][core.ICDGroupLength]bool
	icdTagCounter         [core.MaxTeamPlayerCount][core.ICDTagLength]int
	icdDamageGroupOnTimer [core.MaxTeamPlayerCount][core.ICDGroupLength]bool
	icdDamageGroupCounter [core.MaxTeamPlayerCount][core.ICDGroupLength]int
}

func (t *Tmpl) Type() core.TargettableType { return t.TargetType }
func (t *Tmpl) Index() int                 { return t.TargetIndex }
func (t *Tmpl) SetIndex(ind int)           { t.TargetIndex = ind }
func (t *Tmpl) MaxHP() float64             { return t.HPMax }
func (t *Tmpl) HP() float64                { return t.HPCurrent }
func (t *Tmpl) Shape() core.Shape          { return &t.Hitbox }
func (t *Tmpl) Kill()                      {} // do nothing

func (t *Tmpl) Init(x, y, size float64) {
	t.Hitbox = *core.NewCircle(x, y, size)
}
