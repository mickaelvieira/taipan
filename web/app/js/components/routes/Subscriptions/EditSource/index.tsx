import React, { useState } from "react";
import CloseIcon from "@material-ui/icons/Close";
import Tabs from "@material-ui/core/Tabs";
import Panel from "../../../ui/Panel";
import Tab from "./Tab";
import TabPanel from "./TabPanel";
import Info from "./Info";
import Logs from "./Logs";

interface Props {
  url: URL;
  isOpen: boolean;
  close: () => void;
}

export default function EditSource({ url, isOpen, close }: Props): JSX.Element {
  const [tab, setTab] = useState(0);

  return (
    <Panel
      BackButton={CloseIcon}
      title="Edit web syndication source"
      isOpen={isOpen}
      prev={close}
    >
      <Tabs
        value={tab}
        onChange={(_, value) => setTab(value)}
        aria-label="Source info"
      >
        <Tab label="Information" value={0} />
        <Tab label="Logs" value={1} />
      </Tabs>
      <TabPanel value={tab} index={0}>
        <Info url={url} />
      </TabPanel>
      <TabPanel value={tab} index={1}>
        <Logs url={url} />
      </TabPanel>
    </Panel>
  );
}
