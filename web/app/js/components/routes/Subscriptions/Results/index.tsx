import React, { useRef, useEffect } from "react";
import { useQuery } from "@apollo/react-hooks";
import { makeStyles } from "@material-ui/core/styles";
import { useTheme } from "@material-ui/core/styles";
import useMediaQuery from "@material-ui/core/useMediaQuery";
import Table from "@material-ui/core/Table";
import TableBody from "@material-ui/core/TableBody";
import TableCell from "@material-ui/core/TableCell";
import TableHead from "@material-ui/core/TableHead";
import TableRow from "@material-ui/core/TableRow";
import useWindowBottom from "../../../../hooks/useWindowBottom";
import Loader from "../../../ui/Loader";
import Empty from "../Empty";
import {
  Data,
  Variables,
  LoadMore,
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
  tags: string[];
}

export default React.memo(function SubscriptionsTable({
  terms,
  tags
}: Props): JSX.Element | null {
  const isAtTheBottom = useWindowBottom();
  const loadMore = useRef<LoadMore | undefined>();
  const classes = useStyles();
  const theme = useTheme();
  const md = useMediaQuery(theme.breakpoints.up("md"));
  const { data, loading, error, fetchMore } = useQuery<Data, Variables>(query, {
    fetchPolicy: "network-only",
    variables: { ...variables, search: { terms, tags } }
  });

  useEffect(() => {
    if (isAtTheBottom && loadMore.current) {
      loadMore.current();
    }
  }, [isAtTheBottom, loadMore]);

  if (loading) {
    return <Loader />;
  }

  if (error) {
    return <Empty>{error.message}</Empty>;
  }

  if (!data) {
    return null;
  }

  const { results } = data.subscriptions.subscriptions;

  loadMore.current = getFetchMore(fetchMore, data, {
    ...variables,
    pagination: {
      ...variables.pagination,
      offset: results.length
    },
    search: { terms, tags }
  });

  if (results.length === 0) {
    return <Empty>We could not find any sources matching your query.</Empty>;
  }

  return (
    <>
      <Table className={classes.table} size="small">
        <TableHead>
          <TableRow>
            <TableCell>Title</TableCell>
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
    </>
  );
});
