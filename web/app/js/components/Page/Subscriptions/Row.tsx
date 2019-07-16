import React from "react";
import { makeStyles } from "@material-ui/core/styles";
import Hidden from "@material-ui/core/Hidden";
import TableCell from "@material-ui/core/TableCell";
import TableRow from "@material-ui/core/TableRow";
import Link from "@material-ui/core/Link";
import { Subscription } from "../../../types/subscription";
import Domain from "../../ui/Domain";
import { StatusButton } from "../../ui/Subscriptions/Button";
import SubscriptionTitle from "./Title";

const useStyles = makeStyles({
  link: {
    display: "inline-block",
    padding: "16px 6px 12px",
    position: "relative",
    verticalAlign: "middle",
    boxSizing: "border-box",
    textAlign: "center"
  }
});

interface Props {
  subscription: Subscription;
}

export default React.memo(function Row({ subscription }: Props): JSX.Element {
  const classes = useStyles();
  const { title } = subscription;
  const url = new URL(subscription.url);

  return (
    <TableRow>
      <TableCell>
        {title ? (
          <SubscriptionTitle item={subscription} />
        ) : (
          <Domain item={subscription} />
        )}
      </TableCell>
      <Hidden mdDown>
        <TableCell>
          <Link
            underline="none"
            href={`${url.protocol}//${url.host}`}
            title={subscription.title ? subscription.title : subscription.url}
            target="_blank"
            rel="noopener"
            className={classes.link}
          >
            {`${url.protocol}//${url.host}`}
          </Link>
        </TableCell>
      </Hidden>
      <TableCell align="center">
        <StatusButton subscription={subscription} />
      </TableCell>
    </TableRow>
  );
});
