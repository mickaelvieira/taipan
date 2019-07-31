import React from "react";
import { makeStyles } from "@material-ui/core/styles";
import ButtonBase from "@material-ui/core/ButtonBase";
import List from "@material-ui/core/List";
import ListItem from "@material-ui/core/ListItem";
import ListItemText from "@material-ui/core/ListItemText";
import DocumentSearchSearch, {
  query,
  variables,
  Data
} from "../../apollo/Query/Documents";

interface Props {
  terms: string[];
}

import NoResults from "./NoResults";
import Pagination from "./Pagination";
import Results from "./Results";

const useStyles = makeStyles(() => ({
  list: {
    "& mark": {
      backgroundColor: "yellow"
    }
  },
  button: {}
}));

interface Props {
  terms: string[];
}

export default function SearchDocuments({ terms }: Props): JSX.Element {
  const classes = useStyles();

  console.log(terms);
  return (
    <DocumentSearchSearch
      skip={terms.length === 0}
      variables={{ ...variables, search: { terms } }}
    >
      {({ data, loading, error, fetchMore }) => {
        if (error) {
          return <div>{error.message}</div>;
        }

        if (loading) {
          return <div>loading...</div>;
        }

        if (!data) {
          return null;
        }

        const {
          documents: {
            search: { results, total }
          }
        } = data;

        if (results.length === 0) {
          return <NoResults terms={terms} />;
        }

        const showLoadMoreButton = results.length < total;

        return (
          <List className={classes.list}>
            <Pagination count={results.length} total={total} terms={terms} />
            <Results results={results} terms={terms} />
            <Pagination count={results.length} total={total} terms={terms} />
            {showLoadMoreButton && (
              <ListItem>
                <ListItemText>
                  <ButtonBase
                    className={classes.button}
                    onClick={() =>
                      fetchMore({
                        query,
                        variables: {
                          ...variables,
                          pagination: {
                            ...variables.pagination,
                            offset: results.length
                          },
                          search: { terms }
                        },
                        updateQuery: (
                          prev: Data,
                          { fetchMoreResult: next }
                        ) => {
                          if (!next) {
                            return prev;
                          }
                          return {
                            documents: {
                              ...prev.documents,
                              search: {
                                ...prev.documents.search,
                                limit: next.documents.search.limit,
                                offset: next.documents.search.offset,
                                results: [
                                  ...prev.documents.search.results,
                                  ...next.documents.search.results
                                ]
                              }
                            }
                          };
                        }
                      })
                    }
                  >
                    Load more
                  </ButtonBase>
                </ListItemText>
              </ListItem>
            )}
          </List>
        );
      }}
    </DocumentSearchSearch>
  );
}
