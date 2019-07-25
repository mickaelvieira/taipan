import React from "react";
import { makeStyles } from "@material-ui/core/styles";
import Hidden from "@material-ui/core/Hidden";
import Table from "@material-ui/core/Table";
import Paper from "@material-ui/core/Paper";
import TableBody from "@material-ui/core/TableBody";
import TableCell from "@material-ui/core/TableCell";
import TableHead from "@material-ui/core/TableHead";
import TableRow from "@material-ui/core/TableRow";
import Button from "@material-ui/core/Button";
import Loader from "../../ui/Loader";

import SubscriptionsQuery, {
  variables,
  query,
  Data
} from "../../apollo/Query/Subscriptions";
import Row from "./Row";

const useStyles = makeStyles(() => ({
  table: {
    width: "100%"
  },
  fetchMore: {
    display: "flex",
    flexDirection: "column",
    alignItems: "center"
  },
  button: {
    margin: "12px"
  },
  message: {
    padding: 24
  }
}));

interface Props {
  terms: string[];
}

export default function SubscriptionsTable({ terms }: Props): JSX.Element {
  const classes = useStyles();
  return (
    <SubscriptionsQuery
      variables={
        terms.length === 0 ? variables : { ...variables, search: { terms } }
      }
    >
      {({ data, loading, error, fetchMore }) => {
        if (loading) {
          return <Loader />;
        }

        if (error) {
          return <Paper className={classes.message}>{error.message}</Paper>;
        }

        if (!data) {
          return null;
        }

        const { total, results } = data.subscriptions.subscriptions;
        const showLoadMoreButton = results.length < total;

        if (results.length === 0) {
          return (
            <Paper className={classes.message}>
              We could not find any sources matching your query.
            </Paper>
          );
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
                  <TableCell>Title</TableCell>
                  <Hidden mdDown>
                    <TableCell>Domain</TableCell>
                  </Hidden>
                  <Hidden mdDown>
                    <TableCell>Updated</TableCell>
                  </Hidden>
                  <TableCell align="center">Subscribed</TableCell>
                </TableRow>
              </TableHead>
              <TableBody>
                {results.map(subscription => (
                  <Row key={subscription.id} subscription={subscription} />
                ))}
              </TableBody>
            </Table>
            <div className={classes.fetchMore}>
              {showLoadMoreButton && (
                <Button
                  className={classes.button}
                  onClick={() =>
                    fetchMore({
                      query,
                      variables: {
                        ...variables,
                        pagination: {
                          ...variables.pagination,
                          offset:
                            data.subscriptions.subscriptions.results.length
                        }
                      },
                      updateQuery: (prev: Data, { fetchMoreResult: next }) => {
                        if (!next) {
                          return prev;
                        }
                        return {
                          subscriptions: {
                            ...prev.subscriptions,
                            subscriptions: {
                              ...prev.subscriptions.subscriptions,
                              limit: next.subscriptions.subscriptions.limit,
                              offset: next.subscriptions.subscriptions.offset,
                              results: [
                                ...prev.subscriptions.subscriptions.results,
                                ...next.subscriptions.subscriptions.results
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
              )}
            </div>
          </>
        );
      }}
    </SubscriptionsQuery>
  );
}
