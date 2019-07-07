import React from "react";
import { makeStyles } from "@material-ui/core/styles";
import IconButton from "@material-ui/core/IconButton";
import ShareIcon from "@material-ui/icons/ShareOutlined";
import { FeedItem } from "../../../../types/feed";

interface Props {
  item: FeedItem;
  onSuccess: (message: string) => void;
  onError: (message: string) => void;
}

const useStyles = makeStyles(({ palette }) => ({
  button: {
    color: palette.primary.main
  }
}));

function canShare(): boolean {
  return "share" in navigator;
}

export default React.memo(function ShareButton({
  item,
  onSuccess,
  onError
}: Props): JSX.Element | null {
  const classes = useStyles();
  return !canShare() ? null : (
    <IconButton
      aria-label="Share"
      className={classes.button}
      onClick={() => {
        navigator
          .share({
            title: item.title,
            url: item.url
          })
          .then(() => {
            onSuccess("Thanks for sharing!");
          })
          .catch((error: Error) => {
            onError(error.message);
          });
      }}
    >
      <ShareIcon />
    </IconButton>
  );
});
