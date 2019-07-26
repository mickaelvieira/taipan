import React from "react";
import { makeStyles } from "@material-ui/core/styles";
import IconButton from "@material-ui/core/IconButton";
import ShareIcon from "@material-ui/icons/ShareOutlined";
import { FeedItem } from "../../../../types/feed";

interface Props {
  item: FeedItem;
  onSucceed: (message: string) => void;
  onFail: (message: string) => void;
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
  onSucceed,
  onFail
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
            onSucceed("Thanks for sharing!");
          })
          .catch((error: Error) => {
            onFail(error.message);
          });
      }}
    >
      <ShareIcon />
    </IconButton>
  );
});
