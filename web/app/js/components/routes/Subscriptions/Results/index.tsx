import React from "react";
import { useQuery } from "@apollo/react-hooks";
import { makeStyles } from "@material-ui/core/styles";
import { useTheme } from "@material-ui/core/styles";
import useMediaQuery from "@material-ui/core/useMediaQuery";
import Table from "@material-ui/core/Table";
import TableBody from "@material-ui/core/TableBody";
import TableCell from "@material-ui/core/TableCell";
import TableHead from "@material-ui/core/TableHead";
import TableRow from "@material-ui/core/TableRow";
import Button from "@material-ui/core/Button";
import Loader from "../../../ui/Loader";
import Empty from "../Empty";
import {
  Data,
  Variables,
  query,
  variables,
  getFetchMore
} from "../../../apollo/Query/Subscriptions";
import Row from "./Result";

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
  }
}));

interface Props {
  terms: string[];
}

export default React.memo(function SubscriptionsTable({
  terms
}: Props): JSX.Element | null {
  const classes = useStyles();
  const theme = useTheme();
  const md = useMediaQuery(theme.breakpoints.up("md"));
  const { data, loading, error, fetchMore } = useQuery<Data, Variables>(query, {
    fetchPolicy: "network-only",
    variables: { ...variables, search: { terms } }
  });

  if (loading) {
    return <Loader />;
  }

  if (error) {
    return <Empty>{error.message}</Empty>;
  }

  if (!data) {
    return null;
  }

  const { total, results } = data.subscriptions.subscriptions;
  const showLoadMoreButton = results.length < total;

  if (results.length === 0) {
    return <Empty>We could not find any sources matching your query.</Empty>;
  }

  return (
    <>
      <Table className={classes.table} size="small">
        <TableHead>
          <TableRow>
            <TableCell>Title</TableCell>
            {md && <TableCell>Domain</TableCell>}
            {md && <TableCell>Updated</TableCell>}
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
            onClick={getFetchMore(fetchMore, data, {
              ...variables,
              pagination: {
                ...variables.pagination,
                offset: results.length
              },
              search: { terms }
            })}
          >
            Load more
          </Button>
        )}
      </div>
    </>
  );
});
