import React from "react";
import { makeStyles } from "@material-ui/core/styles";
import Table from "@material-ui/core/Table";
import TableBody from "@material-ui/core/TableBody";
import TableCell from "@material-ui/core/TableCell";
import TableHead from "@material-ui/core/TableHead";
import TableRow from "@material-ui/core/TableRow";
import red from "@material-ui/core/colors/red";
import SourceQuery from "../../../apollo/Query/Source";
import Loader from "../../../ui/Loader";

const useStyles = makeStyles(({ typography }) => ({
  logs: {
    width: "100%",
    overflowX: "hidden",
    overflowY: "auto",
    maxHeight: 360
  },
  table: {
    minWidth: 650
  },
  error: {
    color: red[500],
    fontWeight: typography.fontWeightBold
  }
}));

interface Props {
  url: string;
}

export default React.memo(function Logs({ url }: Props): JSX.Element {
  const classes = useStyles();

  return (
    <SourceQuery variables={{ url }}>
      {({ data, loading, error }) => {
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
          syndication: { source }
        } = data;

        return (
          <div className={classes.logs}>
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
                {source.logEntries.map(entry => (
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
      }}
    </SourceQuery>
  );
});
