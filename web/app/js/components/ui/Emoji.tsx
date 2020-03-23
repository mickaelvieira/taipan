import React from "react";
import { makeStyles } from "@material-ui/core/styles";
import { Emoji as EmojiBase } from "emoji-mart/dist-modern/index";

const useStyles = makeStyles(({ spacing }) => ({
  container: {
    display: "inline-block",
    verticalAlign: "middle",
    padding: `0 ${spacing(0.5)}px`,
    lineHeight: 0,
  },
}));

interface Props {
  emoji: string;
  className?: string;
  size?: number;
}

export default function Emoji({
  emoji,
  className,
  ...rest
}: Props): JSX.Element {
  const classes = useStyles();
  return (
    <span className={`${classes.container} ${className ? className : ""}`}>
      <EmojiBase
        set="google"
        emoji={emoji}
        sheetSize={32}
        size={32}
        {...rest}
      />
    </span>
  );
}
