import React, { useContext } from "react";
import { makeStyles } from "@material-ui/core/styles";
import { RouteAccountProps } from "../../../types/routes";
import { UserContext } from "../../context";
import Grid from "../../ui/Grid";
import Prolile from "./Profile";
import Emails from "./Emails";

const useStyles = makeStyles(() => ({
  content: {
    paddingLeft: 12,
    paddingRight: 12
  }
}));

export default function Account(_: RouteAccountProps): JSX.Element | null {
  const classes = useStyles();
  const user = useContext(UserContext);
  if (!user) {
    return null;
  }

  return (
    <Grid className={classes.content}>
      <Prolile user={user} />
      <Emails user={user} />
    </Grid>
  );
}
