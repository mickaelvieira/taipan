import React, { useState } from "react";
import { makeStyles } from "@material-ui/core/styles";
import FormControl from "@material-ui/core/FormControl";
import FormGroup from "@material-ui/core/FormGroup";
import FormControlLabel from "@material-ui/core/FormControlLabel";
import Switch from "@material-ui/core/Switch";
import Layout from "../../Layout";
import SyndicationList from "./List";

const useStyles = makeStyles({
  switch: {
    alignSelf: "flex-start"
  }
});

export default function Syndication(): JSX.Element {
  const classes = useStyles();
  const [pausedSourcesOnly, showPausedSourcesOnly] = useState(false);

  return (
    <Layout>
      <FormControl component="fieldset" className={classes.switch}>
        <FormGroup>
          <FormControlLabel
            control={
              <Switch
                onChange={() => showPausedSourcesOnly(!pausedSourcesOnly)}
                checked={pausedSourcesOnly}
              />
            }
            label="Paused sources"
          />
        </FormGroup>
      </FormControl>
      <SyndicationList showPausedSources={pausedSourcesOnly} />
    </Layout>
  );
}
