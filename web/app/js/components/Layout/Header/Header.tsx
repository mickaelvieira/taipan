import React from "react";
import { makeStyles } from "@material-ui/core/styles";
import { fade } from "@material-ui/core/styles/colorManipulator";
import AppBar from "@material-ui/core/AppBar";
import Toolbar from "@material-ui/core/Toolbar";
import IconButton from "@material-ui/core/IconButton";
import MenuIcon from "@material-ui/icons/Menu";
import InputBase from "@material-ui/core/InputBase";
import SearchIcon from "@material-ui/icons/Search";

const useStyles = makeStyles(
  ({ shape, palette, spacing, breakpoints, transitions }) => ({
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

interface Props {
  toggleDrawer: (status: boolean) => void;
}

export default function Header({ toggleDrawer }: Props) {
  const classes = useStyles();
  return (
    <>
      <AppBar position="fixed">
        <Toolbar>
          <IconButton
            className={classes.menuButton}
            color="inherit"
            aria-label="Menu"
            onClick={() => toggleDrawer(true)}
          >
            <MenuIcon />
          </IconButton>

          <div className={classes.search}>
            <div className={classes.searchIcon}>
              <SearchIcon />
            </div>
            <InputBase
              placeholder="Search..."
              classes={{
                root: classes.inputRoot,
                input: classes.inputInput
              }}
            />
          </div>
        </Toolbar>
      </AppBar>
    </>
  );
}
