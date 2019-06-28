import React, { useContext } from "react";
import { UserContext } from "../../context";
import { makeStyles } from "@material-ui/core/styles";

const useStyles = makeStyles(({ palette }) => ({
  userInfo: {
    fontSize: "1.2rem",
    fontWeight: 500,
    lineHeight: 1.33,
    letterSpacing: "0em",
    color: palette.grey[900],
    textAlign: "center",
    margin: 0,
    padding: "1.2rem 0",
    backgroundColor: palette.grey[100]
  }
}));

export default function UserInfo(): JSX.Element | null {
  const classes = useStyles();
  const user = useContext(UserContext);

  return !user ? null : (
    <p className={classes.userInfo}>
      {user.firstname} {user.lastname}
    </p>
  );
}
