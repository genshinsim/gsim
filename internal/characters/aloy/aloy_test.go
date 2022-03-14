package aloy

import (
	"testing"

	"github.com/genshinsim/gcsim/internal/testhelper"
	"github.com/genshinsim/gcsim/pkg/core"
	"github.com/genshinsim/gcsim/pkg/coretype"
)

func TestBasicAbilUsage(t *testing.T) {
	c := testhelper.NewTestCore()
	prof := testhelper.CharProfile(core.Aloy, coretype.Cryo, 6)
	x, err := NewChar(c, prof)
	//cast it to *char so we can access private members
	// this := x.(*char)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	testhelper.TestBowCharacter(c, x)
}

func TestCD(t *testing.T) {
	c := testhelper.NewTestCore()
	prof := testhelper.CharProfile(core.Aloy, core.Cryo, 6)
	x, err := NewChar(c, prof)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	err = testhelper.TestSkillCDSingleCharge(c, x, 20*60)
	if err != nil {
		t.Error(err)
	}
}
