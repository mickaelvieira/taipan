import React, { useState } from "react";
import CloseIcon from "@material-ui/icons/Close";
import Tabs from "@material-ui/core/Tabs";
import Tab from "@material-ui/core/Tab";
import Panel from "../../../ui/Panel";
import TabPanel from "./TabPanel";
import Info from "./Info";
import Logs from "./Logs";

interface Props {
  url: string;
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
      <Tabs value={tab} onChange={(_, value) => setTab(value)}>
        <Tab label="Information" value={0}></Tab>
        <Tab label="Logs" value={1}></Tab>
      </Tabs>
      <TabPanel isVisible={tab === 0}>
        <Info url={url} />
      </TabPanel>
      <TabPanel isVisible={tab === 1}>
        <Logs url={url} />
      </TabPanel>
    </Panel>
  );
}
