import React from "react";
// import { makeStyles } from "@material-ui/core/styles";
import Link from "@material-ui/core/Link";
import ListItem from "@material-ui/core/ListItem";
import ListItemText from "@material-ui/core/ListItemText";
import Typography from "@material-ui/core/Typography";
import Highlighter from "react-highlight-words";
import { Bookmark } from "../../../types/bookmark";
import { Document } from "../../../types/document";

import Domain from "../../ui/Domain";
import { truncate } from "../../../helpers/string";

// const useStyles = makeStyles(({ spacing, palette }) => ({
//   list: {
//     "& mark": {
//       backgroundColor: "yellow"
//     }
//   },
//   term: {
//     display: "inline-block",
//     padding: "0 4px",
//     margin: "0 2px",
//     color: palette.common.white,
//     backgroundColor: palette.primary.main
//   },
//   message: {
//     padding: spacing(2)
//   },
//   button: {}
// }));

interface Props {
  terms: string[];
  results: Bookmark[] | Document[];
}

export default function SearchBookmarks({
  results,
  terms
}: Props): JSX.Element {
  // const classes = useStyles();

  console.log(terms);
  return results.map(result => (
    <ListItem key={result.id}>
      <ListItemText>
        <Link
          underline="none"
          href={result.url}
          title={result.title}
          target="_blank"
          rel="noopener"
        >
          {result.title ? (
            <Highlighter
              autoEscape
              searchWords={terms}
              textToHighlight={result.title}
            />
          ) : (
            "[no title available]"
          )}
        </Link>
        <Typography>
          {result.description ? (
            <Highlighter
              autoEscape
              searchWords={terms}
              textToHighlight={truncate(result.description, 200)}
            />
          ) : (
            "[no description available]"
          )}
        </Typography>
        <Domain item={result} />
      </ListItemText>
    </ListItem>
  ));
}
