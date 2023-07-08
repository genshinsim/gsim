import CharDataGen from '@ui/Data/char_data.generated.json';
import WeaponDataGen from '@ui/Data/weapon_data.generated.json';
import TextMap from './GenshinData/EnkaTextMapEN.json';

import {
  GOODArtifact,
  GOODArtifactSetKey,
  GOODCharacter,
  GOODSlotKey,
  GOODStatKey,
  GOODWeapon,
  IGOOD,
  ISubstat,
} from '../GOOD/GOODTypes';
import {
  EnkaData,
  FightProp,
  GenshinItemReliquary,
  GenshinItemWeapon,
  ReliquaryEquipType,
} from './EnkaTypes';

const characterMapV2 = Object.values(CharDataGen.data).reduce(
  (acc, val) => {
    acc[val.id] = val;
    return acc;
  },
  {} as {
    [id: number]: {
      id: number;
      key: string;
      rarity: string;
      body: string;
      region: string;
      element: string;
      weapon_class: string;
      icon_name: string;
      skill_details: {
        skill: number;
        burst: number;
        attack: number;
        burst_energy_cost: number;
      };
    };
  }
);

const weaponMapV2 = Object.values(WeaponDataGen.data).reduce(
  (acc, val) => {
    acc[val.id] = val;
    return acc;
  },
  {} as {
    [id: number]: {
      id: number;
      key: string;
      rarity: number;
      weapon_class: string;
      image_name: string;
    };
  }
);

export default function EnkaToGOOD(enkaData: EnkaData): IGOOD {
  const characters: GOODCharacter[] = [];
  const artifacts: GOODArtifact[] = [];
  const weapons: GOODWeapon[] = [];

  enkaData.avatarInfoList.forEach(
    ({ avatarId, propMap, talentIdList, skillLevelMap, equipList }) => {
      const characterData = characterMapV2[avatarId];
      if (!characterData || !characterData.key) {
        console.log('Missing/Unimplemented character', avatarId);
        return;
      }

      const character: GOODCharacter = {
        //this is already in srl key but goodtosrl is idempotent so its fine
        key: characterData.key,
        level: parseInt(propMap['4001'].val) ?? 1,
        ascension: parseInt(propMap['1002'].val) ?? 1,
        constellation: talentIdList?.length ?? 0,
        talent: getCharacterTalentV2(
          characterData.skill_details,
          skillLevelMap
        ),
      };
      characters.push(character);

      equipList.forEach((equip) => {
        if (equip.flat.itemType == 'ITEM_WEAPON') {
          const { weapon: enkaWeapon, itemId } = equip as GenshinItemWeapon;
          const weaponData = weaponMapV2[itemId];
          if (!weaponData || !weaponData.key) {
            // goodtosrl handles no weapon characters
            return;
          }
          const weapon: GOODWeapon = {
            key: weaponData.key,
            level: enkaWeapon.level,
            ascension: enkaWeapon.promoteLevel ?? 0,
            refinement: determineWeaponRefinement(enkaWeapon.affixMap),
            location: character.key,
            lock: false,
          };
          weapons.push(weapon);
        } else {
          const { flat, reliquary: enkaReliquary } =
            equip as GenshinItemReliquary;

          const artifact: GOODArtifact = {
            setKey: getGOODKeyFromReliquaryNameTextMapHash(
              flat.setNameTextMapHash
            ),
            level: enkaReliquary.level - 1,
            slotKey: reliquaryTypeToGOODKey(flat.equipType),
            rarity: flat.rankLevel,
            location: character.key,
            lock: false,
            mainStatKey: fightPropToGOODKey(flat.reliquaryMainstat.mainPropId),
            substats: getGOODSubstatsFromReliquarySubstats(
              flat.reliquarySubstats
            ),
          };
          artifacts.push(artifact);
        }
      });
    }
  );

  return {
    format: 'GOOD' as IGOOD['format'],
    version: 2,
    source: 'gcsimFromEnka',
    characters,
    weapons,
    artifacts,
  };
}

interface IENTextMap {
  [key: string]: string;
}
const textMap: IENTextMap = TextMap.en;

function determineWeaponRefinement(affixMap?: { [key: number]: number }) {
  return affixMap
    ? (Object.entries(affixMap)[0] != null
        ? Object.entries(affixMap)[0][1]
        : 0) + 1
    : 1;
}

