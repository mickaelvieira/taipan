import React from "react";
import { useQuery } from "@apollo/react-hooks";
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

import {
  Data,
  Variables,
  query,
  variables,
  getFetchMore
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
  showDeleted: boolean;
  pausedOnly: boolean;
  canEdit?: boolean;
  editSource?: (url: URL) => void;
}

export default React.memo(function SubscriptionsTable({
  terms,
  showDeleted,
  pausedOnly,
  canEdit = false,
  editSource
}: Props): JSX.Element | null {
  const classes = useStyles();
  const theme = useTheme();
  const md = useMediaQuery(theme.breakpoints.up("md"));
  const { data, loading, error, fetchMore } = useQuery<Data, Variables>(query, {
    fetchPolicy: "network-only",
    variables: { ...variables, search: { terms, pausedOnly, showDeleted } }
  });

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
        {terms.length === 0 &&
          "You haven't subscribed to any feeds. Use the search field above to find new sources of excitement."}
        {terms.length > 0 &&
          "We could not find any sources matching your query."}
      </Paper>
    );
  }

  return (
    <>
      <Table className={classes.table} size="small">
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
              getFetchMore(fetchMore, data, {
                ...variables,
                pagination: {
                  ...variables.pagination,
                  offset: results.length
                },
                search: { terms, pausedOnly, showDeleted }
              })
            }
          >
            Load more
          </Button>
        )}
      </div>
    </>
  );
});
