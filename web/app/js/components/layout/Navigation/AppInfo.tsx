import React, { useContext } from "react";
import { makeStyles } from "@material-ui/core/styles";
import Button from "@material-ui/core/Button";
import { AppContext } from "../../context";
import { getAppInfoFromEnv } from "../../../helpers/app";

const useStyles = makeStyles(({ breakpoints, palette }) => ({
  root: {
    display: "flex",
    flexDirection: "column"
  },
  appInfo: {
    lineHeight: 1.33,
    color: palette.grey[600],
    textAlign: "center",
    paddingTop: "1.2rem",
    margin: 0,
    [breakpoints.up("md")]: {
      color: palette.grey[100]
    }
  },
  reload: {
    color: palette.primary.main,
    textDecoration: "underline"
  }
}));

export default function AppInfo(): JSX.Element {
  const classes = useStyles();
  const info = useContext(AppContext);
  const [name, version] = getAppInfoFromEnv();
  const shouldReload = info && (name !== info.name || version !== info.version);

  return (
    <div className={classes.root}>
      <p className={classes.appInfo}>
        {name} v{version}
      </p>
      {shouldReload && (
        <Button
          className={classes.reload}
          onClick={() => window.location.reload()}
        >
          Reload
        </Button>
      )}
    </div>
  );
}
