import React from "react";
import { useQuery } from "@apollo/react-hooks";
import { makeStyles } from "@material-ui/core/styles";
import Table from "@material-ui/core/Table";
import TableBody from "@material-ui/core/TableBody";
import TableCell from "@material-ui/core/TableCell";
import TableHead from "@material-ui/core/TableHead";
import TableRow from "@material-ui/core/TableRow";
import red from "@material-ui/core/colors/red";
import { Data, Variables, query, variables } from "../../../apollo/Query/Logs";
import Loader from "../../Loader";

const useStyles = makeStyles(({ breakpoints, typography }) => ({
  container: {
    width: "100%",
    overflowX: "hidden",
    overflowY: "auto",
    height: 380,
    maxHeight: 380,
  },
  table: {
    [breakpoints.up("md")]: {
      minWidth: 650,
    },
  },
  error: {
    color: red[500],
    fontWeight: typography.fontWeightBold,
  },
}));

interface Props {
  url: URL;
}

export default React.memo(function Logs({ url }: Props): JSX.Element | null {
  const classes = useStyles();
  const { data, loading, error } = useQuery<Data, Variables>(query, {
    variables: { ...variables, url },
  });

  if (loading) {
    return <Loader />;
  }

  if (error) {
    return <span>{error.message}</span>;
  }

  if (!data) {
    return null;
  }

  const {
    bot: { logs },
  } = data;

  return (
    <div className={classes.container}>
      <Table className={classes.table} size="small">
        <TableHead>
          <TableRow>
            <TableCell align="center">Checksum</TableCell>
            <TableCell align="center">Status Code</TableCell>
            <TableCell align="center">Content Type</TableCell>
            <TableCell align="center">Date</TableCell>
          </TableRow>
        </TableHead>
        <TableBody>
          {logs.results.map((entry) => (
            <TableRow key={entry.id}>
              <TableCell
                align="center"
                className={entry.hasFailed ? classes.error : ""}
              >
                {entry.checksum.substr(0, 6)}
              </TableCell>
              <TableCell
                align="center"
                className={entry.hasFailed ? classes.error : ""}
              >
                {entry.statusCode}
              </TableCell>
              <TableCell
                align="center"
                className={entry.hasFailed ? classes.error : ""}
              >
                {entry.contentType}
              </TableCell>
              <TableCell
                align="center"
                className={entry.hasFailed ? classes.error : ""}
              >
                {entry.createdAt}
              </TableCell>
            </TableRow>
          ))}
        </TableBody>
      </Table>
    </div>
  );
});
