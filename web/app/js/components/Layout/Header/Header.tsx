import React from "react";
import {
  withStyles,
  WithStyles,
  createStyles,
  Theme
} from "@material-ui/core/styles";
import { fade } from "@material-ui/core/styles/colorManipulator";
import AppBar from "@material-ui/core/AppBar";
import Toolbar from "@material-ui/core/Toolbar";
import IconButton from "@material-ui/core/IconButton";
import MenuIcon from "@material-ui/icons/Menu";
import InputBase from "@material-ui/core/InputBase";
import SearchIcon from "@material-ui/icons/Search";

const styles = ({ breakpoints, palette, spacing, transitions, shape }: Theme) =>
  createStyles({
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
      marginRight: spacing.unit * 2,
      marginLeft: 0,
      width: "100%",
      [breakpoints.up("sm")]: {
        marginLeft: spacing.unit * 3,
        width: "auto"
      }
    },
    searchIcon: {
      width: spacing.unit * 9,
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
      paddingTop: spacing.unit,
      paddingRight: spacing.unit,
      paddingBottom: spacing.unit,
      paddingLeft: spacing.unit * 10,
      transition: transitions.create("width"),
      width: "100%",
      [breakpoints.up("md")]: {
        width: 200
      }
    }
  });

interface Props extends WithStyles<typeof styles> {
  toggleDrawer: (status: boolean) => void;
}

export default withStyles(styles)(function Header({
  toggleDrawer,
  classes
}: Props) {
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
});
