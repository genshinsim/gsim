import { Button, Classes, Intent, Position, ProgressBar, Toaster } from "@blueprintjs/core";
import classNames from "classnames";
import { MutableRefObject, RefObject, useCallback, useEffect, useRef } from "react";
import { pool } from "~src/Pages/Sim";
import { ResultSource } from "..";

type Props = {
  src: ResultSource;
  error: string | null;
  current?: number;
  total?: number;
};

// TODO: Add translations + number format
export default ({ src, error, current, total }: Props) => {
  const loadingToast = useRef<Toaster>(null);
  const key = useRef<string | undefined>(undefined);

  useEffect(() => {
    if (error != null) {
      loadingToast.current?.clear();
      return;
    }

    if (current == undefined || total == undefined) {
      key.current = loadingToast.current?.show({
        message: "Loading...",
        icon: "refresh",
        intent: Intent.PRIMARY,
        isCloseButtonShown: false,
        timeout: 0,
      }, key.current);
      return;
    }

    if (current >= total && src == ResultSource.Loaded) {
      key.current = loadingToast.current?.show({
        message: "Loaded " + current + " iterations!",
        icon: "tick",
        intent: Intent.SUCCESS,
        isCloseButtonShown: true,
        timeout: 2000,
      }, key.current);
      return;
    }

    if (key.current == null) {
      return;
    }

    key.current = loadingToast.current?.show({
      message: (
          <ProgressToast
              current={current}
              total={total}
              toastKey={key}
              loadingToast={loadingToast} />
      ),
      className: "w-full !max-w-2xl",
      intent: Intent.NONE,
      isCloseButtonShown: current >= total,
      timeout: current < total ? 0 : 2000
    }, key.current);
  }, [current, total, src, error]);

  return <Toaster ref={loadingToast} position={Position.TOP} />;
};

const ProgressToast = ({ current, total, toastKey, loadingToast }: {
    current: number,
    total: number,
    toastKey: MutableRefObject<string | undefined>,
    loadingToast: RefObject<Toaster> }) => {
  
  const cancel = useCallback(() => pool.cancel(), []);
  useEffect(() => {
    return () => cancel();
  }, [cancel]);

  const val = current / total;
  return (
    <div className="flex flex-row items-center justify-between gap-2">
      <div className="min-w-fit">Running ({current}/{total})</div>
      <ProgressBar
          className={classNames("basis-1/2 flex-auto sm:min-w-", {
            [Classes.PROGRESS_NO_STRIPES]: val >= 1,
          })}
          intent={val < 1 ? Intent.PRIMARY : Intent.SUCCESS}
          value={val}/>
      {action(val, toastKey, loadingToast, cancel)}
    </div>
  );
};

function action(
    val: number,
    key: MutableRefObject<string | undefined>,
    loadingToast: RefObject<Toaster>,
    cancel: () => void) {
  if (val >= 1) {
    return null;
  }
  return (
    <Button
        className="!min-w-fit"
        text="Cancel"
        intent={Intent.DANGER}
        onClick={() => {
          cancel();
          loadingToast.current?.clear();
          key.current = undefined;
        }} />
  );
}