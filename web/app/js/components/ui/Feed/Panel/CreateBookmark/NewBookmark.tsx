import React, { useState, useContext } from "react";
import { makeStyles } from "@material-ui/core/styles";
import Typography from "@material-ui/core/Typography";
import ButtonBase from "@material-ui/core/ButtonBase";
import { MessageContext } from "../../../../context";
import { Document } from "../../../../../types/document";
import { Bookmark } from "../../../../../types/bookmark";
import {
  BookmarkButton,
  BookmarkAndFavoriteButton
} from "../../../../ui/Feed/Button";
import Syndication from "./Syndication";

const useStyles = makeStyles(({ breakpoints, typography, palette }) => ({
  actions: {
    display: "flex",
    flexDirection: "column",
    alignItems: "flex-end",
    [breakpoints.up("md")]: {
      flexDirection: "row",
      alignItems: "center",
      justifyContent: "flex-end"
    }
  },
  button: {
    ...typography.body1,
    padding: "0 0 .06rem .2rem",
    color: palette.primary.main
  },
  cancel: {
    ...typography.body1,
    color: palette.secondary.main
  },
  addto: {
    display: "inline-block",
    paddingLeft: "0.3rem"
  }
}));

interface Props {
  document: Document;
  onFinish: (bookmark: Bookmark) => void;
  onCancel: () => void;
}

export default function NewBookmark({
  document,
  onCancel,
  onFinish
}: Props): JSX.Element {
  const classes = useStyles();
  const setMessageInfo = useContext(MessageContext);
  const [subscriptions, setSubscriptions] = useState<string[]>([]);

  return (
    <>
      {document.syndication && (
        <Syndication
          syndication={document.syndication}
          subscriptions={subscriptions}
          onChange={setSubscriptions}
        />
      )}
      <div className={classes.actions}>
        <Typography>Would you like to add this link to your</Typography>
        <div>
          <BookmarkButton
            textOnly
            document={document}
            className={classes.button}
            onSucceed={({ updateCache, item }) => {
              setMessageInfo({
                message: "The document was added to your reading list"
              });
              updateCache();
              onFinish(item as Bookmark);
            }}
            onFail={message => setMessageInfo({ message })}
          />
          {","}
          <BookmarkAndFavoriteButton
            textOnly
            document={document}
            className={classes.button}
            onSucceed={({ updateCache, item }) => {
              setMessageInfo({
                message: "The document was added to your favorites"
              });
              updateCache();
              onFinish(item as Bookmark);
            }}
            onFail={message => setMessageInfo({ message })}
          />
          <Typography className={classes.addto}>or</Typography>
          <ButtonBase
            onClick={onCancel}
            className={`${classes.button} ${classes.cancel}`}
          >
            cancel
          </ButtonBase>
          {"?"}
        </div>
      </div>
    </>
  );
}
