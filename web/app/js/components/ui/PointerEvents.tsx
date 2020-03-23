import React, { PropsWithChildren } from "react";
import { makeStyles } from "@material-ui/core/styles";

const useStyles = makeStyles({
  scrolling: {
    pointerEvents: "none",
  },
  idle: {
    pointerEvents: "auto",
  },
});

interface Props {
  isScrolling: boolean;
}

export default function PointerEvents({
  children,
  isScrolling,
}: PropsWithChildren<Props>): JSX.Element {
  const classes = useStyles();
  return (
    <div className={isScrolling ? classes.scrolling : classes.idle}>
      {children}
    </div>
  );
}
