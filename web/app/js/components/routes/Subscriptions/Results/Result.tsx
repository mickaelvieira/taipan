import React from "react";
import { useTheme } from "@material-ui/core/styles";
import useMediaQuery from "@material-ui/core/useMediaQuery";
import TableCell from "@material-ui/core/TableCell";
import TableRow from "@material-ui/core/TableRow";
import { Subscription } from "../../../../types/subscription";
import { StatusButton } from "../../../ui/Subscriptions/Button";
import Title from "../Title";
import Domain from "../Domain";

interface Props {
  subscription: Subscription;
}

export default React.memo(function Row({ subscription }: Props): JSX.Element {
  const { title } = subscription;
  const theme = useTheme();
  const md = useMediaQuery(theme.breakpoints.up("md"));

  return (
    <TableRow>
      <TableCell>
        {title ? (
          <Title item={subscription} shouldTruncate />
        ) : (
          <Domain item={subscription} />
        )}
      </TableCell>
      {md && (
        <TableCell>
          <Domain item={subscription} />
        </TableCell>
      )}
      {md && <TableCell align="center">{subscription.frequency}</TableCell>}
      <TableCell align="center">
        <StatusButton subscription={subscription} />
      </TableCell>
    </TableRow>
  );
});
