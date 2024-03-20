// Code generated by "pipeline"; DO NOT EDIT.
package venti

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
	1: {"hold"},
	3: {"travel"},
	7: {"hold", "travel", "weakspot"},
}

func init() {
	base = &model.AvatarData{}
	err := prototext.Unmarshal(pbData, base)
	if err != nil {
		panic(err)
	}
	validation.RegisterCharParamValidationFunc(keys.Venti, ValidateParamKeys)
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
	attack = [][][]float64{
		attack_1,
		{attack_2},
		{attack_3},
		attack_4,
		{attack_5},
		{attack_6},
	}
)

var (
	// attack: aim = [6]
	aim = []float64{
		0.43860000371932983,
		0.47429999709129333,
		0.5099999904632568,
		0.5609999895095825,
		0.5967000126838684,
		0.637499988079071,
		0.6935999989509583,
		0.7497000098228455,
		0.8058000206947327,
		0.8669999837875366,
		0.937125027179718,
		1.019592046737671,
		1.102059006690979,
		1.184525966644287,
		1.2744899988174438,
	}
	// attack: attack_1 = [0 0]
	attack_1 = [][]float64{
		{
			0.20382000505924225,
			0.22041000425815582,
			0.2370000034570694,
			0.260699987411499,
			0.2772899866104126,
			0.29624998569488525,
			0.32232001423835754,
			0.34839001297950745,
			0.37446001172065735,
			0.40290001034736633,
			0.4354870021343231,
			0.47380998730659485,
			0.512133002281189,
			0.5504559874534607,
			0.5922629833221436,
		},
		{
			0.20382000505924225,
			0.22041000425815582,
			0.2370000034570694,
			0.260699987411499,
			0.2772899866104126,
			0.29624998569488525,
			0.32232001423835754,
			0.34839001297950745,
			0.37446001172065735,
			0.40290001034736633,
			0.4354870021343231,
			0.47380998730659485,
			0.512133002281189,
			0.5504559874534607,
			0.5922629833221436,
		},
	}
	// attack: attack_2 = [1]
	attack_2 = []float64{
		0.4437600076198578,
		0.4798800051212311,
		0.515999972820282,
		0.5676000118255615,
		0.6037200093269348,
		0.6449999809265137,
		0.7017599940299988,
		0.7585200071334839,
		0.815280020236969,
		0.8772000074386597,
		0.9481499791145325,
		1.03158700466156,
		1.1150239706039429,
		1.1984620094299316,
		1.2894840240478516,
	}
	// attack: attack_3 = [2]
	attack_3 = []float64{
		0.5237399935722351,
		0.5663700103759766,
		0.609000027179718,
		0.6699000000953674,
		0.7125300168991089,
		0.7612500190734863,
		0.8282399773597717,
		0.8952299952507019,
		0.9622200131416321,
		1.0353000164031982,
		1.1190370321273804,
		1.2175129652023315,
		1.3159879446029663,
		1.4144630432128906,
		1.5218909978866577,
	}
	// attack: attack_4 = [3 3]
	attack_4 = [][]float64{
		{
			0.26058000326156616,
			0.28178998827934265,
			0.30300000309944153,
			0.33329999446868896,
			0.35451000928878784,
			0.3787499964237213,
			0.41207998991012573,
			0.44541001319885254,
			0.47874000668525696,
			0.5151000022888184,
			0.5567619800567627,
			0.605758011341095,
			0.6547530293464661,
			0.7037479877471924,
			0.7571970224380493,
		},
		{
			0.26058000326156616,
			0.28178998827934265,
			0.30300000309944153,
			0.33329999446868896,
			0.35451000928878784,
			0.3787499964237213,
			0.41207998991012573,
			0.44541001319885254,
			0.47874000668525696,
			0.5151000022888184,
			0.5567619800567627,
			0.605758011341095,
			0.6547530293464661,
			0.7037479877471924,
			0.7571970224380493,
		},
	}
	// attack: attack_5 = [4]
	attack_5 = []float64{
		0.5065400004386902,
		0.5477700233459473,
		0.5889999866485596,
		0.6478999853134155,
		0.6891300082206726,
		0.7362499833106995,
		0.80103999376297,
		0.8658300042152405,
		0.930620014667511,
		1.0012999773025513,
		1.0822880268096924,
		1.177528977394104,
		1.2727700471878052,
		1.3680109977722168,
		1.4719109535217285,
	}
	// attack: attack_6 = [5]
	attack_6 = []float64{
		0.7095000147819519,
		0.7672500014305115,
		0.824999988079071,
		0.9075000286102295,
		0.9652500152587891,
		1.03125,
		1.121999979019165,
		1.21274995803833,
		1.3035000562667847,
		1.402500033378601,
		1.5159369707107544,
		1.649340033531189,
		1.7827420234680176,
		1.9161449670791626,
		2.0616750717163086,
	}
	// attack: fullaim = [7]
	fullaim = []float64{
		1.2400000095367432,
		1.3329999446868896,
		1.4259999990463257,
		1.5499999523162842,
		1.6430000066757202,
		1.7359999418258667,
		1.8600000143051147,
		1.9839999675750732,
		2.1080000400543213,
		2.2320001125335693,
		2.360960006713867,
		2.529599905014038,
		2.698240041732788,
		2.866879940032959,
		3.035520076751709,
	}
	// attack: highPlunge = [10]
	highPlunge = []float64{
		1.4193439483642578,
		1.534872055053711,
		1.6504000425338745,
		1.815440058708191,
		1.9309680461883545,
		2.062999963760376,
		2.24454402923584,
		2.4260880947113037,
		2.6076319217681885,
		2.8056800365448,
		3.003727912902832,
		3.2017760276794434,
		3.3998239040374756,
		3.597872018814087,
		3.795919895172119,
	}
	// skill: skillHold = [2]
	skillHold = []float64{
		3.799999952316284,
		4.085000038146973,
		4.369999885559082,
		4.75,
		5.034999847412109,
		5.320000171661377,
		5.699999809265137,
		6.079999923706055,
		6.460000038146973,
		6.840000152587891,
		7.21999979019165,
		7.599999904632568,
		8.074999809265137,
		8.550000190734863,
		9.024999618530273,
	}
	// skill: skillPress = [0]
	skillPress = []float64{
		2.759999990463257,
		2.9670000076293945,
		3.1740000247955322,
		3.450000047683716,
		3.6570000648498535,
		3.864000082015991,
		4.139999866485596,
		4.415999889373779,
		4.691999912261963,
		4.9679999351501465,
		5.24399995803833,
		5.519999980926514,
		5.864999771118164,
		6.210000038146973,
		6.554999828338623,
	}
	// burst: burstAbsorbDot = [1]
	burstAbsorbDot = []float64{
		0.18799999356269836,
		0.2020999938249588,
		0.21619999408721924,
		0.23499999940395355,
		0.249099999666214,
		0.2632000148296356,
		0.28200000524520874,
		0.30079999566078186,
		0.319599986076355,
		0.3384000062942505,
		0.3571999967098236,
		0.37599998712539673,
		0.3995000123977661,
		0.4230000078678131,
		0.4465000033378601,
	}
	// burst: burstDot = [0]
	burstDot = []float64{
		0.37599998712539673,
		0.4041999876499176,
		0.4323999881744385,
		0.4699999988079071,
		0.498199999332428,
		0.5264000296592712,
		0.5640000104904175,
		0.6015999913215637,
		0.63919997215271,
		0.676800012588501,
		0.7143999934196472,
		0.7519999742507935,
		0.7990000247955322,
		0.8460000157356262,
		0.8930000066757202,
	}
)
