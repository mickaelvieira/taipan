import React from "react";
import { makeStyles } from "@material-ui/core/styles";
import Dialog from "@material-ui/core/Dialog";
import Tabs from "@material-ui/core/Tabs";
import Tab from "@material-ui/core/Tab";
import AddBookmark from "./AddBookmark";
import AddFeed from "./AddFeed";
import { Bookmark } from "../../types/bookmark";
import { Source } from "../../types/syndication";

interface Props {
  isOpen: boolean;
  toggleDialog: (status: boolean) => void;
  onBookmarkCreated: (bookmark: Bookmark) => void;
  onSyndicationSourceCreated: (source: Source) => void;
}

const useStyles = makeStyles(() => ({
  dialog: {}
}));

export default function AddForm({
  isOpen,
  toggleDialog,
  onBookmarkCreated,
  onSyndicationSourceCreated
}: Props) {
  const classes = useStyles();
  const [value, setValue] = React.useState(0);

  function handleChange(event: React.ChangeEvent<{}>, newValue: number) {
    setValue(newValue);
  }

  return (
    <Dialog
      open={isOpen}
      onClose={() => toggleDialog(false)}
      aria-labelledby="form-dialog-title"
      className={classes.dialog}
    >
      <Tabs
        value={value}
        onChange={handleChange}
        indicatorColor="primary"
        textColor="primary"
        variant="fullWidth"
      >
        <Tab value={0} label="Bookmark" />
        <Tab value={1} label="Feed" />
      </Tabs>
      {value === 0 && (
        <AddBookmark
          onBookmarkCreated={onBookmarkCreated}
          toggleDialog={toggleDialog}
        />
      )}
      {value === 1 && (
        <AddFeed
          onSyndicationSourceCreated={onSyndicationSourceCreated}
          toggleDialog={toggleDialog}
        />
      )}
    </Dialog>
  );
}
