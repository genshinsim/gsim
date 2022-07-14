package simulation

import (
	//artifacts
	_ "github.com/genshinsim/gcsim/internal/artifacts/archaic"
	_ "github.com/genshinsim/gcsim/internal/artifacts/blizzard"
	_ "github.com/genshinsim/gcsim/internal/artifacts/bloodstained"
	_ "github.com/genshinsim/gcsim/internal/artifacts/bolide"
	_ "github.com/genshinsim/gcsim/internal/artifacts/crimson"
	_ "github.com/genshinsim/gcsim/internal/artifacts/echoes"
	_ "github.com/genshinsim/gcsim/internal/artifacts/exile"
	_ "github.com/genshinsim/gcsim/internal/artifacts/gambler"
	_ "github.com/genshinsim/gcsim/internal/artifacts/gladiator"
	_ "github.com/genshinsim/gcsim/internal/artifacts/heartofdepth"
	_ "github.com/genshinsim/gcsim/internal/artifacts/huskofopulentdreams"
	_ "github.com/genshinsim/gcsim/internal/artifacts/instructor"
	_ "github.com/genshinsim/gcsim/internal/artifacts/lavawalker"
	_ "github.com/genshinsim/gcsim/internal/artifacts/maiden"
	_ "github.com/genshinsim/gcsim/internal/artifacts/noblesse"
	_ "github.com/genshinsim/gcsim/internal/artifacts/oceanhuedclam"
	_ "github.com/genshinsim/gcsim/internal/artifacts/paleflame"
	_ "github.com/genshinsim/gcsim/internal/artifacts/reminiscence"
	_ "github.com/genshinsim/gcsim/internal/artifacts/scholar"
	_ "github.com/genshinsim/gcsim/internal/artifacts/seal"
	_ "github.com/genshinsim/gcsim/internal/artifacts/tenacity"
	_ "github.com/genshinsim/gcsim/internal/artifacts/thunderingfury"
	_ "github.com/genshinsim/gcsim/internal/artifacts/thundersoother"
	_ "github.com/genshinsim/gcsim/internal/artifacts/vermillion"
	_ "github.com/genshinsim/gcsim/internal/artifacts/viridescent"
	_ "github.com/genshinsim/gcsim/internal/artifacts/wanderer"

	//char
	_ "github.com/genshinsim/gcsim/internal/characters/albedo"
	_ "github.com/genshinsim/gcsim/internal/characters/aloy"
	_ "github.com/genshinsim/gcsim/internal/characters/amber"
	_ "github.com/genshinsim/gcsim/internal/characters/ayaka"
	_ "github.com/genshinsim/gcsim/internal/characters/ayato"
	_ "github.com/genshinsim/gcsim/internal/characters/barbara"
	_ "github.com/genshinsim/gcsim/internal/characters/beidou"
	_ "github.com/genshinsim/gcsim/internal/characters/bennett"
	_ "github.com/genshinsim/gcsim/internal/characters/chongyun"
	_ "github.com/genshinsim/gcsim/internal/characters/diluc"
	_ "github.com/genshinsim/gcsim/internal/characters/diona"
	_ "github.com/genshinsim/gcsim/internal/characters/eula"
	_ "github.com/genshinsim/gcsim/internal/characters/fischl"
	_ "github.com/genshinsim/gcsim/internal/characters/ganyu"
	_ "github.com/genshinsim/gcsim/internal/characters/gorou"
	_ "github.com/genshinsim/gcsim/internal/characters/heizou"
	_ "github.com/genshinsim/gcsim/internal/characters/hutao"
	_ "github.com/genshinsim/gcsim/internal/characters/itto"
	_ "github.com/genshinsim/gcsim/internal/characters/jean"
	_ "github.com/genshinsim/gcsim/internal/characters/kaeya"
	_ "github.com/genshinsim/gcsim/internal/characters/kazuha"
	_ "github.com/genshinsim/gcsim/internal/characters/keqing"
	_ "github.com/genshinsim/gcsim/internal/characters/klee"
	_ "github.com/genshinsim/gcsim/internal/characters/kokomi"
	_ "github.com/genshinsim/gcsim/internal/characters/kuki"
	_ "github.com/genshinsim/gcsim/internal/characters/lisa"
	_ "github.com/genshinsim/gcsim/internal/characters/mona"
	_ "github.com/genshinsim/gcsim/internal/characters/ningguang"
	_ "github.com/genshinsim/gcsim/internal/characters/noelle"
	_ "github.com/genshinsim/gcsim/internal/characters/qiqi"
	_ "github.com/genshinsim/gcsim/internal/characters/raiden"
	_ "github.com/genshinsim/gcsim/internal/characters/rosaria"
	_ "github.com/genshinsim/gcsim/internal/characters/sara"
	_ "github.com/genshinsim/gcsim/internal/characters/shenhe"
	_ "github.com/genshinsim/gcsim/internal/characters/sucrose"
	_ "github.com/genshinsim/gcsim/internal/characters/tartaglia"
	_ "github.com/genshinsim/gcsim/internal/characters/travelerelectro"
	_ "github.com/genshinsim/gcsim/internal/characters/travelergeo"
	_ "github.com/genshinsim/gcsim/internal/characters/venti"
	_ "github.com/genshinsim/gcsim/internal/characters/xiangling"
	_ "github.com/genshinsim/gcsim/internal/characters/xiao"
	_ "github.com/genshinsim/gcsim/internal/characters/xingqiu"
	_ "github.com/genshinsim/gcsim/internal/characters/yaemiko"
	_ "github.com/genshinsim/gcsim/internal/characters/yanfei"
	_ "github.com/genshinsim/gcsim/internal/characters/yelan"
	_ "github.com/genshinsim/gcsim/internal/characters/yoimiya"
	_ "github.com/genshinsim/gcsim/internal/characters/yunjin"
	_ "github.com/genshinsim/gcsim/internal/characters/zhongli"

	//weapons
	_ "github.com/genshinsim/gcsim/internal/weapons/bow/alley"
	_ "github.com/genshinsim/gcsim/internal/weapons/bow/amos"
	_ "github.com/genshinsim/gcsim/internal/weapons/bow/aqua"
	_ "github.com/genshinsim/gcsim/internal/weapons/bow/blackcliff"
	_ "github.com/genshinsim/gcsim/internal/weapons/bow/compound"
	_ "github.com/genshinsim/gcsim/internal/weapons/bow/elegy"
	_ "github.com/genshinsim/gcsim/internal/weapons/bow/favonius"
	_ "github.com/genshinsim/gcsim/internal/weapons/bow/hamayumi"
	_ "github.com/genshinsim/gcsim/internal/weapons/bow/huntersbow"
	_ "github.com/genshinsim/gcsim/internal/weapons/bow/mitternachtswaltz"
	_ "github.com/genshinsim/gcsim/internal/weapons/bow/mouunsmoon"
	_ "github.com/genshinsim/gcsim/internal/weapons/bow/polarstar"
	_ "github.com/genshinsim/gcsim/internal/weapons/bow/predator"
	_ "github.com/genshinsim/gcsim/internal/weapons/bow/prototype"
	_ "github.com/genshinsim/gcsim/internal/weapons/bow/recurve"
	_ "github.com/genshinsim/gcsim/internal/weapons/bow/royal"
	_ "github.com/genshinsim/gcsim/internal/weapons/bow/rust"
	_ "github.com/genshinsim/gcsim/internal/weapons/bow/sacrificial"
	_ "github.com/genshinsim/gcsim/internal/weapons/bow/sharpshooter"
	_ "github.com/genshinsim/gcsim/internal/weapons/bow/skyward"
	_ "github.com/genshinsim/gcsim/internal/weapons/bow/slingshot"
	_ "github.com/genshinsim/gcsim/internal/weapons/bow/stringless"
	_ "github.com/genshinsim/gcsim/internal/weapons/bow/thundering"
	_ "github.com/genshinsim/gcsim/internal/weapons/bow/twilight"
	_ "github.com/genshinsim/gcsim/internal/weapons/bow/viridescent"
	_ "github.com/genshinsim/gcsim/internal/weapons/bow/windblume"
	_ "github.com/genshinsim/gcsim/internal/weapons/catalyst/apprentice"
	_ "github.com/genshinsim/gcsim/internal/weapons/catalyst/blackcliff"
	_ "github.com/genshinsim/gcsim/internal/weapons/catalyst/dodoco"
	_ "github.com/genshinsim/gcsim/internal/weapons/catalyst/favonius"
	_ "github.com/genshinsim/gcsim/internal/weapons/catalyst/frostbearer"
	_ "github.com/genshinsim/gcsim/internal/weapons/catalyst/hakushin"
	_ "github.com/genshinsim/gcsim/internal/weapons/catalyst/kagura"
	_ "github.com/genshinsim/gcsim/internal/weapons/catalyst/magicguide"
	_ "github.com/genshinsim/gcsim/internal/weapons/catalyst/mappa"
	_ "github.com/genshinsim/gcsim/internal/weapons/catalyst/memory"
	_ "github.com/genshinsim/gcsim/internal/weapons/catalyst/moonglow"
	_ "github.com/genshinsim/gcsim/internal/weapons/catalyst/oathsworneye"
	_ "github.com/genshinsim/gcsim/internal/weapons/catalyst/perception"
	_ "github.com/genshinsim/gcsim/internal/weapons/catalyst/prayer"
	_ "github.com/genshinsim/gcsim/internal/weapons/catalyst/prototype"
	_ "github.com/genshinsim/gcsim/internal/weapons/catalyst/royal"
	_ "github.com/genshinsim/gcsim/internal/weapons/catalyst/sacrifical"
	_ "github.com/genshinsim/gcsim/internal/weapons/catalyst/skyward"
	_ "github.com/genshinsim/gcsim/internal/weapons/catalyst/solar"
	_ "github.com/genshinsim/gcsim/internal/weapons/catalyst/thrilling"
	_ "github.com/genshinsim/gcsim/internal/weapons/catalyst/widsith"
	_ "github.com/genshinsim/gcsim/internal/weapons/catalyst/wine"
	_ "github.com/genshinsim/gcsim/internal/weapons/claymore/akuoumaru"
	_ "github.com/genshinsim/gcsim/internal/weapons/claymore/bell"
	_ "github.com/genshinsim/gcsim/internal/weapons/claymore/blackcliff"
	_ "github.com/genshinsim/gcsim/internal/weapons/claymore/favonius"
	_ "github.com/genshinsim/gcsim/internal/weapons/claymore/nagamasa"
	_ "github.com/genshinsim/gcsim/internal/weapons/claymore/pines"
	_ "github.com/genshinsim/gcsim/internal/weapons/claymore/prototype"
	_ "github.com/genshinsim/gcsim/internal/weapons/claymore/rainslasher"
	_ "github.com/genshinsim/gcsim/internal/weapons/claymore/redhorn"
	_ "github.com/genshinsim/gcsim/internal/weapons/claymore/royal"
	_ "github.com/genshinsim/gcsim/internal/weapons/claymore/sacrifical"
	_ "github.com/genshinsim/gcsim/internal/weapons/claymore/sealord"
	_ "github.com/genshinsim/gcsim/internal/weapons/claymore/skyrider"
	_ "github.com/genshinsim/gcsim/internal/weapons/claymore/skyward"
	_ "github.com/genshinsim/gcsim/internal/weapons/claymore/spine"
	_ "github.com/genshinsim/gcsim/internal/weapons/claymore/starsilver"
	_ "github.com/genshinsim/gcsim/internal/weapons/claymore/unforged"
	_ "github.com/genshinsim/gcsim/internal/weapons/claymore/waster"
	_ "github.com/genshinsim/gcsim/internal/weapons/claymore/whiteblind"
	_ "github.com/genshinsim/gcsim/internal/weapons/claymore/wolf"
	_ "github.com/genshinsim/gcsim/internal/weapons/spear/blackcliff"
	_ "github.com/genshinsim/gcsim/internal/weapons/spear/blacktassel"
	_ "github.com/genshinsim/gcsim/internal/weapons/spear/calamity"
	_ "github.com/genshinsim/gcsim/internal/weapons/spear/catch"
	_ "github.com/genshinsim/gcsim/internal/weapons/spear/crescent"
	_ "github.com/genshinsim/gcsim/internal/weapons/spear/deathmatch"
	_ "github.com/genshinsim/gcsim/internal/weapons/spear/dragonbane"
	_ "github.com/genshinsim/gcsim/internal/weapons/spear/dragonspine"
	_ "github.com/genshinsim/gcsim/internal/weapons/spear/favonius"
	_ "github.com/genshinsim/gcsim/internal/weapons/spear/grasscutter"
	_ "github.com/genshinsim/gcsim/internal/weapons/spear/homa"
	_ "github.com/genshinsim/gcsim/internal/weapons/spear/ironpoint"
	_ "github.com/genshinsim/gcsim/internal/weapons/spear/kitain"
	_ "github.com/genshinsim/gcsim/internal/weapons/spear/lithic"
	_ "github.com/genshinsim/gcsim/internal/weapons/spear/primordial"
	_ "github.com/genshinsim/gcsim/internal/weapons/spear/prototype"
	_ "github.com/genshinsim/gcsim/internal/weapons/spear/royal"
	_ "github.com/genshinsim/gcsim/internal/weapons/spear/skyward"
	_ "github.com/genshinsim/gcsim/internal/weapons/spear/vortex"
	_ "github.com/genshinsim/gcsim/internal/weapons/spear/wavebreaker"
	_ "github.com/genshinsim/gcsim/internal/weapons/spear/whitetassel"
	_ "github.com/genshinsim/gcsim/internal/weapons/sword/alley"
	_ "github.com/genshinsim/gcsim/internal/weapons/sword/amenoma"
	_ "github.com/genshinsim/gcsim/internal/weapons/sword/aquila"
	_ "github.com/genshinsim/gcsim/internal/weapons/sword/black"
	_ "github.com/genshinsim/gcsim/internal/weapons/sword/blackcliff"
	_ "github.com/genshinsim/gcsim/internal/weapons/sword/cinnabar"
	_ "github.com/genshinsim/gcsim/internal/weapons/sword/coolsteel"
	_ "github.com/genshinsim/gcsim/internal/weapons/sword/darkironsword"
	_ "github.com/genshinsim/gcsim/internal/weapons/sword/dullblade"
	_ "github.com/genshinsim/gcsim/internal/weapons/sword/favonius"
	_ "github.com/genshinsim/gcsim/internal/weapons/sword/festering"
	_ "github.com/genshinsim/gcsim/internal/weapons/sword/filletblade"
	_ "github.com/genshinsim/gcsim/internal/weapons/sword/flute"
	_ "github.com/genshinsim/gcsim/internal/weapons/sword/freedom"
	_ "github.com/genshinsim/gcsim/internal/weapons/sword/haran"
	_ "github.com/genshinsim/gcsim/internal/weapons/sword/harbinger"
	_ "github.com/genshinsim/gcsim/internal/weapons/sword/ironsting"
	_ "github.com/genshinsim/gcsim/internal/weapons/sword/lion"
	_ "github.com/genshinsim/gcsim/internal/weapons/sword/lithic"
	_ "github.com/genshinsim/gcsim/internal/weapons/sword/mistsplitter"
	_ "github.com/genshinsim/gcsim/internal/weapons/sword/primordial"
	_ "github.com/genshinsim/gcsim/internal/weapons/sword/prototype"
	_ "github.com/genshinsim/gcsim/internal/weapons/sword/royal"
	_ "github.com/genshinsim/gcsim/internal/weapons/sword/sacrifical"
	_ "github.com/genshinsim/gcsim/internal/weapons/sword/silversword"
	_ "github.com/genshinsim/gcsim/internal/weapons/sword/skyrider"
	_ "github.com/genshinsim/gcsim/internal/weapons/sword/skyward"
	_ "github.com/genshinsim/gcsim/internal/weapons/sword/summit"
	_ "github.com/genshinsim/gcsim/internal/weapons/sword/swordofdescension"
	_ "github.com/genshinsim/gcsim/internal/weapons/sword/travelershandysword"
)
