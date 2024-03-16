// Code generated by "pipeline"; DO NOT EDIT.
package hutao

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
	attack = [][][]float64{
		{attack_1},
		{attack_2},
		{attack_3},
		{attack_4},
		attack_5,
		{attack_6},
	}
)

var (
	// attack: attack_1 = [0]
	attack_1 = []float64{
		0.46886399388313293,
		0.5008320212364197,
		0.532800018787384,
		0.5754240155220032,
		0.6073920130729675,
		0.6446880102157593,
		0.6926400065422058,
		0.7405920028686523,
		0.7885439991950989,
		0.8364959955215454,
		0.8844479918479919,
		0.9323999881744385,
		0.980351984500885,
		1.0283039808273315,
		1.0762560367584229,
	}
	// attack: attack_2 = [1]
	attack_2 = []float64{
		0.48253899812698364,
		0.5154399871826172,
		0.5483400225639343,
		0.5922070145606995,
		0.625108003616333,
		0.6634910106658936,
		0.7128419876098633,
		0.7621930241584778,
		0.8115429878234863,
		0.8608940243721008,
		0.9102439880371094,
		0.9595950245857239,
		1.0089459419250488,
		1.0582959651947021,
		1.1076469421386719,
	}
	// attack: attack_3 = [2]
	attack_3 = []float64{
		0.6104999780654907,
		0.6521250009536743,
		0.6937500238418579,
		0.7492499947547913,
		0.7908750176429749,
		0.8394380211830139,
		0.9018750190734863,
		0.9643129706382751,
		1.0267499685287476,
		1.0891879796981812,
		1.1516250371932983,
		1.214063048362732,
		1.2764999866485596,
		1.3389379978179932,
		1.4013750553131104,
	}
	// attack: attack_4 = [3]
	attack_4 = []float64{
		0.6564099788665771,
		0.7011650204658508,
		0.7459200024604797,
		0.8055940270423889,
		0.8503490090370178,
		0.9025629758834839,
		0.9696959853172302,
		1.0368289947509766,
		1.1039619445800781,
		1.1710939407348633,
		1.2382270097732544,
		1.305359959602356,
		1.372493028640747,
		1.4396259784698486,
		1.5067579746246338,
	}
	// attack: attack_5 = [4 5]
	attack_5 = [][]float64{
		{
			0.332736998796463,
			0.3554230034351349,
			0.37810999155044556,
			0.40835899114608765,
			0.4310449957847595,
			0.4575130045413971,
			0.4915429949760437,
			0.5255730152130127,
			0.5596029758453369,
			0.5936329960823059,
			0.6276630163192749,
			0.6616929769515991,
			0.6957219839096069,
			0.7297520041465759,
			0.7637820243835449,
		},
		{
			0.35199999809265137,
			0.37599998712539673,
			0.4000000059604645,
			0.4320000112056732,
			0.4560000002384186,
			0.48399999737739563,
			0.5199999809265137,
			0.5559999942779541,
			0.5920000076293945,
			0.628000020980835,
			0.6639999747276306,
			0.699999988079071,
			0.7360000014305115,
			0.7720000147819519,
			0.8080000281333923,
		},
	}
	// attack: attack_6 = [6]
	attack_6 = []float64{
		0.8595839738845825,
		0.9181920289993286,
		0.9768000245094299,
		1.0549440383911133,
		1.1135519742965698,
		1.181928038597107,
		1.2698400020599365,
		1.3577519655227661,
		1.4456640481948853,
		1.5335760116577148,
		1.6214879751205444,
		1.7094000577926636,
		1.7973120212554932,
		1.8852239847183228,
		1.9731359481811523,
	}
	// attack: charge = [7]
	charge = []float64{
		1.3595999479293823,
		1.452299952507019,
		1.5449999570846558,
		1.6685999631881714,
		1.761299967765808,
		1.8694499731063843,
		2.008500099182129,
		2.147550106048584,
		2.286600112915039,
		2.425649881362915,
		2.56469988822937,
		2.703749895095825,
		2.8427999019622803,
		2.9818499088287354,
		3.1208999156951904,
	}
	// attack: collision = [9]
	collision = []float64{
		0.6541919708251953,
		0.6987959742546082,
		0.743399977684021,
		0.8028720021247864,
		0.8474760055541992,
		0.8995140194892883,
		0.9664199948310852,
		1.0333260297775269,
		1.1002320051193237,
		1.1671379804611206,
		1.2340439558029175,
		1.300950050354004,
		1.3678560256958008,
		1.4347620010375977,
		1.5016679763793945,
	}
	// attack: highPlunge = [11]
	highPlunge = []float64{
		1.633895993232727,
		1.7452980279922485,
		1.8566999435424805,
		2.0052359104156494,
		2.116637945175171,
		2.2466070652008057,
		2.413710117340088,
		2.580812931060791,
		2.7479159832000732,
		2.9150190353393555,
		3.0821220874786377,
		3.249224901199341,
		3.416327953338623,
		3.5834310054779053,
		3.7505340576171875,
	}
	// attack: lowPlunge = [10]
	lowPlunge = []float64{
		1.3081070184707642,
		1.3972959518432617,
		1.4864850044250488,
		1.605404019355774,
		1.6945929527282715,
		1.7986470460891724,
		1.9324309825897217,
		2.066214084625244,
		2.199997901916504,
		2.3337810039520264,
		2.4675650596618652,
		2.601349115371704,
		2.7351319789886475,
		2.8689160346984863,
		3.002700090408325,
	}
	// skill: bb = [2]
	bb = []float64{
		0.6399999856948853,
		0.6880000233650208,
		0.7360000014305115,
		0.800000011920929,
		0.8479999899864197,
		0.8960000276565552,
		0.9599999785423279,
		1.0240000486373901,
		1.0880000591278076,
		1.1519999504089355,
		1.215999960899353,
		1.2799999713897705,
		1.3600000143051147,
		1.440000057220459,
		1.5199999809265137,
	}
	// skill: ppatk = [1]
	ppatk = []float64{
		0.0384100005030632,
		0.040709998458623886,
		0.04301000013947487,
		0.04600000008940697,
		0.04830000177025795,
		0.050599999725818634,
		0.05358999967575073,
		0.05657999962568283,
		0.05956999957561493,
		0.06255999952554703,
		0.06554999947547913,
		0.06853999942541122,
		0.07152999937534332,
		0.07451999932527542,
		0.07750999927520752,
	}
	// burst: burst = [0]
	burst = []float64{
		3.0327200889587402,
		3.214319944381714,
		3.3959200382232666,
		3.631999969482422,
		3.8136000633239746,
		3.9951999187469482,
		4.2312798500061035,
		4.467360019683838,
		4.703440189361572,
		4.939519882202148,
		5.175600051879883,
		5.411680221557617,
		5.647759914398193,
		5.883840084075928,
		6.119919776916504,
	}
	// burst: burstLow = [1]
	burstLow = []float64{
		3.7908999919891357,
		4.017899990081787,
		4.244900226593018,
		4.539999961853027,
		4.767000198364258,
		4.99399995803833,
		5.289100170135498,
		5.584199905395508,
		5.879300117492676,
		6.1743998527526855,
		6.4695000648498535,
		6.764599800109863,
		7.059700012207031,
		7.354800224304199,
		7.649899959564209,
	}
	// burst: regen = [2]
	regen = []float64{
		0.06262499839067459,
		0.0663750022649765,
		0.07012499868869781,
		0.07500000298023224,
		0.07874999940395355,
		0.08250000327825546,
		0.08737500011920929,
		0.09224999696016312,
		0.09712500125169754,
		0.10199999809265137,
		0.10687500238418579,
		0.11174999922513962,
		0.11662500351667404,
		0.12150000035762787,
		0.1263750046491623,
	}
	// burst: regenLow = [3]
	regenLow = []float64{
		0.08349999785423279,
		0.0885000005364418,
		0.09350000321865082,
		0.10000000149011612,
		0.10499999672174454,
		0.10999999940395355,
		0.11649999767541885,
		0.12300000339746475,
		0.12950000166893005,
		0.13600000739097595,
		0.14249999821186066,
		0.14900000393390656,
		0.15549999475479126,
		0.16200000047683716,
		0.16850000619888306,
	}
)
