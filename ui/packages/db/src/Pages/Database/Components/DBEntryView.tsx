import { model } from "@gcsim/types";
import { Long } from "protobufjs";
import { useState, useTransition } from "react";

function useTranslation() {
  return (text: string) => text;
}

//displays one database entry
export default function DBEntryView({ dbEntry }: { dbEntry: model.IDBEntry }) {
  const t = useTranslation();
  return (
    <div className="flex flex-row bg-slate-800 max-w-fit p-4 gap-4 max-h-44">
      <div className="flex gap-2">
        {dbEntry.team &&
          dbEntry.team.map((char, index) => {
            return (
              <DBEntryCharacterPortrait {...char} key={index.toString()} />
            );
          })}
      </div>
      <div className="flex flex-col max-w-md">
        <div className="capitalize text-lg font-semibold">
          {dbEntry?.charNames?.toString().replaceAll(",", ", ")}
        </div>
        <DBEntryTags tags={dbEntry.tags} />
        <span className="  overflow-hidden">{dbEntry?.description}</span>
      </div>

      <DBEntryDetails {...dbEntry} />
      <DBEntryActions />
    </div>
  );
}

function DBEntryCharacterPortrait({ name, sets, weapon }: model.ICharacter) {
  return (
    <div className="bg-slate-700 p-2 flex flex-col">
      {name && (
        <img
          src={"https://gcsim.app/api/assets/avatar/" + name + ".png"}
          alt={name}
          className="w-24 h-24"
        />
      )}
      <div className="flex flex-row ">
        <PortraitWeaponComponent weapon={weapon} />
        <PortraitArtifactsComponent artifactSet={sets} />
      </div>
    </div>
  );
}
function PortraitWeaponComponent({
  weapon,
}: {
  weapon: model.IWeapon | undefined | null;
}) {
  if (!weapon || !weapon.name) {
    return <div className="h-16 w-16">No weapon</div>;
  }
  return (
    <div className=" relative max-h-min flex w-8 h-8">
      <img
        src={"https://gcsim.app/api/assets/weapons/" + weapon.name + ".png"}
        alt={weapon.name}
      />
      <div className=" absolute bottom-0 right-0  text-xs  font-semibold">
        R{weapon?.refine?.toString()}
      </div>
    </div>
  );
}
function PortraitArtifactsComponent({
  artifactSet,
}: {
  artifactSet:
    | {
        [k: string]: number | Long;
      }
    | undefined
    | null;
}) {
  if (!artifactSet) {
    return <div className="h-4 w-4">No artifact</div>;
  }
  return (
    <div className="flex flex-row   ">
      {Object.entries(artifactSet).map(
        ([setName, setCount]) =>
          setCount > 0 && (
            <div className=" relative h-8 w-8 ">
              <img
                src={
                  "https://gcsim.app/api/assets/artifacts/" +
                  setName +
                  "_flower.png"
                }
                alt={setName}
              />
              <div className=" absolute bottom-0 right-0  text-xs  font-semibold">
                {setCount.toString()}
              </div>
            </div>
          )
      )}
    </div>
  );
}

function DBEntryTags({ tags }: { tags: string[] | undefined | null }) {
  const t = useTranslation();
  const [showAll, setShowAll] = useState(false);
  return (
    <div
      className={
        "flex flex-row h-full flex-wrap  relative  max-w-md " +
        (showAll ? " " : " truncate")
      }
    >
      {tags?.map((tag) => (
        <div
          className="bg-slate-700 text-xs font-semibold rounded-full px-2 py-1 mr-2 mt-1 whitespace-nowrap "
          key={tag}
        >
          {t(tag)}
        </div>
      ))}
      <button
        className=" absolute right-0 top-1  text-xs font-semibold     bg-blue-400/30 p-1 mr-1   whitespace-nowrap rounded-sm h-fit   "
        onClick={() => setShowAll(!showAll)}
      >
        {showAll ? "▲" : "▼"}
      </button>
    </div>
  );
}

function DBEntryDetails({
  targetCount,
  meanDpsPerTarget,
  runDate,
}: model.IDBEntry) {
  return (
    <div className="flex flex-col justify-center">
      {targetCount && <div>Target Count: {targetCount}</div>}
      {meanDpsPerTarget && <div>Mean DPS Per Target: {meanDpsPerTarget}</div>}
      {runDate && <div>Run Date: {JSON.stringify(runDate)}</div>}
    </div>
  );
}

function DBEntryActions() {
  const simulation_key = "test"; // TODO: get simulation key from dbEntry
  return (
    <div className="flex flex-col justify-center">
      <a
        href={`https://gcsim.app/v3/viewer/share/${simulation_key}`}
        target="_blank"
        className="text-white bg-slate-600 rounded-md p-2"
        rel="noreferrer"
      >
        Open in Viewer
      </a>
    </div>
  );
}
