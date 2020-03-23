import React, { PropsWithChildren } from "react";
import { makeStyles } from "@material-ui/core/styles";
import Card from "@material-ui/core/Card";

const useStyles = makeStyles(({ breakpoints }) => ({
  card: {
    marginBottom: 24,
    display: "flex",
    flexDirection: "column",
    borderRadius: 0,
    [breakpoints.up("sm")]: {
      borderRadius: 4,
    },
  },
}));

export default React.memo(function Item({
  children,
}: PropsWithChildren<{}>): JSX.Element {
  const classes = useStyles();

  return (
    <div>
      <Card className={`${classes.card} feed-item`}>{children}</Card>
    </div>
  );
});
