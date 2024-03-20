// Code generated by "pipeline"; DO NOT EDIT.
package kaveh

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
	5: {"collision"},
	6: {"collision"},
}

func init() {
	base = &model.AvatarData{}
	err := prototext.Unmarshal(pbData, base)
	if err != nil {
		panic(err)
	}
	validation.RegisterCharParamValidationFunc(keys.Kaveh, ValidateParamKeys)
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
	}
)

var (
	// attack: attack_1 = [0]
	attack_1 = []float64{
		0.7618569731712341,
		0.8238679766654968,
		0.8858799934387207,
		0.9744679927825928,
		1.0364799499511719,
		1.1073499917984009,
		1.2047970294952393,
		1.302243947982788,
		1.39969003200531,
		1.5059959888458252,
		1.6123019456863403,
		1.718606948852539,
		1.8249130249023438,
		1.9312180280685425,
		2.0375239849090576,
	}
	// attack: attack_2 = [1]
	attack_2 = []float64{
		0.6963850259780884,
		0.7530679702758789,
		0.8097500205039978,
		0.8907250165939331,
		0.9474070072174072,
		1.0121879577636719,
		1.1012599468231201,
		1.1903330087661743,
		1.2794049978256226,
		1.376574993133545,
		1.4737449884414673,
		1.5709149837493896,
		1.668084979057312,
		1.7652549743652344,
		1.8624249696731567,
	}
	// attack: attack_3 = [2]
	attack_3 = []float64{
		0.8426110148429871,
		0.9111949801445007,
		0.9797800183296204,
		1.0777579545974731,
		1.1463429927825928,
		1.2247250080108643,
		1.3325010538101196,
		1.4402769804000854,
		1.5480519533157349,
		1.665626049041748,
		1.7832000255584717,
		1.900773048400879,
		2.0183470249176025,
		2.1359200477600098,
		2.2534940242767334,
	}
	// attack: attack_4 = [3]
	attack_4 = []float64{
		1.0268830060958862,
		1.1104669570922852,
		1.1940499544143677,
		1.3134549856185913,
		1.3970390558242798,
		1.492563009262085,
		1.6239080429077148,
		1.7552540302276611,
		1.8865989446640015,
		2.0298850536346436,
		2.173171043395996,
		2.3164570331573486,
		2.459743022918701,
		2.6030290126800537,
		2.7463150024414062,
	}
	// attack: collision = [8]
	collision = []float64{
		0.7458779811859131,
		0.8065890073776245,
		0.8672999739646912,
		0.9540299773216248,
		1.0147409439086914,
		1.08412504196167,
		1.179527997970581,
		1.2749309539794922,
		1.3703340291976929,
		1.474410057067871,
		1.5784859657287598,
		1.682561993598938,
		1.7866380214691162,
		1.8907140493392944,
		1.994789958000183,
	}
	// attack: highPlunge = [10]
	highPlunge = []float64{
		1.862889051437378,
		2.0145199298858643,
		2.1661500930786133,
		2.382765054702759,
		2.534395933151245,
		2.707688093185425,
		2.9459640979766846,
		3.1842410564422607,
		3.4225170612335205,
		3.682455062866211,
		3.9423930644989014,
		4.202331066131592,
		4.462268829345703,
		4.722207069396973,
		4.982144832611084,
	}
	// attack: lowPlunge = [9]
	lowPlunge = []float64{
		1.4914400577545166,
		1.6128360033035278,
		1.734233021736145,
		1.907655954360962,
		2.0290520191192627,
		2.1677908897399902,
		2.358556032180786,
		2.5493218898773193,
		2.7400870323181152,
		2.948194980621338,
		3.1563029289245605,
		3.3644111156463623,
		3.572519063949585,
		3.7806270122528076,
		3.9887349605560303,
	}
	// skill: skill = [0]
	skill = []float64{
		2.0399999618530273,
		2.193000078201294,
		2.3459999561309814,
		2.549999952316284,
		2.703000068664551,
		2.8559999465942383,
		3.059999942779541,
		3.2639999389648438,
		3.4679999351501465,
		3.671999931335449,
		3.875999927520752,
		4.079999923706055,
		4.335000038146973,
		4.590000152587891,
		4.84499979019165,
	}
	// burst: burst = [0]
	burst = []float64{
		1.600000023841858,
		1.7200000286102295,
		1.840000033378601,
		2,
		2.119999885559082,
		2.240000009536743,
		2.4000000953674316,
		2.559999942779541,
		2.7200000286102295,
		2.880000114440918,
		3.0399999618530273,
		3.200000047683716,
		3.4000000953674316,
		3.5999999046325684,
		3.799999952316284,
	}
	// burst: burstDmgBonus = [2]
	burstDmgBonus = []float64{
		0.27487999200820923,
		0.29549598693847656,
		0.3161120116710663,
		0.34360000491142273,
		0.36421599984169006,
		0.3848319947719574,
		0.41231998801231384,
		0.4398080110549927,
		0.4672960042953491,
		0.49478399753570557,
		0.522271990776062,
		0.5497599840164185,
		0.584119975566864,
		0.6184800267219543,
		0.6528400182723999,
	}
)
