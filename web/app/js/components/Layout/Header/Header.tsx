import React from "react";
import { withRouter } from "react-router";
import { makeStyles } from "@material-ui/core/styles";
import { fade } from "@material-ui/core/styles/colorManipulator";
import AppBar from "@material-ui/core/AppBar";
import Hidden from "@material-ui/core/Hidden";
import Toolbar from "@material-ui/core/Toolbar";
import IconButton from "@material-ui/core/IconButton";
import MenuIcon from "@material-ui/icons/Menu";
import { RouteFeedProps } from "../../../types/routes";
import { Typography } from "@material-ui/core";
import { SIDEBAR_WIDTH } from "../../../constant/sidebar";

const useStyles = makeStyles(
  ({ shape, palette, spacing, breakpoints, transitions }) => ({
    appBar: {
      marginLeft: SIDEBAR_WIDTH,
      [breakpoints.up("md")]: {
        width: `calc(100% - ${SIDEBAR_WIDTH}px)`
      }
    },
    menuButton: {
      marginLeft: -12,
      marginRight: 20
    },
    search: {
      position: "relative",
      borderRadius: shape.borderRadius,
      backgroundColor: fade(palette.common.white, 0.15),
      "&:hover": {
        backgroundColor: fade(palette.common.white, 0.25)
      },
      marginRight: spacing(2),
      marginLeft: 0,
      width: "100%",
      [breakpoints.up("sm")]: {
        marginLeft: spacing(3),
        width: "auto"
      }
    },
    searchIcon: {
      width: spacing(9),
      height: "100%",
      position: "absolute",
      pointerEvents: "none",
      display: "flex",
      alignItems: "center",
      justifyContent: "center"
    },
    inputRoot: {
      color: "inherit",
      width: "100%"
    },
    inputInput: {
      paddingTop: spacing(1),
      paddingRight: spacing(1),
      paddingBottom: spacing(1),
      paddingLeft: spacing(10),
      transition: transitions.create("width"),
      width: "100%",
      [breakpoints.up("md")]: {
        width: 200
      }
    }
  })
);

interface Props extends RouteFeedProps {
  toggleDrawer: (status: boolean) => void;
}

export default withRouter(function Header({
  toggleDrawer,
  match
}: Props): JSX.Element {
  const classes = useStyles();
  let title = "";
  if (match.path === "/") {
    title = "News";
  } else if (match.path === "/reading-list") {
    title = "Reading list";
  } else if (match.path === "/favorites") {
    title = "Favorites";
  }

  return (
    <AppBar position="fixed" className={classes.appBar}>
      <Toolbar>
        <Hidden mdUp>
          <IconButton
            className={classes.menuButton}
            color="inherit"
            aria-label="Menu"
            onClick={() => toggleDrawer(true)}
          >
            <MenuIcon />
          </IconButton>
          <Typography component="h6" variant="h5">
            {title}
          </Typography>
        </Hidden>
        {/* <div className={classes.search}>
            <div className={classes.searchIcon}>
              <SearchIcon />
            </div>
            <label htmlFor="search-field">
              <InputBase
                id="search-field"
                placeholder="Search..."
                classes={{
                  root: classes.inputRoot,
                  input: classes.inputInput
                }}
              />
            </label>
          </div> */}
      </Toolbar>
    </AppBar>
  );
});
