import React, { useRef, useState } from "react";
import { makeStyles } from "@material-ui/core/styles";
import Fade from "@material-ui/core/Fade";
import Card from "@material-ui/core/Card";
import { CacheUpdater } from "../../../../types";

const useStyles = makeStyles(({ breakpoints }) => ({
  card: {
    marginBottom: 24,
    display: "flex",
    flexDirection: "column",
    borderRadius: 0,
    [breakpoints.up("sm")]: {
      borderRadius: 4
    }
  }
}));

interface Props {
  children: (props: RenderProps) => JSX.Element;
}

interface RenderProps {
  remove: (cb: CacheUpdater) => void;
}

export default function Item({ children }: Props): JSX.Element {
  const classes = useStyles();
  const ref = useRef<CacheUpdater>();
  const [visible, setIsVisible] = useState(true);

  const remove = (cb: CacheUpdater): void => {
    ref.current = cb;
    setIsVisible(false);
  };

  return (
    <Fade
      in={visible}
      unmountOnExit
      timeout={{
        enter: 500,
        exit: 400
      }}
      onExited={() => {
        if (typeof ref.current === "function") {
          ref.current();
        }
      }}
    >
      <Card className={classes.card}>
        {children({
          remove
        })}
      </Card>
    </Fade>
  );
}
