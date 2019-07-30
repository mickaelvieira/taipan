import React, { useEffect } from "react";
import { useTheme } from "@material-ui/core/styles";
import useMediaQuery from "@material-ui/core/useMediaQuery";
import { withRouter } from "react-router";
import { makeStyles } from "@material-ui/core/styles";
import AppBar from "@material-ui/core/AppBar";
import Toolbar from "@material-ui/core/Toolbar";
import IconButton from "@material-ui/core/IconButton";
import MenuIcon from "@material-ui/icons/Menu";
import { RoutesProps } from "../../../types/routes";
import Typography from "@material-ui/core/Typography";
import { SIDEBAR_WIDTH } from "../../../constant/sidebar";
import { getSectionTitle, getPageTitle } from "../helpers/navigation";
import Search from "./Search";

const useStyles = makeStyles(({ breakpoints }) => ({
  appBar: {
    marginLeft: SIDEBAR_WIDTH,
    [breakpoints.up("md")]: {
      width: `calc(100% - ${SIDEBAR_WIDTH}px)`
    }
  },
  toolbar: {
    display: "flex",
    alignItems: "center",
    justifyContent: "space-between"
  },
  title: {
    flexGrow: 1,
    textAlign: "center"
  },
  menuButton: {
    marginLeft: -12
  },
  container: {
    display: "flex",
    flexDirection: "column",
    overflow: "hidden"
  }
}));

interface Props extends RoutesProps {
  toggleDrawer: (status: boolean) => void;
}

export default withRouter(function Header({
  toggleDrawer,
  match
}: Props): JSX.Element {
  const classes = useStyles();
  const theme = useTheme();
  const md = useMediaQuery(theme.breakpoints.up("md"));
  const pageTitle = getPageTitle(window.location.pathname);

  // @TODO add search terms to the document title when looking up documents or articles
  useEffect(() => {
    document.title = pageTitle;
  }, [pageTitle]);

  return (
    <AppBar position="fixed" className={classes.appBar}>
      <Toolbar className={classes.toolbar}>
        {!md && (
          <>
            <IconButton
              className={classes.menuButton}
              color="inherit"
              aria-label="Menu"
              onClick={() => toggleDrawer(true)}
            >
              <MenuIcon />
            </IconButton>
          </>
        )}
        {!md && (
          <div className={classes.container}>
            <Typography component="h6" variant="h5" className={classes.title}>
              {getSectionTitle(match.path)}
            </Typography>
          </div>
        )}
        {md && <Search />}
      </Toolbar>
    </AppBar>
  );
});
