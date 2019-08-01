import React from "react";
import { makeStyles } from "@material-ui/core/styles";
import List from "@material-ui/core/List";
import ListItem from "@material-ui/core/ListItem";
import ListItemText from "@material-ui/core/ListItemText";
import Typography from "@material-ui/core/Typography";
import Highlighter from "react-highlight-words";
import { SearchItem, SearchType } from "../../../types/search";
import Domain from "../../ui/Domain";
import ExternalLink from "../../ui/Link/External";
import { truncate } from "../../../helpers/string";
import NoResults from "./NoResults";

interface Props {
  type: SearchType;
  terms: string[];
  results: SearchItem[];
}
const useStyles = makeStyles(() => ({
  list: {
    "& mark": {
      backgroundColor: "yellow"
    }
  }
}));

export default function Results({ results, type, terms }: Props): JSX.Element {
  const classes = useStyles();
  if (results.length === 0) {
    return <NoResults type={type} terms={terms} />;
  }

  return (
    <List className={classes.list}>
      {results.map(result => (
        <ListItem key={result.id}>
          <ListItemText>
            <ExternalLink href={result.url} title={result.title}>
              {result.title ? (
                <Highlighter
                  autoEscape
                  searchWords={terms}
                  textToHighlight={result.title}
                />
              ) : (
                "[no title available]"
              )}
            </ExternalLink>
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
      ))}
    </List>
  );
}
