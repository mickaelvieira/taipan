import React, { useState } from "react";
import { makeStyles } from "@material-ui/core/styles";
import Tabs from "@material-ui/core/Tabs";
import Tab from "@material-ui/core/Tab";
import IconButton from "@material-ui/core/IconButton";
import CloseIcon from "@material-ui/icons/Close";
import Typography from "@material-ui/core/Typography";
import Panel from "../../../ui/Panel";
import TabPanel from "./TabPanel";

import Info from "./Info";
import Logs from "./Logs";

const useStyles = makeStyles(({ palette }) => ({
  header: {
    display: "flex",
    flexDirection: "row",
    justifyContent: "start",
    margin: 0,
    padding: 0,
    backgroundColor: palette.grey[200]
  },
  title: {
    paddingTop: 12,
    paddingBottom: 12
  }
}));

interface Props {
  url: string;
  isOpen: boolean;
  close: () => void;
}

export default function EditSource({ url, isOpen, close }: Props): JSX.Element {
  const classes = useStyles();
  const [tab, setTab] = useState(0);

  return (
    <Panel isOpen={isOpen} close={close}>
      <header className={classes.header}>
        <IconButton onClick={() => close()}>
          <CloseIcon />
        </IconButton>
        <Typography component="h5" variant="h6" className={classes.title}>
          Edit web syndication source
        </Typography>
      </header>
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
