import React from "react";
import { makeStyles } from "@material-ui/core/styles";
import Domain, { DomainProps } from "../../Domain";

const useStyles = makeStyles(() => ({
  link: {
    padding: 12
  }
}));

export default function FeedItemDomain(props: DomainProps): JSX.Element {
  const classes = useStyles();
  return <Domain {...props} className={classes.link} />;
}
