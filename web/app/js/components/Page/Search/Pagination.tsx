import React from "react";
// import { makeStyles } from "@material-ui/core/styles";
import ListItem from "@material-ui/core/ListItem";
import ListItemText from "@material-ui/core/ListItemText";
import Terms from "./Terms";

// const useStyles = makeStyles(({ spacing, palette }) => ({
//   list: {
//     "& mark": {
//       backgroundColor: "yellow"
//     }
//   },
//   message: {
//     padding: spacing(2)
//   },
//   button: {}
// }));

interface Props {
  count: number;
  total: number;
  terms: string[];
}

export default function SearchBookmarks({
  count,
  total,
  terms
}: Props): JSX.Element {
  // const classes = useStyles();

  return (
    <ListItem>
      <ListItemText>
        {count} results of {total} matching <Terms terms={terms} />
      </ListItemText>
    </ListItem>
  );
}
