import React, { useState, useEffect } from "react";
import { useTheme } from "@material-ui/core/styles";
import useMediaQuery from "@material-ui/core/useMediaQuery";
import { withRouter } from "react-router";
import { makeStyles } from "@material-ui/core/styles";
// import { fade } from "@material-ui/core/styles/colorManipulator";
import AppBar from "@material-ui/core/AppBar";
// import Hidden from "@material-ui/core/Hidden";
import Toolbar from "@material-ui/core/Toolbar";
import ButtonBase from "@material-ui/core/ButtonBase";
import IconButton from "@material-ui/core/IconButton";
import MenuIcon from "@material-ui/icons/Menu";
// import InputBase from "@material-ui/core/InputBase";
import SearchIcon from "@material-ui/icons/Search";
import CloseIcon from "@material-ui/icons/Close";
import { RoutesProps } from "../../../types/routes";
import Typography from "@material-ui/core/Typography";
// import Fade from "@material-ui/core/Fade";
// import Slide from "@material-ui/core/Slide";
import { SIDEBAR_WIDTH } from "../../../constant/sidebar";
import { getSectionTitle, getPageTitle } from "../helpers/navigation";

const useStyles = makeStyles(
  ({ shape, spacing, breakpoints, transitions }) => ({
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
    },
    search: {
      flexGrow: 1,
      borderRadius: shape.borderRadius,
      display: "flex",
      alignItems: "center",
      margin: "0 16px"
    },
    searchLabel: {
      flexGrow: 1
    },
    searchIcon: {
      display: "flex",
      alignItems: "center",
      justifyContent: "center"
    },
    inputRoot: {
      width: "100%"
    },
    inputInput: {
      padding: spacing(1),
      transition: transitions.create("width"),
      width: "100%"
      // [breakpoints.up("md")]: {
      //   width: 200
      // }
    }
  })
);

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
  const [isOpen, setSearchIsOpen] = useState(false);
  const pageTitle = getPageTitle(match.path);

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
          <>
            <div className={classes.container}>
              {/* <Slide
              in={!isOpen}
              direction="down"
              timeout={{
                enter: 200,
                exit: 200
              }}
            > */}
              <Typography component="h6" variant="h5" className={classes.title}>
                {getSectionTitle(match.path)}
              </Typography>
              {/* </Slide> */}
              {/* <Slide
              in={isOpen}
              direction="up"
              timeout={{
                enter: 200,
                exit: 200
              }}
            > */}
              {/* <div className={classes.search}>
                <div className={classes.searchIcon}>
                  <SearchIcon />
                </div>
                <label htmlFor="search-field" className={classes.searchLabel}>
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
              {/* </Slide> */}
            </div>
            <div>
              <ButtonBase onClick={() => setSearchIsOpen(!isOpen)}>
                {isOpen ? <CloseIcon /> : <SearchIcon />}
              </ButtonBase>
            </div>
          </>
        )}
      </Toolbar>
    </AppBar>
  );
});
