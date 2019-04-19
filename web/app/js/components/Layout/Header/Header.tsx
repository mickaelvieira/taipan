import React from "react";
import { withStyles, WithStyles, createStyles } from "@material-ui/core/styles";
import AppBar from "@material-ui/core/AppBar";
import Toolbar from "@material-ui/core/Toolbar";
import IconButton from "@material-ui/core/IconButton";
import MenuIcon from "@material-ui/icons/Menu";

const styles = () =>
  createStyles({
    menuButton: {
      marginLeft: -12,
      marginRight: 20
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
        </Toolbar>
      </AppBar>
    </>
  );
});
