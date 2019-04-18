import React from "react";
import { WithStyles, createStyles } from "@material-ui/core";
import { withStyles } from "@material-ui/core/styles";
import AppBar from "@material-ui/core/AppBar";
import Toolbar from "@material-ui/core/Toolbar";
import Button from "@material-ui/core/Button";
import IconButton from "@material-ui/core/IconButton";
import MenuIcon from "@material-ui/icons/Menu";

const styles = () =>
  createStyles({
    root: {
      flexGrow: 1
    },
    grow: {
      flexGrow: 1
    },
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
    <div className={classes.root}>
      <AppBar position="static">
        <Toolbar>
          <IconButton
            className={classes.menuButton}
            color="inherit"
            aria-label="Menu"
          >
            <MenuIcon onClick={() => toggleDrawer(true)} />
          </IconButton>
          <div className={classes.grow} />
          <Button color="inherit">Login</Button>
        </Toolbar>
      </AppBar>
    </div>
  );
});
