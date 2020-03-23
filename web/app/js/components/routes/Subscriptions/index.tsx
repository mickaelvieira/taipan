import React, { useContext, useState } from "react";
import { makeStyles } from "@material-ui/core/styles";
import Paper from "@material-ui/core/Paper";
import { RouteSubscriptionProps } from "../../../types/routes";
import { LayoutContext, MessageContext } from "../../context";
import CreateSubscription from "../../ui/Panel/CreateSubscription";
import { AddButton } from "../../ui/Fab";
import Grid from "../../ui/Grid";
import Search from "./Search";

const useStyles = makeStyles(() => ({
  paper: {
    display: "flex",
    flexDirection: "column",
  },
}));

export default function Subscriptions(_: RouteSubscriptionProps): JSX.Element {
  const classes = useStyles();
  const setMessageInfo = useContext(MessageContext);
  const setIsContained = useContext(LayoutContext);
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
      <CreateSubscription
        isOpen={isModalOpen}
        toggleDialog={(status) => {
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
