import React from "react";
import { makeStyles } from "@material-ui/core/styles";
import Hidden from "@material-ui/core/Hidden";
import Table from "@material-ui/core/Table";
import TableBody from "@material-ui/core/TableBody";
import TableCell from "@material-ui/core/TableCell";
import TableHead from "@material-ui/core/TableHead";
import TableRow from "@material-ui/core/TableRow";
import Button from "@material-ui/core/Button";
import Loader from "../../ui/Loader";

import SyndicationQuery, {
  variables,
  query
} from "../../apollo/Query/Syndication";
import Row from "./Row";

const useStyles = makeStyles(({ breakpoints }) => ({
  table: {
    width: "100%",
    [breakpoints.up("md")]: {
      width: "80%"
    }
  },
  button: {
    margin: "24px"
  }
}));

export default function SourcesTable(): JSX.Element {
  const classes = useStyles();

  return (
    <SyndicationQuery
      variables={{
        ...variables,
        search: {
          isPaused: false
        }
      }}
    >
      {({ data, loading, error, fetchMore }) => {
        if (loading) {
          return <Loader />;
        }

        if (error) {
          return <span>{error.message}</span>;
        }

        if (!data) {
          return null;
        }

        return (
          <>
            <Table
              className={classes.table}
              aria-labelledby="tableTitle"
              size="small"
            >
              <TableHead>
                <TableRow>
                  <TableCell align="center">Title</TableCell>
                  <Hidden mdDown>
                    <TableCell align="center">Domain</TableCell>
                  </Hidden>
                  <TableCell align="center">Active</TableCell>
                  <Hidden mdDown>
                    <TableCell></TableCell>
                  </Hidden>
                </TableRow>
              </TableHead>
              <TableBody>
                {data.syndication.sources.results.map(source => {
                  return <Row key={source.id} source={source} />;
                })}
              </TableBody>
            </Table>

            <Button
              className={classes.button}
              onClick={() =>
                fetchMore({
                  query,
                  variables: {
                    ...variables,
                    pagination: {
                      ...variables.pagination,
                      offset: data.syndication.sources.results.length
                    }
                  },
                  updateQuery: (prev, { fetchMoreResult: next }) => {
                    if (!next) {
                      return prev;
                    }
                    return {
                      syndication: {
                        ...prev.syndication,
                        sources: {
                          ...prev.syndication.sources,
                          limit: next.syndication.sources.limit,
                          offset: next.syndication.sources.offset,
                          results: [
                            ...prev.syndication.sources.results,
                            ...next.syndication.sources.results
                          ]
                        }
                      }
                    };
                  }
                })
              }
            >
              Load more
            </Button>
          </>
        );
      }}
    </SyndicationQuery>
  );
}
