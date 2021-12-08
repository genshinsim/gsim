package core

type Character interface {
	Init(index int) //to be called when everything including weapon and artifacts has been loaded
	Tick()          //function to be called every frame

	//information functions
	Name() string
	CharIndex() int
	Ele() EleType
	Level() int
	WeaponClass() WeaponClass
	Zone() ZoneType
	CurrentEnergy() float64 //current energy
	MaxEnergy() float64
	HP() float64
	MaxHP() float64
	ModifyHP(float64)
	Stat(s StatType) float64

	AddTask(fun func(), name string, delay int)

	//actions; each action should return 2 ints:
	//	the earliest frame at which the next action may be queued, and;
	// 	the total length of the animation state
	Attack(p map[string]int) (int, int)
	Aimed(p map[string]int) (int, int)
	ChargeAttack(p map[string]int) (int, int)
	HighPlungeAttack(p map[string]int) (int, int)
	LowPlungeAttack(p map[string]int) (int, int)
	Skill(p map[string]int) (int, int)
	Burst(p map[string]int) (int, int)
	Dash(p map[string]int) (int, int)

	//info methods
	ActionReady(a ActionType, p map[string]int) bool
	ActionStam(a ActionType, p map[string]int) float64

	//number of frames this action will take
	// ActionFrames(a ActionType, p map[string]int) int
	//return the number of frames the current action must wait before it can be
	//executed;
	ActionInterruptableDelay(next ActionType) int

	//char stat mods
	AddMod(mod CharStatMod)
	AddPreDamageMod(mod PreDamageMod)
	AddWeaponInfuse(inf WeaponInfusion)
	AddReactBonusMod(mod ReactionBonusMod)
	ReactBonus(AttackInfo) float64

	//cooldown stuff
	SetCD(a ActionType, dur int)
	Cooldown(a ActionType) int
	ResetActionCooldown(a ActionType)
	ReduceActionCooldown(a ActionType, v int)
	AddCDAdjustFunc(adj CDAdjust)

	//status stuff
	Tag(key string) int

	//energy
	QueueParticle(src string, num int, ele EleType, delay int)
	ReceiveParticle(p Particle, isActive bool, partyCount int)
	AddEnergy(e float64)

	//combat
	Snapshot(a *AttackInfo) Snapshot
	PreDamageSnapshotAdjust(*AttackEvent, Target) [EndStatType]float64
	ResetNormalCounter()
}

type ZoneType int

const (
	ZoneMondstadt ZoneType = iota
	ZoneLiyue
)

type CharStatMod struct {
	Key          string
	AffectedStat StatType
	Amount       func(a AttackTag) ([EndStatType]float64, bool)
	Expiry       int
}

type PreDamageMod struct {
	Key    string
	Amount func(atk *AttackEvent, t Target) ([EndStatType]float64, bool)
	Expiry int
}

type WeaponInfusion struct {
	Key    string
	Ele    EleType
	Tags   []AttackTag
	Expiry int
}

type CDAdjust struct {
	Key    string
	Amount func(a ActionType) float64
	Expiry int
}

type Particle struct {
	Source string
	Num    int
	Ele    EleType
}

type ReactionBonusMod struct {
	Key    string
	Amount func(AttackInfo) (float64, bool)
	Expiry int
}
