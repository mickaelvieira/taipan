import React, { useState, useContext } from "react";
import { makeStyles } from "@material-ui/core/styles";
import Slide from "@material-ui/core/Slide";
import Paper from "@material-ui/core/Paper";
import IconButton from "@material-ui/core/IconButton";
import CloseIcon from "@material-ui/icons/Close";
import { Document } from "../../../../types/document";
import { Bookmark } from "../../../../types/bookmark";
import FormDocument from "./FormDocument";
import FormBookmark from "./FormBookmark";
import { Typography } from "@material-ui/core";
import { FeedsCacheContext } from "../../../context";

const useStyles = makeStyles(({ palette }) => ({
  dialog: {},
  paper: {
    zIndex: 10000,
    position: "fixed",
    width: "100vw",
    height: "100vh"
  },
  header: {
    display: "flex",
    flexDirection: "row",
    justifyContent: "start",
    margin: 0,
    padding: 0,
    backgroundColor: palette.grey[200]
  },
  title: {
    paddingTop: 12,
    paddingBottom: 12
  },
  container: {
    padding: 16,
    display: "flex",
    flexDirection: "column"
  }
}));

interface Props {
  isOpen: boolean;
  setIsPanelOpen: (isOpen: boolean) => void;
  onBookmarkCreated: (bookmark: Bookmark) => void;
}

export default function AddForm({
  isOpen,
  setIsPanelOpen,
  onBookmarkCreated
}: Props): JSX.Element {
  const classes = useStyles();
  const updater = useContext(FeedsCacheContext);
  const [document, setDocument] = useState<Document | null>(null);

  return (
    <Slide in={isOpen} direction="up" mountOnEnter unmountOnExit>
      <Paper className={classes.paper} elevation={0} square>
        <header className={classes.header}>
          <IconButton onClick={() => setIsPanelOpen(false)}>
            <CloseIcon />
          </IconButton>
          <Typography component="h5" variant="h6" className={classes.title}>
            Bookmark a webpage
          </Typography>
        </header>
        <section className={classes.container}>
          {!document && <FormDocument setDocument={setDocument} />}
          {document && (
            <FormBookmark
              document={document}
              onRemoveBookmark={() => setDocument(null)}
              onBookmarkCreated={bookmark => {
                if (updater) {
                  updater.bookmark(bookmark);
                }
                setDocument(null);
                onBookmarkCreated(bookmark);
              }}
            />
          )}
        </section>
      </Paper>
    </Slide>
  );
}