function getCharacterTalentV2(
  skill_details: {
    skill: number;
    burst: number;
    attack: number;
  },
  skillLevelMap: { [key: number]: number }
) {
  return {
    auto: skillLevelMap[skill_details.attack],
    skill: skillLevelMap[skill_details.skill],
    burst: skillLevelMap[skill_details.burst],
  };
}

function textToGOODKey(string: string) {
  function toTitleCase(str: string) {
    return str.replace(/-/g, ' ').replace(/\w\S*/g, function (txt: string) {
      return txt.charAt(0).toUpperCase() + txt.substring(1).toLowerCase();
    });
  }
  return toTitleCase(string || '').replace(/[^A-Za-z]/g, '');
}

function reliquaryTypeToGOODKey(
  reliquaryType: ReliquaryEquipType
): GOODSlotKey {
  switch (reliquaryType) {
    case ReliquaryEquipType.EQUIP_BRACER:
      return 'flower';
    case ReliquaryEquipType.EQUIP_NECKLACE:
      return 'plume';
    case ReliquaryEquipType.EQUIP_SHOES:
      return 'sands';
    case ReliquaryEquipType.EQUIP_RING:
      return 'goblet';
    case ReliquaryEquipType.EQUIP_DRESS:
      return 'circlet';
  }
}

function fightPropToGOODKey(fightProp: FightProp): GOODStatKey {
  switch (fightProp) {
    case FightProp.FIGHT_PROP_HP:
      return 'hp';
    case FightProp.FIGHT_PROP_HP_PERCENT:
      return 'hp_';
    case FightProp.FIGHT_PROP_ATTACK:
      return 'atk';
    case FightProp.FIGHT_PROP_ATTACK_PERCENT:
      return 'atk_';
    case FightProp.FIGHT_PROP_DEFENSE:
      return 'def';
    case FightProp.FIGHT_PROP_DEFENSE_PERCENT:
      return 'def_';
    case FightProp.FIGHT_PROP_CHARGE_EFFICIENCY:
      return 'enerRech_';
    case FightProp.FIGHT_PROP_ELEMENT_MASTERY:
      return 'eleMas';
    case FightProp.FIGHT_PROP_CRITICAL:
      return 'critRate_';
    case FightProp.FIGHT_PROP_CRITICAL_HURT:
      return 'critDMG_';
    case FightProp.FIGHT_PROP_HEAL_ADD:
      return 'heal_';
    case FightProp.FIGHT_PROP_FIRE_ADD_HURT:
      return 'pyro_dmg_';
    case FightProp.FIGHT_PROP_ELEC_ADD_HURT:
      return 'electro_dmg_';
    case FightProp.FIGHT_PROP_ICE_ADD_HURT:
      return 'cryo_dmg_';
    case FightProp.FIGHT_PROP_WATER_ADD_HURT:
      return 'hydro_dmg_';
    case FightProp.FIGHT_PROP_WIND_ADD_HURT:
      return 'anemo_dmg_';
    case FightProp.FIGHT_PROP_ROCK_ADD_HURT:
      return 'geo_dmg_';
    case FightProp.FIGHT_PROP_GRASS_ADD_HURT:
      return 'dendro_dmg_';
    case FightProp.FIGHT_PROP_PHYSICAL_ADD_HURT:
      return 'physical_dmg_';
    default:
      return '';
  }
}

function getGOODKeyFromReliquaryNameTextMapHash(
  reliquaryNameTextMapHash: string
): GOODArtifactSetKey {
  const res = textToGOODKey(
    textMap[reliquaryNameTextMapHash]
  ) as GOODArtifactSetKey;
  return res;
}

function getGOODSubstatsFromReliquarySubstats(
  reliquarySubstats: {
    appendPropId: FightProp;
    statValue: number;
  }[]
): ISubstat[] {
  if (reliquarySubstats.length == 0 || reliquarySubstats.length > 4) {
    return [];
  }
  const GOODSubstats: ISubstat[] = [];
  for (const reliquarySubstat of reliquarySubstats) {
    GOODSubstats.push({
      key: fightPropToGOODKey(reliquarySubstat.appendPropId),
      value: reliquarySubstat.statValue,
    });
  }
  return GOODSubstats;
}
