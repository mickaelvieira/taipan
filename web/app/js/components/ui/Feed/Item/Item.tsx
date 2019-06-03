import React, { ReactNode, useState } from "react";
import { makeStyles } from "@material-ui/core/styles";
import Fade from "@material-ui/core/Fade";
import Card from "@material-ui/core/Card";

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

interface RenderProps {
  remove: () => void;
}

interface Props {
  children: (props: RenderProps) => ReactNode;
}

export default function Item({ children }: Props) {
  const classes = useStyles();
  const [visible, setIsVisible] = useState(true);

  return (
    <Fade
      in={visible}
      unmountOnExit
      timeout={{
        enter: 1000,
        exit: 300
      }}
    >
      <Card className={classes.card}>
        {children({
          remove: () => setIsVisible(false)
        })}
      </Card>
    </Fade>
  );
}
