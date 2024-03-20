// Code generated by "pipeline"; DO NOT EDIT.
package diluc

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
	validation.RegisterCharParamValidationFunc(keys.Diluc, ValidateParamKeys)
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
		0.8969799876213074,
		0.9699900150299072,
		1.0429999828338623,
		1.1473000049591064,
		1.2203099727630615,
		1.3037500381469727,
		1.4184800386428833,
		1.533210039138794,
		1.6479400396347046,
		1.7731000185012817,
		1.9165129661560059,
		2.0851659774780273,
		2.253818988800049,
		2.4224720001220703,
		2.606456995010376,
	}
	// attack: attack_2 = [1]
	attack_2 = []float64{
		0.8763399720191956,
		0.9476699829101562,
		1.0190000534057617,
		1.12090003490448,
		1.192229986190796,
		1.2737499475479126,
		1.3858400583267212,
		1.4979300498962402,
		1.6100200414657593,
		1.732300043106079,
		1.8724119663238525,
		2.037184953689575,
		2.2019569873809814,
		2.3667290210723877,
		2.546480894088745,
	}
	// attack: attack_3 = [2]
	attack_3 = []float64{
		0.9881399869918823,
		1.0685700178146362,
		1.1490000486373901,
		1.2639000415802002,
		1.3443299531936646,
		1.4362499713897705,
		1.5626399517059326,
		1.6890300512313843,
		1.8154200315475464,
		1.9532999992370605,
		2.1112871170043945,
		2.2970809936523438,
		2.4828739166259766,
		2.6686670780181885,
		2.8713510036468506,
	}
	// attack: attack_4 = [3]
	attack_4 = []float64{
		1.3398799896240234,
		1.4489400386810303,
		1.5579999685287476,
		1.7137999534606934,
		1.8228600025177002,
		1.9474999904632568,
		2.118880033493042,
		2.290260076522827,
		2.461639881134033,
		2.6486001014709473,
		2.8628249168395996,
		3.1147539615631104,
		3.3666820526123047,
		3.6186110973358154,
		3.893441915512085,
	}
	// attack: collision = [8]
	collision = []float64{
		0.8950539827346802,
		0.9679070115089417,
		1.0407600402832031,
		1.1448359489440918,
		1.217689037322998,
		1.300950050354004,
		1.4154340028762817,
		1.5299170017242432,
		1.644400954246521,
		1.7692919969558716,
		1.8941830396652222,
		2.019073963165283,
		2.14396595954895,
		2.268857002258301,
		2.3937480449676514,
	}
	// attack: highPlunge = [10]
	highPlunge = []float64{
		2.235466957092285,
		2.4174230098724365,
		2.5993800163269043,
		2.8593180179595947,
		3.0412750244140625,
		3.249224901199341,
		3.5351569652557373,
		3.821089029312134,
		4.107019901275635,
		4.418945789337158,
		4.73087215423584,
		5.042797088623047,
		5.35472297668457,
		5.666647911071777,
		5.978573799133301,
	}
	// attack: lowPlunge = [9]
	lowPlunge = []float64{
		1.789728045463562,
		1.9354029893875122,
		2.0810790061950684,
		2.289186954498291,
		2.434861898422241,
		2.601349115371704,
		2.8302669525146484,
		3.0591859817504883,
		3.288105010986328,
		3.5378339290618896,
		3.7875640392303467,
		4.037292957305908,
		4.287023067474365,
		4.536752223968506,
		4.786481857299805,
	}
	// skill: skill = [0 1 2]
	skill = [][]float64{
		{
			0.9440000057220459,
			1.014799952507019,
			1.0856000185012817,
			1.1799999475479126,
			1.2508000135421753,
			1.3215999603271484,
			1.4160000085830688,
			1.5104000568389893,
			1.6047999858856201,
			1.6992000341415405,
			1.7935999631881714,
			1.8880000114440918,
			2.00600004196167,
			2.124000072479248,
			2.242000102996826,
		},
		{
			0.9760000109672546,
			1.0492000579833984,
			1.1224000453948975,
			1.2200000286102295,
			1.2932000160217285,
			1.3664000034332275,
			1.4639999866485596,
			1.5615999698638916,
			1.6591999530792236,
			1.7568000555038452,
			1.8544000387191772,
			1.9520000219345093,
			2.0739998817443848,
			2.196000099182129,
			2.318000078201294,
		},
		{
			1.2879999876022339,
			1.384600043296814,
			1.4811999797821045,
			1.6100000143051147,
			1.7065999507904053,
			1.8032000064849854,
			1.9320000410079956,
			2.060800075531006,
			2.1895999908447266,
			2.3183999061584473,
			2.447200059890747,
			2.5759999752044678,
			2.736999988555908,
			2.8980000019073486,
			3.059000015258789,
		},
	}
	// burst: burstDOT = [1]
	burstDOT = []float64{
		0.6000000238418579,
		0.6449999809265137,
		0.6899999976158142,
		0.75,
		0.7950000166893005,
		0.8399999737739563,
		0.8999999761581421,
		0.9599999785423279,
		1.0199999809265137,
		1.0800000429153442,
		1.1399999856948853,
		1.2000000476837158,
		1.274999976158142,
		1.350000023841858,
		1.4249999523162842,
	}
	// burst: burstExplode = [0]
	burstExplode = []float64{
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
	// burst: burstInitial = [0]
	burstInitial = []float64{
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
)
