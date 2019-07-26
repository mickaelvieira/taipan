import React, { useContext } from "react";
import { makeStyles } from "@material-ui/core/styles";
import { useTheme } from "@material-ui/core/styles";
import useMediaQuery from "@material-ui/core/useMediaQuery";
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
import { UserContext } from "../../context";
import { isAdmin } from "../../../helpers/users";

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
  showDeleted: boolean;
  pausedOnly: boolean;
  editSource: (url: string) => void;
}

export default React.memo(function SubscriptionsTable({
  terms,
  showDeleted,
  pausedOnly,
  editSource
}: Props): JSX.Element {
  const classes = useStyles();
  const user = useContext(UserContext);
  const theme = useTheme();
  const md = useMediaQuery(theme.breakpoints.up("md"));
  const canEdit = isAdmin(user);

  return (
    <SubscriptionsQuery
      fetchPolicy="network-only"
      variables={{ ...variables, search: { terms, pausedOnly, showDeleted } }}
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
                  {md && <TableCell>Domain</TableCell>}
                  {md && <TableCell>{canEdit ? "" : "Updated"}</TableCell>}
                  <TableCell align="center">Subscribed</TableCell>
                </TableRow>
              </TableHead>
              <TableBody>
                {results.map(subscription => (
                  <Row
                    key={subscription.id}
                    canEdit={canEdit}
                    editSource={editSource}
                    subscription={subscription}
                  />
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
                        },
                        search: { terms, pausedOnly, showDeleted }
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
});
