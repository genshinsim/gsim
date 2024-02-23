import { FormGroup, InputGroup } from "@blueprintjs/core";
import { ExecutorSupplier, ServerExecutor } from "@gcsim/executors";
import { UI } from "@gcsim/ui";
import React, { ReactNode, useRef } from "react";
import { useTranslation } from "react-i18next";

let exec: ServerExecutor | undefined;

const ServerMode = ({ children }: { children: ReactNode }) => {
  const { t } = useTranslation();
  const [url, setURL] = React.useState("http://127.0.0.1:54321");

  React.useEffect(() => {
    if (exec != null) {
      exec.set_url(url);
    }
  }, [url]);

  const supplier = useRef<ExecutorSupplier<ServerExecutor>>(() => {
    if (exec == null) {
      exec = new ServerExecutor(url);
    }
    return exec;
  });

  return (
    <UI
      exec={supplier.current}
      gitCommit={import.meta.env.VITE_GIT_COMMIT_HASH}
      mode={import.meta.env.MODE}
    >
      <FormGroup className="!m-0" label={t<string>("simple.workers")}>
        {children}
        <FormGroup
          helperText="default: http://127.0.0.1.54321"
          label="Server URL"
          labelFor="text-input"
          labelInfo="(required)"
        >
          <InputGroup
            id="text-input"
            value={url}
            onChange={(e) => {
              setURL(e.target.value);
            }}
            fill
          />
        </FormGroup>
      </FormGroup>
    </UI>
  );
};

export default ServerMode;
