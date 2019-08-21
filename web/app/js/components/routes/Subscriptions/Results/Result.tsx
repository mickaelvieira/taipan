import React from "react";
import { useTheme } from "@material-ui/core/styles";
import useMediaQuery from "@material-ui/core/useMediaQuery";
import IconButton from "@material-ui/core/Button";
import IconEdit from "@material-ui/icons/Edit";
import TableCell from "@material-ui/core/TableCell";
import TableRow from "@material-ui/core/TableRow";
import { Subscription } from "../../../../types/subscription";
import { StatusButton } from "../../../ui/Subscriptions/Button";
import Title from "../Title";
import Domain from "../Domain";

interface Props {
  subscription: Subscription;
  canEdit?: boolean;
  editSource?: (url: URL) => void;
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
          <Domain item={subscription} />
        )}
      </TableCell>
      {md && (
        <TableCell>
          <Domain item={subscription} />
        </TableCell>
      )}
      {md && (
        <TableCell align="center">
          {!canEdit ? (
            subscription.frequency
          ) : (
            <IconButton
              aria-label="Edit syndication source"
              size="small"
              onClick={() => {
                if (typeof editSource === "function") {
                  editSource(subscription.url);
                }
              }}
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
