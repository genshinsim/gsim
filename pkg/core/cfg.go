package core

//SimulationConfig describes the required settings to run an simulation
type SimulationConfig struct {
	//these settings relate to each simulation iteration
	DamageMode bool
	Targets    []EnemyProfile
	Characters struct {
		Initial CharKey
		Profile []CharacterProfile
	}
	Rotation []ActionBlock
	Hurt     HurtEvent
	Energy   EnergyEvent
	Duration int //duration of the simulation in seconds; only if DamageMode is false

	Options SimulationOption
}

func (c *SimulationConfig) Clone() SimulationConfig {
	r := *c

	//clone targets
	r.Targets = make([]EnemyProfile, len(c.Targets))

	for i, v := range c.Targets {
		r.Targets[i] = v.Clone()
	}

	//clone characters
	r.Characters.Profile = make([]CharacterProfile, len(c.Characters.Profile))
	for i, v := range c.Characters.Profile {
		r.Characters.Profile[i] = v.Clone()
	}

	return r
}

//SimulationOption contains options such as number of iteration, number of workers, etc..
type SimulationOption struct {
	Iteration     int  `json:"iter"`
	Workers       int  `json:"workers"`
	GenerateDebug bool `json:"debug"`
	LogDetails    bool `json:"log_details"`
	ERCalcMode    bool `json:"er_calc_mode"`
}

type CharacterProfile struct {
	Base      CharacterBase             `json:"base"`
	Weapon    WeaponProfile             `json:"weapon"`
	Talents   TalentProfile             `json:"talents"`
	Stats     []float64                 `json:"stats"`
	Sets      map[string]int            `json:"sets"`
	SetParams map[string]map[string]int `json:"-"`
	Params    map[string]int            `json:"-"`
}

func (c *CharacterProfile) Clone() CharacterProfile {
	r := *c
	r.Weapon.Params = make(map[string]int)
	for k, v := range c.Weapon.Params {
		r.Weapon.Params[k] = v
	}
	r.Stats = make([]float64, len(c.Stats))
	copy(r.Stats, c.Stats)
	r.Sets = make(map[string]int)
	for k, v := range c.Sets {
		r.Sets[k] = v
	}

	return r
}

type CharacterBase struct {
	Key      CharKey `json:"-"`
	Name     string  `json:"name"`
	Element  EleType `json:"element"`
	Level    int     `json:"level"`
	MaxLevel int     `json:"max_level"`
	HP       float64 `json:"-"`
	Atk      float64 `json:"-"`
	Def      float64 `json:"-"`
	Cons     int     `json:"cons"`
	StartHP  float64 `json:"-"`
}

type WeaponProfile struct {
	Name     string         `json:"name"`
	Key      string         `json:""` //use this to match with weapon curve mapping
	Class    WeaponClass    `json:"-"`
	Refine   int            `json:"refine"`
	Level    int            `json:"level"`
	MaxLevel int            `json:"max_level"`
	Atk      float64        `json:"-"`
	Params   map[string]int `json:"-"`
}

type TalentProfile struct {
	Attack int `json:"attack"`
	Skill  int `json:"skill"`
	Burst  int `json:"burst"`
}

type EnemyProfile struct {
	Level          int                 `json:"level"`
	HP             float64             `json:"-"`
	Resist         map[EleType]float64 `json:"-"`
	Size           float64             `json:"-"`
	CoordX, CoordY float64             `json:"-"`
}

func (e *EnemyProfile) Clone() EnemyProfile {
	r := EnemyProfile{
		Level:  e.Level,
		Resist: make(map[EleType]float64),
	}
	for k, v := range e.Resist {
		r.Resist[k] = v
	}
	return r
}

type EnergyEvent struct {
	Active    bool
	Once      bool //how often
	Start     int
	End       int
	Particles int
}

type HurtEvent struct {
	Active bool
	Once   bool //how often
	Start  int  //
	End    int
	Min    float64
	Max    float64
	Ele    EleType
}
