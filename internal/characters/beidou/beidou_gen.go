// Code generated by "pipeline"; DO NOT EDIT.
package beidou

import (
	_ "embed"

	"fmt"
	"github.com/genshinsim/gcsim/pkg/core/action"
	"github.com/genshinsim/gcsim/pkg/core/keys"
	"github.com/genshinsim/gcsim/pkg/gcs/validation"
	"github.com/genshinsim/gcsim/pkg/model"
	"google.golang.org/protobuf/encoding/prototext"
	"slices"
)

//go:embed data_gen.textproto
var pbData []byte
var base *model.AvatarData
var paramKeysValidation = map[action.Action][]string{
	1: {"counter"},
}

func init() {
	base = &model.AvatarData{}
	err := prototext.Unmarshal(pbData, base)
	if err != nil {
		panic(err)
	}
	validation.RegisterCharParamValidationFunc(keys.Beidou, ValidateParamKeys)
}

func ValidateParamKeys(a action.Action, keys []string) error {
	valid, ok := paramKeysValidation[a]
	if !ok {
		return nil
	}
	for _, v := range keys {
		if !slices.Contains(valid, v) {
			return fmt.Errorf("key %v is invalid for action %v", v, a.String())
		}
	}
	return nil
}

func (x *char) Data() *model.AvatarData {
	return base
}

var (
	attack = [][]float64{
		attack_1,
		attack_2,
		attack_3,
		attack_4,
		attack_5,
	}
)

var (
	// attack: attack_1 = [0]
	attack_1 = []float64{
		0.71122,
		0.76911,
		0.827,
		0.9097,
		0.96759,
		1.03375,
		1.12472,
		1.21569,
		1.30666,
		1.4059,
		1.519613,
		1.653338,
		1.787064,
		1.92079,
		2.066673,
	}
	// attack: attack_2 = [1]
	attack_2 = []float64{
		0.70864,
		0.76632,
		0.824,
		0.9064,
		0.96408,
		1.03,
		1.12064,
		1.21128,
		1.30192,
		1.4008,
		1.5141,
		1.647341,
		1.780582,
		1.913822,
		2.059176,
	}
	// attack: attack_3 = [2]
	attack_3 = []float64{
		0.88322,
		0.95511,
		1.027,
		1.1297,
		1.20159,
		1.28375,
		1.39672,
		1.50969,
		1.62266,
		1.7459,
		1.887112,
		2.053178,
		2.219244,
		2.38531,
		2.566473,
	}
	// attack: attack_4 = [3]
	attack_4 = []float64{
		0.86516,
		0.93558,
		1.006,
		1.1066,
		1.17702,
		1.2575,
		1.36816,
		1.47882,
		1.58948,
		1.7102,
		1.848525,
		2.011195,
		2.173865,
		2.336536,
		2.513994,
	}
	// attack: attack_5 = [4]
	attack_5 = []float64{
		1.12144,
		1.21272,
		1.304,
		1.4344,
		1.52568,
		1.63,
		1.77344,
		1.91688,
		2.06032,
		2.2168,
		2.3961,
		2.606957,
		2.817814,
		3.02867,
		3.258696,
	}
	// skill: shieldBase = [1]
	shieldBase = []float64{
		1386.3678,
		1525.0239,
		1675.2346,
		1837,
		2010.3201,
		2195.1948,
		2391.6243,
		2599.6084,
		2819.147,
		3050.2405,
		3292.8887,
		3547.0913,
		3812.8489,
		4090.161,
		4379.028,
	}
	// skill: shieldPer = [0]
	shieldPer = []float64{
		0.144,
		0.1548,
		0.1656,
		0.18,
		0.1908,
		0.2016,
		0.216,
		0.2304,
		0.2448,
		0.2592,
		0.2736,
		0.288,
		0.306,
		0.324,
		0.342,
	}
	// skill: skillbonus = [3]
	skillbonus = []float64{
		1.6,
		1.72,
		1.84,
		2,
		2.12,
		2.24,
		2.4,
		2.56,
		2.72,
		2.88,
		3.04,
		3.2,
		3.4,
		3.6,
		3.8,
	}
	// burst: burstonhit = [0]
	burstonhit = []float64{
		1.216,
		1.3072,
		1.3984,
		1.52,
		1.6112,
		1.7024,
		1.824,
		1.9456,
		2.0672,
		2.1888,
		2.3104,
		2.432,
		2.584,
		2.736,
		2.888,
	}
	// burst: burstproc = [1]
	burstproc = []float64{
		0.96,
		1.032,
		1.104,
		1.2,
		1.272,
		1.344,
		1.44,
		1.536,
		1.632,
		1.728,
		1.824,
		1.92,
		2.04,
		2.16,
		2.28,
	}
	// burst: skillbase = [0]
	skillbase = []float64{
		1.216,
		1.3072,
		1.3984,
		1.52,
		1.6112,
		1.7024,
		1.824,
		1.9456,
		2.0672,
		2.1888,
		2.3104,
		2.432,
		2.584,
		2.736,
		2.888,
	}
)
