package chongyun

import (
	"testing"

	"github.com/genshinsim/gcsim/internal/testhelper"
	"github.com/genshinsim/gcsim/pkg/core"
	"github.com/genshinsim/gcsim/pkg/coretype"
)

func TestBasicAbilUsage(t *testing.T) {
	c := testhelper.NewTestCore()
	prof := testhelper.CharProfile(core.Chongyun, coretype.Cryo, 6)
	x, err := NewChar(c, prof)
	//cast it to *char so we can access private members
	// this := x.(*char)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	testhelper.TestClaymoreCharacter(c, x)
}

func TestCD(t *testing.T) {
	c := testhelper.NewTestCore()
	prof := testhelper.CharProfile(core.Chongyun, core.Cryo, 6)
	x, err := NewChar(c, prof)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	err = testhelper.TestSkillCDSingleCharge(c, x, 900)
	if err != nil {
		t.Error(err)
	}
}
