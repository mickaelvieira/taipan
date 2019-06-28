import React, { useContext } from "react";
import { makeStyles } from "@material-ui/core/styles";
import Button from "@material-ui/core/Button";
import { AppContext } from "../../context";

const useStyles = makeStyles(({ palette }) => ({
  root: {
    display: "flex",
    flexDirection: "column"
  },
  appInfo: {
    lineHeight: 1.33,
    color: palette.grey[500],
    textAlign: "center",
    paddingTop: "1.2rem",
    margin: 0
  },
  reload: {
    color: palette.primary.main,
    textDecoration: "underline"
  }
}));

function getInfoFromEnv(): [string, string] {
  return [process.env.APP_NAME || "", process.env.APP_VERSION || ""];
}

export default function AppInfo(): JSX.Element {
  const classes = useStyles();
  const info = useContext(AppContext);
  const [name, version] = getInfoFromEnv();
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
