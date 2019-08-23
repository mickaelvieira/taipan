import React, { useState } from "react";
import CloseIcon from "@material-ui/icons/Close";
import Tabs from "@material-ui/core/Tabs";
import Panel from "..";
import Tab from "./Tab";
import TabPanel from "./TabPanel";
import Info from "./Info";
import Tags from "./Tags";
import Logs from "./Logs";

interface Props {
  url: URL | null;
  isOpen: boolean;
  close: () => void;
}

export default function EditSource({
  url,
  isOpen,
  close
}: Props): JSX.Element | null {
  const [tab, setTab] = useState(0);

  return !url ? null : (
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
        <Tab label="Tags" value={1} />
        <Tab label="Logs" value={2} />
      </Tabs>
      <TabPanel value={tab} index={0}>
        <Info url={url} />
      </TabPanel>
      <TabPanel value={tab} index={1}>
        <Tags url={url} />
      </TabPanel>
      <TabPanel value={tab} index={2}>
        <Logs url={url} />
      </TabPanel>
    </Panel>
  );
}
