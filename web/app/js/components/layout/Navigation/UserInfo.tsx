import React, { useContext } from "react";
import { UserContext } from "../../context";
import { makeStyles } from "@material-ui/core/styles";
import Avatar from "@material-ui/core/Avatar";

const useStyles = makeStyles(({ palette, breakpoints }) => ({
  userInfo: {
    display: "flex",
    flexDirection: "row",
    justifyContent: "space-evenly",
    alignItems: "center",
    padding: "0 8px",
    color: palette.grey[900],
    [breakpoints.up("md")]: {
      backgroundColor: "#1d1d1d",
      color: palette.primary.main
    }
  },
  name: {
    fontSize: "1.2rem",
    fontWeight: 500,
    lineHeight: 1.33,
    letterSpacing: "0em"
  },
  avatar: {
    width: 35,
    height: 35
  }
}));

export default function UserInfo(): JSX.Element | null {
  const classes = useStyles();
  const user = useContext(UserContext);

  return !user ? null : (
    <div className={classes.userInfo}>
      {user.image && (
        <Avatar
          alt={`${user.firstname} ${user.lastname}`}
          src={`${user.image.url}`}
          className={classes.avatar}
        />
      )}
      <p className={classes.name}>
        {user.firstname} {user.lastname}
      </p>
    </div>
  );
}
