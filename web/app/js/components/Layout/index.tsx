import React, { useState, FunctionComponent } from "react";
import { WithStyles, createStyles, Theme } from "@material-ui/core";
import { withStyles } from "@material-ui/core/styles";
import Paper from "@material-ui/core/Paper";
import Grid from "@material-ui/core/Grid";
import Header from "./Header";
import Sidebar from "./Navigation/Sidebar";

const styles = ({ palette, spacing }: Theme) =>
  createStyles({
    root: {
      flexGrow: 1
    },
    paper: {
      padding: spacing.unit * 2,
      textAlign: "center",
      color: palette.text.secondary
    }
  });

interface Props extends WithStyles<typeof styles> {}

const Layout: FunctionComponent<Props> = ({ children, classes }) => {
  const [isOpen, setIsOpen] = useState(false);

  return (
    <div className={classes.root}>
      <Sidebar isOpen={isOpen} toggleDrawer={setIsOpen} />
      <Header toggleDrawer={setIsOpen} />
      <Grid container spacing={24}>
        <Grid item xs={12}>
          <Paper className={classes.paper}>{children}</Paper>
        </Grid>
      </Grid>
    </div>
  );
};

export default withStyles(styles)(Layout);
