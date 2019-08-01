import React, { useState } from "react";
import { makeStyles } from "@material-ui/core/styles";
import Paper from "@material-ui/core/Paper";
import AddSubscriptionModal from "../../ui/Subscriptions/Panel/AddSubscription";
import { AddButton } from "../../ui/Fab";
import { LayoutRenderProps } from "../../Layout/Layout";
import Grid from "../../ui/Grid";
import Search from "./Search";

const useStyles = makeStyles(() => ({
  paper: {
    display: "flex",
    flexDirection: "column"
  }
}));

export default function Subscriptions({
  setIsContained,
  setMessageInfo
}: LayoutRenderProps): JSX.Element {
  const classes = useStyles();
  const [isModalOpen, setModalStatus] = useState(false);

  return (
    <>
      <Grid>
        <Paper className={classes.paper}>
          <Search />
        </Paper>
      </Grid>
      <AddButton
        onClick={() => {
          setIsContained(true);
          setModalStatus(true);
        }}
      />
      <AddSubscriptionModal
        isOpen={isModalOpen}
        toggleDialog={status => {
          setIsContained(status);
          setModalStatus(status);
        }}
        onSubscriptionCreated={() => {
          setMessageInfo({ message: "Nice one! The feed was added" });
          setIsContained(false);
          setModalStatus(false);
        }}
      />
    </>
  );
}
