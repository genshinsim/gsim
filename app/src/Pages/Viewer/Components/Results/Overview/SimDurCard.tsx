import { useTranslation } from "react-i18next";
import { SimResults } from "~src/Pages/Viewer/SimResults";
import SummaryCard from "../SummaryCard";

export default ({ data, color }: { data: SimResults | null, color: string }) => {
  const { i18n } = useTranslation();
  const fmt = (val?: number) => val?.toLocaleString(i18n.language, { maximumFractionDigits: 2 });
  const duration = data?.statistics?.duration;

  return (
    <SummaryCard
        key="duration"
        color={color}
        title="Sim Duration"
        label="s"
        value={fmt(duration?.mean)}
        auxStats={[
          { title: "min", value: fmt(duration?.min) },
          { title: "max", value: fmt(duration?.max) },
          { title: "std", value: fmt(duration?.sd) },
          { title: "p25", value: fmt(duration?.q1) },
          { title: "p50", value: fmt(duration?.q2) },
          { title: "p75", value: fmt(duration?.q3) },
        ]}
        tooltip="help"
        drawerTitle="Sim Duration Statistics">
      <div></div>
    </SummaryCard>
  );
};