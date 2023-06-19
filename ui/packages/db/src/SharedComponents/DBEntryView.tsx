import { db, model } from "@gcsim/types";
import { Long } from "protobufjs";
import { ReactI18NextChild, useTranslation } from "react-i18next";
import DBEntryActions from "./DBEntryViewComponents/DBEntryActions";
import { DBEntryPortrait } from "./DBEntryViewComponents/DBEntryPortrait";

//displays one database entry
export default function DBEntryView({ dbEntry }: { dbEntry: db.IEntry }) {
  // const t = useTranslation();
  const team = dbEntry.summary?.team ?? [];
  if (team.length < 4) {
    const diff = 4 - team.length;
    for (let i = 0; i < diff; i++) {
      team.push({} as model.ICharacter);
    }
  }

  return (
    <>
      <div className="hidden lg:flex  flex-row bg-slate-800  p-4 gap-4 w-full max-w-7xl">
        <div className="flex gap-2 flex-row min-w-fit ">
          {team &&
            team.map((char, index) => {
              return <DBEntryPortrait {...char} key={index.toString()} />;
            })}
        </div>
        <div className="flex flex-col grow ">
          <div className="max-w-2xl">
            <div className="flex flex-col ">
              {/* <DBEntryTags tags={dbEntry.accepted_tags} /> */}
              <span className="  overflow-hidden">{dbEntry?.description}</span>
            </div>

            <DBEntryDetails
              {...dbEntry.summary}
              run_date={dbEntry.last_update}
            />
          </div>
        </div>
        <DBEntryActions share_key={dbEntry.share_key} id={dbEntry.id} />
      </div>
      <div className="lg:hidden flex flex-col items-center  bg-slate-700 max-w-xs p-5 border  gap-4 ">
        <div className="grid grid-cols-2 grid-row-2  gap-4">
          {team &&
            team.map((char, index) => {
              return <DBEntryPortrait {...char} key={index.toString()} />;
            })}
        </div>
        <DBEntryActions share_key={dbEntry.share_key} id={dbEntry.id} />
      </div>
    </>
  );
}

function DBEntryDetails({
  target_count,
  mean_dps_per_target,
  mode,
  sim_duration,
  total_damage,
  run_date,
}: // total_damage,
// description,
NonNullable<db.IEntry["summary"]> & {
  run_date?: number | Long | null;
}) {
  const { t: translate } = useTranslation();

  const t = (key: string) => translate(key) as ReactI18NextChild; // idk why this is needed
  let date = t("db.unknown");
  if (run_date && typeof run_date === "number") {
    date = new Date(run_date).toLocaleDateString();
  }
  return (
    <table className="bp4-html-table  ">
      <thead>
        <tr className="">
          <th className="">
            <div>{t("db.simMode") as ReactI18NextChild}</div>
          </th>
          <th className="">{t("db.targetCount")}</th>
          <th className="">{t("db.dpsPerTarget")}</th>
          <th className="">{t("db.avgTotalDmg")}</th>
          <th className="">{t("db.avgSimTime")}</th>
          <th className="">{t("db.runDate")}</th>
        </tr>
      </thead>
      <tbody>
        <tr className=" text-xs ">
          <td className="">{mode ? t("db.ttk") : t("db.duration")}</td>
          <td className="">{target_count}</td>
          <td className="">{mean_dps_per_target?.toPrecision(8)}</td>
          <td className="">{total_damage?.mean?.toPrecision(8)}</td>
          <td className="">
            {sim_duration?.mean
              ? `${sim_duration.mean.toPrecision(3)}s`
              : t("db.unknown")}
          </td>
          <td className="">{date}</td>
        </tr>
      </tbody>
    </table>
  );
}
