// Code generated by "pipeline"; DO NOT EDIT.
package charlotte

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
	}
)

var (
	// attack: arkhe = [8]
	arkhe = []float64{
		0.11168000102043152,
		0.12005600333213806,
		0.1284320056438446,
		0.1395999938249588,
		0.14797599613666534,
		0.1563519984483719,
		0.16752000153064728,
		0.17868800461292267,
		0.18985599279403687,
		0.20102399587631226,
		0.21219199895858765,
		0.22336000204086304,
		0.23732000589370728,
		0.2512800097465515,
		0.26524001359939575,
	}
	// attack: attack_1 = [0]
	attack_1 = []float64{
		0.4984560012817383,
		0.5358399748802185,
		0.5732240080833435,
		0.6230700016021729,
		0.6604539752006531,
		0.6978380084037781,
		0.7476840019226074,
		0.7975299954414368,
		0.8473749756813049,
		0.897221028804779,
		0.9470660090446472,
		0.9969120025634766,
		1.0592190027236938,
		1.1215260028839111,
		1.1838330030441284,
	}
	// attack: attack_2 = [1]
	attack_2 = []float64{
		0.43375200033187866,
		0.46628299355506897,
		0.49881500005722046,
		0.5421900153160095,
		0.5747209787368774,
		0.6072530150413513,
		0.6506279706954956,
		0.6940029859542847,
		0.7373780012130737,
		0.780754029750824,
		0.8241289854049683,
		0.8675040006637573,
		0.9217230081558228,
		0.9759420156478882,
		1.0301610231399536,
	}
	// attack: attack_3 = [2]
	attack_3 = []float64{
		0.6460080146789551,
		0.6944590210914612,
		0.7429090142250061,
		0.8075100183486938,
		0.8559610247612,
		0.9044110178947449,
		0.9690120220184326,
		1.0336129665374756,
		1.098214030265808,
		1.1628140211105347,
		1.2274149656295776,
		1.2920160293579102,
		1.3727669715881348,
		1.453518033027649,
		1.5342689752578735,
	}
	// attack: charge = [3]
	charge = []float64{
		1.005120038986206,
		1.0805039405822754,
		1.1558879613876343,
		1.2563999891281128,
		1.3317840099334717,
		1.4071680307388306,
		1.507680058479309,
		1.608191967010498,
		1.7087039947509766,
		1.809216022491455,
		1.9097280502319336,
		2.010240077972412,
		2.1358799934387207,
		2.2615199089050293,
		2.387160062789917,
	}
	// attack: collision = [5]
	collision = []float64{
		0.5682880282402039,
		0.6145439743995667,
		0.6607999801635742,
		0.7268800139427185,
		0.7731360197067261,
		0.8259999752044678,
		0.898688018321991,
		0.9713760018348694,
		1.0440640449523926,
		1.12336003780365,
		1.2026560306549072,
		1.2819520235061646,
		1.3612480163574219,
		1.4405440092086792,
		1.5198400020599365,
	}
	// attack: highPlunge = [7]
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
	// attack: lowPlunge = [6]
	lowPlunge = []float64{
		1.136335015296936,
		1.228827953338623,
		1.3213200569152832,
		1.4534519910812378,
		1.5459439754486084,
		1.6516499519348145,
		1.7969950437545776,
		1.9423400163650513,
		2.087686061859131,
		2.246243953704834,
		2.404802083969116,
		2.5633609294891357,
		2.721919059753418,
		2.8804779052734375,
		3.0390360355377197,
	}
	// skill: skillHold = [1]
	skillHold = []float64{
		1.3919999599456787,
		1.496399998664856,
		1.6008000373840332,
		1.7400000095367432,
		1.8444000482559204,
		1.948799967765808,
		2.0880000591278076,
		2.2272000312805176,
		2.3664000034332275,
		2.5055999755859375,
		2.6447999477386475,
		2.7839999198913574,
		2.9579999446868896,
		3.131999969482422,
		3.305999994277954,
	}
	// skill: skillHoldMark = [5]
	skillHoldMark = []float64{
		0.4059999883174896,
		0.4364500045776367,
		0.4668999910354614,
		0.5074999928474426,
		0.5379499793052673,
		0.5684000253677368,
		0.609000027179718,
		0.6496000289916992,
		0.6901999711990356,
		0.7307999730110168,
		0.771399974822998,
		0.8119999766349792,
		0.8627499938011169,
		0.9135000109672546,
		0.9642500281333923,
	}
	// skill: skillPress = [0]
	skillPress = []float64{
		0.671999990940094,
		0.7224000096321106,
		0.7728000283241272,
		0.8399999737739563,
		0.8903999924659729,
		0.9408000111579895,
		1.0080000162124634,
		1.0751999616622925,
		1.1424000263214111,
		1.2095999717712402,
		1.2768000364303589,
		1.343999981880188,
		1.4279999732971191,
		1.5119999647140503,
		1.5959999561309814,
	}
	// skill: skillPressMark = [2]
	skillPressMark = []float64{
		0.3919999897480011,
		0.4214000105857849,
		0.45080000162124634,
		0.49000000953674316,
		0.5194000005722046,
		0.548799991607666,
		0.5879999995231628,
		0.6272000074386597,
		0.6664000153541565,
		0.7056000232696533,
		0.7447999715805054,
		0.7839999794960022,
		0.8330000042915344,
		0.8820000290870667,
		0.9309999942779541,
	}
	// burst: burst = [2]
	burst = []float64{
		0.7761600017547607,
		0.8343719840049744,
		0.8925840258598328,
		0.9702000021934509,
		1.0284119844436646,
		1.086624026298523,
		1.1642400026321411,
		1.2418559789657593,
		1.3194719552993774,
		1.3970880508422852,
		1.4747040271759033,
		1.5523200035095215,
		1.649340033531189,
		1.746359944343567,
		1.8433799743652344,
	}
	// burst: burstDot = [5]
	burstDot = []float64{
		0.06468000262975693,
		0.06953100115060806,
		0.0743819996714592,
		0.08084999769926071,
		0.08570100367069244,
		0.09055200219154358,
		0.09702000021934509,
		0.1034879982471466,
		0.10995599627494812,
		0.11642400175333023,
		0.12289199978113174,
		0.12936000525951385,
		0.13744500279426575,
		0.14553000032901764,
		0.15361499786376953,
	}
	// burst: burstDotHealFlat = [4]
	burstDotHealFlat = []float64{
		57.44709777832031,
		63.19260787963867,
		69.41690826416016,
		76.12000274658203,
		83.30188751220703,
		90.96256256103516,
		99.10203552246094,
		107.72029876708984,
		116.81735229492188,
		126.39319610595703,
		136.4478302001953,
		146.98126220703125,
		157.99349975585938,
		169.48451232910156,
		181.45431518554688,
	}
	// burst: burstDotHealPer = [3]
	burstDotHealPer = []float64{
		0.09216000139713287,
		0.09907200187444687,
		0.10598400235176086,
		0.1151999980211258,
		0.12211199849843979,
		0.12902399897575378,
		0.1382399946451187,
		0.14745600521564484,
		0.15667200088500977,
		0.1658879965543747,
		0.17510400712490082,
		0.18432000279426575,
		0.1958400011062622,
		0.20735999941825867,
		0.21887999773025513,
	}
	// burst: burstInitialHealFlat = [1]
	burstInitialHealFlat = []float64{
		1608.486328125,
		1769.3572998046875,
		1943.6341552734375,
		2131.31689453125,
		2332.40576171875,
		2546.900390625,
		2774.801025390625,
		3016.107421875,
		3270.81982421875,
		3538.938232421875,
		3820.46240234375,
		4115.392578125,
		4423.728515625,
		4745.470703125,
		5080.61865234375,
	}
	// burst: burstInitialHealPer = [0]
	burstInitialHealPer = []float64{
		2.5657339096069336,
		2.7581639289855957,
		2.950594902038574,
		3.2071681022644043,
		3.3995978832244873,
		3.5920279026031494,
		3.848602056503296,
		4.105175018310547,
		4.361748218536377,
		4.618321895599365,
		4.874895095825195,
		5.131468772888184,
		5.452186107635498,
		5.772902011871338,
		6.093618869781494,
	}
)
