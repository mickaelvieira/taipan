import React, { useContext, useState } from "react";
import { makeStyles } from "@material-ui/core/styles";
import Paper from "@material-ui/core/Paper";
import { RouteSyndicationProps } from "../../../types/routes";
import { LayoutContext, MessageContext } from "../../context";
import CreateSyndication from "../../ui/Panel/CreateSyndication";
import { AddButton } from "../../ui/Fab";
import Grid from "../../ui/Grid";
import EditSource from "../../ui/Panel/EditSource";
import Search from "./Search";

const useStyles = makeStyles(() => ({
  paper: {
    display: "flex",
    flexDirection: "column",
  },
}));

export default function Syndication(_: RouteSyndicationProps): JSX.Element {
  const classes = useStyles();
  const setMessageInfo = useContext(MessageContext);
  const setIsContained = useContext(LayoutContext);
  const [isModalOpen, setModalStatus] = useState(false);
  const [editUrl, setEditURL] = useState<URL | null>(null);

  return (
    <>
      <Grid>
        <Paper className={classes.paper}>
          <Search editSource={setEditURL} />
        </Paper>
      </Grid>
      <AddButton
        onClick={() => {
          setIsContained(true);
          setModalStatus(true);
        }}
      />
      <CreateSyndication
        isOpen={isModalOpen}
        toggleDialog={(status) => {
          setIsContained(status);
          setModalStatus(status);
        }}
        onSyndicationCreated={() => {
          setMessageInfo({ message: "Nice one! The feed was added" });
          setIsContained(false);
          setModalStatus(false);
        }}
      />
      <EditSource
        url={editUrl}
        isOpen={editUrl !== null}
        close={() => setEditURL(null)}
      />
    </>
  );
}
