import React from "react";
import { ApolloError } from "apollo-client";
import { makeStyles } from "@material-ui/core/styles";
import Grid from "@material-ui/core/Grid";
import Paper from "@material-ui/core/Paper";
import Typography from "@material-ui/core/Typography";

const useStyles = makeStyles(({ palette, spacing }) => ({
  root: {
    display: "flex",
    flexDirection: "column",
    alignItems: "center"
  },
  inner: {
    width: "100%",
    maxWidth: 800
  },
  content: {
    display: "flex",
    flexDirection: "column",
    alignItems: "center",
    color: palette.text.secondary,
    paddingTop: 70
  },
  fab: {
    margin: spacing(1),
    position: "fixed",
    bottom: spacing(2),
    right: spacing(2),
    backgroundColor: palette.secondary.main,
    "&:hover": {
      backgroundColor: palette.secondary.light
    }
  },
  message: {
    display: "flex",
    alignItems: "center"
  }
}));

interface Props {
  error: ApolloError;
}

export default React.memo(function Error({ error }: Props) {
  const classes = useStyles();

  return (
    <Grid container className={classes.root}>
      <Grid item xs={12} className={classes.inner}>
        <Paper className={classes.root}>
          <Typography variant="h5" component="h3">
            Oops
          </Typography>
          <Typography component="p">Something went wrong</Typography>
          <Typography component="p">{error.message}</Typography>
          <Typography>
            {error.graphQLErrors.map(({ message, locations, path }, index) => {
              const msg = `[GraphQL error]: Message: ${message}, Location: ${locations}, Path: ${path}`;
              return <span key={index}>{msg}</span>;
            })}
          </Typography>
        </Paper>
      </Grid>
    </Grid>
  );
});
