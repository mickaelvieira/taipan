import React from "react";
import { makeStyles } from "@material-ui/core/styles";

import Paper from "@material-ui/core/Paper";
import Layout from "../../Layout/Subscriptions";
import ScrollToTop from "../../ui/ScrollToTop";
import Search from "./Search";

const useStyles = makeStyles(() => ({
  paper: {
    display: "flex",
    flexDirection: "column"
  }
}));

export default function Subscriptions(): JSX.Element {
  const classes = useStyles();

  return (
    <Layout>
      <ScrollToTop>
        <Paper className={classes.paper}>
          <Search />
        </Paper>
      </ScrollToTop>
    </Layout>
  );
}
