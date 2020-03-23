import React, { PropsWithChildren } from "react";
import { makeStyles } from "@material-ui/core/styles";
import useScrollStatus from "../../../hooks/useScrollStatus";

const useStyles = makeStyles({
  idle: {
    pointerEvents: "auto",
  },
  scrolling: {
    pointerEvents: "none",
  },
});

export default React.memo(function PointerEvents({
  children,
}: PropsWithChildren<{}>): JSX.Element {
  const classes = useStyles();
  const isScrolling = useScrollStatus();

  return (
    <div className={`${isScrolling ? classes.scrolling : classes.idle}`}>
      {children}
    </div>
  );
});
