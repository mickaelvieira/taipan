import React from "react";
import { makeStyles } from "@material-ui/core/styles";
import Layout from "../../Layout/Syndication";
import Table from "./Table";
import Typography from "@material-ui/core/Typography";

const useStyles = makeStyles(({ typography, palette }) => ({
  title: {
    alignSelf: "flex-start",
    margin: "24px",
    fontWeight: 500,
    fontSize: typography.h5.fontSize,
    color: palette.grey[900],
  }
}));

export default function Syndication(): JSX.Element {
  const classes = useStyles();

  return (
    <Layout>
      <Typography component="h2" variant="h2" className={classes.title}>Web syndication</Typography>
      <Table />
    </Layout>
  );
}
