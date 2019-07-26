import React, { useContext } from "react";
import { makeStyles } from "@material-ui/core/styles";
import Typography from "@material-ui/core/Typography";
import ButtonBase from "@material-ui/core/ButtonBase";
import { MessageContext } from "../../../../context";
import { Bookmark } from "../../../../../types/bookmark";
import { FavoriteButton } from "../../../../ui/Feed/Button";

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
  intro: {
    alignSelf: "flex-end",
    textAlign: "right"
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
  bookmark: Bookmark;
  onFinish: (bookmark: Bookmark) => void;
  onCancel: () => void;
}

export default function ExistingBookmark({
  bookmark,
  onCancel,
  onFinish
}: Props): JSX.Element {
  const classes = useStyles();
  const setMessageInfo = useContext(MessageContext);

  return (
    <>
      {!bookmark.isFavorite && (
        <Typography className={classes.intro}>
          You really like this page, this link is already in your bookmarks.
        </Typography>
      )}
      <div className={classes.actions}>
        {!bookmark.isFavorite && (
          <Typography>Do you want me to add it to your</Typography>
        )}
        {bookmark.isFavorite && (
          <Typography>
            You really like this link, it is already in your favorites.
          </Typography>
        )}
        <div>
          {!bookmark.isFavorite && (
            <>
              <FavoriteButton
                textOnly
                bookmark={bookmark}
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
            </>
          )}
          <ButtonBase
            onClick={onCancel}
            className={`${classes.button} ${classes.cancel}`}
          >
            Cancel
          </ButtonBase>
          {!bookmark.isFavorite && "?"}
        </div>
      </div>
    </>
  );
}
