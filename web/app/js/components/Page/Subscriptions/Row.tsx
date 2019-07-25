import React from "react";
import { useTheme } from "@material-ui/core/styles";
import useMediaQuery from "@material-ui/core/useMediaQuery";
import IconButton from "@material-ui/core/Button";
import IconEdit from "@material-ui/icons/Edit";
import TableCell from "@material-ui/core/TableCell";
import TableRow from "@material-ui/core/TableRow";
import { Subscription } from "../../../types/subscription";
import { StatusButton } from "../../ui/Subscriptions/Button";
import Title from "./Title";
import Link from "./Link";

interface Props {
  canEdit: boolean;
  subscription: Subscription;
  editSource: (url: string) => void;
}

export default React.memo(function Row({
  subscription,
  editSource,
  canEdit = false
}: Props): JSX.Element {
  const { title } = subscription;
  const theme = useTheme();
  const md = useMediaQuery(theme.breakpoints.up("md"));

  return (
    <TableRow>
      <TableCell>
        {title ? (
          <Title item={subscription} shouldTruncate />
        ) : (
          <Link item={subscription} />
        )}
      </TableCell>
      {md && (
        <TableCell>
          <Link item={subscription} />
        </TableCell>
      )}
      {md && (
        <TableCell align="center">
          {!canEdit ? (
            subscription.frequency
          ) : (
            <IconButton
              size="small"
              onClick={() => editSource(subscription.url)}
            >
              <IconEdit fontSize="small" />
            </IconButton>
          )}
        </TableCell>
      )}
      <TableCell align="center">
        <StatusButton subscription={subscription} />
      </TableCell>
    </TableRow>
  );
});
