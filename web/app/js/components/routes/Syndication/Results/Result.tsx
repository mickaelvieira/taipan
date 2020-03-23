import React from "react";
import { makeStyles } from "@material-ui/core/styles";
import IconButton from "@material-ui/core/Button";
import Chip from "@material-ui/core/Chip";
import IconEdit from "@material-ui/icons/Edit";
import TableCell from "@material-ui/core/TableCell";
import TableRow from "@material-ui/core/TableRow";
import { Source } from "../../../../types/syndication";
import Link from "../Link";
import { sort } from "../../../../helpers/tags";

const useStyles = makeStyles(({ palette, spacing }) => ({
  chip: {
    margin: spacing(1),
  },
  active: {
    color: palette.common.white,
    backgroundColor: palette.primary.main,
  },
}));

interface Props {
  source: Source;
  editSource: (url: URL) => void;
}

export default React.memo(function Row({
  source,
  editSource,
}: Props): JSX.Element {
  const classes = useStyles();
  const list = sort(source.tags ? source.tags : []);
  return (
    <TableRow>
      <TableCell>
        <Link item={source} />
      </TableCell>
      <TableCell>
        {list.map((tag) => (
          <Chip
            key={tag.id}
            size="small"
            color="primary"
            label={tag.label}
            className={classes.chip}
          />
        ))}
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
