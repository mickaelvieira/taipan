import React from "react";
import IconButton from "@material-ui/core/Button";
import IconEdit from "@material-ui/icons/Edit";
import TableCell from "@material-ui/core/TableCell";
import TableRow from "@material-ui/core/TableRow";
import { Source } from "../../../../types/syndication";
import Link from "../Link";

interface Props {
  source: Source;
  editSource: (url: URL) => void;
}

export default React.memo(function Row({
  source,
  editSource
}: Props): JSX.Element {
  return (
    <TableRow>
      <TableCell>
        <Link item={source} />
      </TableCell>
      <TableCell align="center">
        <IconButton
          aria-label="Edit syndication source"
          size="small"
          onClick={() => editSource(source.url)}
        >
          <IconEdit fontSize="small" />
        </IconButton>
      </TableCell>
    </TableRow>
  );
});
