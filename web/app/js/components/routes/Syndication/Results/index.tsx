import React from "react";
import { useQuery } from "@apollo/react-hooks";
import { makeStyles } from "@material-ui/core/styles";
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
} from "../../../apollo/Query/Syndication";
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
  showDeleted: boolean;
  pausedOnly: boolean;
  editSource: (url: URL) => void;
}

export default React.memo(function SyndicationTable({
  terms,
  tags,
  showDeleted,
  pausedOnly,
  editSource
}: Props): JSX.Element | null {
  const classes = useStyles();
  const { data, loading, error, fetchMore } = useQuery<Data, Variables>(query, {
    fetchPolicy: "network-only",
    variables: {
      ...variables,
      search: { terms, tags, pausedOnly, showDeleted }
    }
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

  const { total, results } = data.syndication.sources;
  const showLoadMoreButton = results.length < total;

  if (results.length === 0) {
    return <Empty>We could not find any sources matching your query.</Empty>;
  }

  return (
    <>
      <Table className={classes.table} size="small">
        <TableHead>
          <TableRow>
            <TableCell>Feed</TableCell>
            <TableCell>Tags</TableCell>
            <TableCell>&nbsp;</TableCell>
          </TableRow>
        </TableHead>
        <TableBody>
          {results.map(source => (
            <Row key={source.id} editSource={editSource} source={source} />
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
              search: { terms, tags, pausedOnly, showDeleted }
            })}
          >
            Load more
          </Button>
        )}
      </div>
    </>
  );
});
