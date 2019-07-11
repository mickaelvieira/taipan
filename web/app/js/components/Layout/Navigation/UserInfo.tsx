import React, { useContext } from "react";
import { UserContext } from "../../context";
import { makeStyles } from "@material-ui/core/styles";
import Avatar from "@material-ui/core/Avatar";

const useStyles = makeStyles(({ palette }) => ({
  userInfo: {
    display: "flex",
    flexDirection: "row",
    justifyContent: "center",
    alignItems: "center",
    margin: 0,
    padding: 0,
    backgroundColor: palette.grey[100]
  },
  name: {
    fontSize: "1.2rem",
    fontWeight: 500,
    lineHeight: 1.33,
    letterSpacing: "0em",
    color: palette.grey[900]
  },
  avatar: {
    width: 45,
    height: 45,
    marginRight: 12
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
          src={user.image.url}
          className={classes.avatar}
        />
      )}
      <p className={classes.name}>
        {user.firstname} {user.lastname}
      </p>
    </div>
  );
}
