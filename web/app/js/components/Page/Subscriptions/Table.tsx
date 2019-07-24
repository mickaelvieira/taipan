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
  query
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

export default function SubscriptionsTable(): JSX.Element {
  const classes = useStyles();

  return (
    <SubscriptionsQuery>
      {({ data, loading, error, fetchMore }) => {
        if (loading) {
          return <Loader />;
        }

        if (error) {
          return <span>{error.message}</span>;
        }

        if (!data) {
          return (
            <Paper className={classes.message}>
              You don&apos;t have any web syndication sources yet.
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
                {data.subscriptions.subscriptions.results.map(subscription => {
                  return (
                    <Row key={subscription.id} subscription={subscription} />
                  );
                })}
              </TableBody>
            </Table>
            <div className={classes.fetchMore}>
              <Button
                className={classes.button}
                onClick={() =>
                  fetchMore({
                    query,
                    variables: {
                      ...variables,
                      pagination: {
                        ...variables.pagination,
                        offset: data.subscriptions.subscriptions.results.length
                      }
                    },
                    updateQuery: (prev, { fetchMoreResult: next }) => {
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
            </div>
          </>
        );
      }}
    </SubscriptionsQuery>
  );
}
