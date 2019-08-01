import React, { useContext } from "react";
import { makeStyles } from "@material-ui/core/styles";
import { UserContext } from "../../context";
import Grid from "../../ui/Grid";
import Prolile from "./Profile";

const useStyles = makeStyles(() => ({
  content: {
    paddingLeft: 12,
    paddingRight: 12
  }
}));

export default function Account(): JSX.Element | null {
  const classes = useStyles();
  const user = useContext(UserContext);
  if (!user) {
    return null;
  }

  return (
    <Grid className={classes.content}>
      <Prolile user={user} />
    </Grid>
  );
}
