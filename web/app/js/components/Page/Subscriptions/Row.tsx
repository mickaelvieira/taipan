import React from "react";
import Hidden from "@material-ui/core/Hidden";
import TableCell from "@material-ui/core/TableCell";
import TableRow from "@material-ui/core/TableRow";
import { Subscription } from "../../../types/subscription";
import { StatusButton } from "../../ui/Subscriptions/Button";
import Title from "./Title";
import Link from "./Link";

interface Props {
  subscription: Subscription;
}

export default React.memo(function Row({ subscription }: Props): JSX.Element {
  const { title } = subscription;

  return (
    <TableRow>
      <TableCell>
        {title ? (
          <Title subscription={subscription} />
        ) : (
          <Link subscription={subscription} />
        )}
      </TableCell>
      <Hidden mdDown>
        <TableCell>
          <Link subscription={subscription} />
        </TableCell>
      </Hidden>
      <Hidden mdDown>
        <TableCell align="center">{subscription.frequency}</TableCell>
      </Hidden>
      <TableCell align="center">
        <StatusButton subscription={subscription} />
      </TableCell>
    </TableRow>
  );
});
