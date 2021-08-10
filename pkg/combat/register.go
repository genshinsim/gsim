package combat

import (
	"sync"

	"github.com/genshinsim/gsim/pkg/core"
	"go.uber.org/zap"
)

var (
	mu        sync.RWMutex
	charMap   = make(map[string]NewCharacterFunc)
	setMap    = make(map[string]NewSetFunc)
	weaponMap = make(map[string]NewWeaponFunc)
)

type NewCharacterFunc func(s core.Sim, log *zap.SugaredLogger, p core.CharacterProfile) (core.Character, error)
type NewSetFunc func(c core.Character, s core.Sim, log core.Logger, count int)
type NewWeaponFunc func(c core.Character, s core.Sim, log core.Logger, r int, param map[string]int)

func RegisterCharFunc(name string, f NewCharacterFunc) {
	mu.Lock()
	defer mu.Unlock()
	if _, dup := charMap[name]; dup {
		panic("combat: RegisterChar called twice for character " + name)
	}
	charMap[name] = f
}

func RegisterSetFunc(name string, f NewSetFunc) {
	mu.Lock()
	defer mu.Unlock()
	if _, dup := setMap[name]; dup {
		panic("combat: RegisterSetBonus called twice for character " + name)
	}
	setMap[name] = f
}

func RegisterWeaponFunc(name string, f NewWeaponFunc) {
	mu.Lock()
	defer mu.Unlock()
	if _, dup := weaponMap[name]; dup {
		panic("combat: RegisterWeapon called twice for character " + name)
	}
	weaponMap[name] = f
}
