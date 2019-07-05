import React, { useState, PropsWithChildren, useContext } from "react";
import { makeStyles } from "@material-ui/core/styles";
import Fab from "@material-ui/core/Fab";
import AddIcon from "@material-ui/icons/Add";
import Grid from "@material-ui/core/Grid";
import AddSource from "../ui/Syndication/Modal/AddSource";
import { MessageContext } from "../context";
import MainLayout from "./Layout";

const useStyles = makeStyles(({ palette, spacing }) => ({
  root: {
    width: "100%"
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
  }
}));

export default function LayoutSyndication({
  children
}: PropsWithChildren<{}>): JSX.Element {
  const classes = useStyles();
  const setInfo = useContext(MessageContext);
  const [isModalOpen, setModalStatus] = useState(false);

  return (
    <MainLayout>
      <Grid item xs={12} className={classes.root}>
        <div className={classes.content}>{children}</div>
        <Fab
          color="primary"
          size="large"
          aria-label="Add"
          className={classes.fab}
          onClick={() => setModalStatus(true)}
        >
          <AddIcon />
        </Fab>
        <AddSource
          isOpen={isModalOpen}
          toggleDialog={setModalStatus}
          onSyndicationSourceCreated={() => {
            setInfo("Nice one! The feed was added");
            setModalStatus(false);
          }}
        />
      </Grid>
    </MainLayout>
  );
}
