// Code generated by "pipeline"; DO NOT EDIT.
package gaming

import (
	_ "embed"

	"github.com/genshinsim/gcsim/pkg/model"
	"google.golang.org/protobuf/encoding/prototext"
)

//go:embed data_gen.textproto
var pbData []byte
var base *model.AvatarData

func init() {
	base = &model.AvatarData{}
	err := prototext.Unmarshal(pbData, base)
	if err != nil {
		panic(err)
	}
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
		0.8385599851608276,
		0.9068149924278259,
		0.9750699996948242,
		1.0725769996643066,
		1.1408319473266602,
		1.2188379764556885,
		1.3260949850082397,
		1.4333529472351074,
		1.5406110286712646,
		1.6576189994812012,
		1.7746269702911377,
		1.8916360139846802,
		2.0086441040039062,
		2.125653028488159,
		2.2426609992980957,
	}
	// attack: attack_2 = [1]
	attack_2 = []float64{
		0.7904430031776428,
		0.854781985282898,
		0.9191200137138367,
		1.011031985282898,
		1.075369954109192,
		1.148900032043457,
		1.2500029802322388,
		1.35110604763031,
		1.4522099494934082,
		1.5625040531158447,
		1.6727980375289917,
		1.783092975616455,
		1.893386960029602,
		2.0036818981170654,
		2.113976001739502,
	}
	// attack: attack_3 = [2]
	attack_3 = []float64{
		1.0664600133895874,
		1.1532649993896484,
		1.2400699853897095,
		1.3640769720077515,
		1.4508819580078125,
		1.5500880479812622,
		1.6864949464797974,
		1.822903037071228,
		1.9593110084533691,
		2.108119010925293,
		2.256927013397217,
		2.405735969543457,
		2.554543972015381,
		2.703352928161621,
		2.852160930633545,
	}
	// attack: attack_4 = [3]
	attack_4 = []float64{
		1.2794909477233887,
		1.3836350440979004,
		1.487779974937439,
		1.6365580558776855,
		1.7407029867172241,
		1.859724998474121,
		2.023380994796753,
		2.1870369911193848,
		2.3506920337677,
		2.529226064682007,
		2.7077600955963135,
		2.8862929344177246,
		3.0648269653320312,
		3.2433600425720215,
		3.421894073486328,
	}
	// attack: collision = [8]
	collision = []float64{
		0.6414570212364197,
		0.693668007850647,
		0.7458800077438354,
		0.820468008518219,
		0.8726800084114075,
		0.9323499798774719,
		1.0143970251083374,
		1.096444010734558,
		1.1784900426864624,
		1.2679959535598755,
		1.3575019836425781,
		1.4470069408416748,
		1.5365129709243774,
		1.6260180473327637,
		1.7155239582061768,
	}
	// attack: highplunge = [10]
	highplunge = []float64{
		1.602084994316101,
		1.7324880361557007,
		1.8628900051116943,
		2.0491790771484375,
		2.1795809268951416,
		2.3286120891571045,
		2.5335299968719482,
		2.738447904586792,
		2.943366050720215,
		3.1669130325317383,
		3.3904600143432617,
		3.614006996154785,
		3.837553024291992,
		4.061100006103516,
		4.284646987915039,
	}
	// attack: lowplunge = [9]
	lowplunge = []float64{
		1.2826379537582397,
		1.3870389461517334,
		1.4914400577545166,
		1.6405839920043945,
		1.7449849843978882,
		1.864300012588501,
		2.028357982635498,
		2.1924169063568115,
		2.3564751148223877,
		2.5354480743408203,
		2.714421033859253,
		2.8933939933776855,
		3.0723659992218018,
		3.2513389587402344,
		3.430311918258667,
	}
	// skill: skill = [0]
	skill = []float64{
		2.303999900817871,
		2.476799964904785,
		2.649600028991699,
		2.880000114440918,
		3.052799940109253,
		3.225600004196167,
		3.4560000896453857,
		3.6863999366760254,
		3.916800022125244,
		4.147200107574463,
		4.377600193023682,
		4.607999801635742,
		4.895999908447266,
		5.184000015258789,
		5.4720001220703125,
	}
	// burst: burst = [0]
	burst = []float64{
		3.7039999961853027,
		3.981800079345703,
		4.2596001625061035,
		4.630000114440918,
		4.907800197601318,
		5.1855998039245605,
		5.556000232696533,
		5.926400184631348,
		6.296800136566162,
		6.667200088500977,
		7.037600040435791,
		7.4079999923706055,
		7.870999813079834,
		8.333999633789062,
		8.79699993133545,
	}
)
