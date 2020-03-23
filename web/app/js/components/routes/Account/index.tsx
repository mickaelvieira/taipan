import React, { useContext } from "react";
import { makeStyles } from "@material-ui/core/styles";
import { RouteAccountProps } from "../../../types/routes";
import { UserContext } from "../../context";
import Grid from "../../ui/Grid";
import Prolile from "./Profile";
import Emails from "./Emails";
import Password from "./Password";
import ScrollToTop from "../../ui/ScrollToTop";

const useStyles = makeStyles(() => ({
  content: {
    paddingLeft: 12,
    paddingRight: 12,
  },
}));

export default function Account(_: RouteAccountProps): JSX.Element | null {
  const classes = useStyles();
  const user = useContext(UserContext);
  if (!user) {
    return null;
  }

  return (
    <ScrollToTop>
      <Grid className={classes.content}>
        <Prolile user={user} />
        <Emails user={user} />
        <Password />
      </Grid>
    </ScrollToTop>
  );
}
