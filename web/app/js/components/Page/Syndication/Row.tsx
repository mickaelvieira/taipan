import React, { useContext } from "react";
import { makeStyles } from "@material-ui/core/styles";
import Hidden from "@material-ui/core/Hidden";
import TableCell from "@material-ui/core/TableCell";
import TableRow from "@material-ui/core/TableRow";
import Link from "@material-ui/core/Link";
import { Source } from "../../../types/syndication";
import Domain from "../../ui/Domain";
import { StatusButton, DeleteButton } from "../../ui/Syndication/Button";
import SourceTitle from "./Title";
import { MessageContext } from "../../context";

const useStyles = makeStyles({
  link: {
    display: "inline-block",
    padding: "16px 6px 12px",
    position: "relative",
    verticalAlign: "middle",
    boxSizing: "border-box",
    textAlign: "center"
  }
});

interface Props {
  source: Source;
}

export default React.memo(function Row({ source }: Props): JSX.Element {
  const classes = useStyles();
  const setMessageInfo = useContext(MessageContext);
  const { title } = source;
  const url = new URL(source.url);

  return (
    <TableRow>
      <TableCell>
        {title ? <SourceTitle item={source} /> : <Domain item={source} />}
      </TableCell>
      <Hidden mdDown>
        <TableCell>
          <Link
            underline="none"
            href={`${url.protocol}//${url.host}`}
            title={source.title ? source.title : source.url}
            target="_blank"
            rel="noopener"
            className={classes.link}
          >
            {`${url.protocol}//${url.host}`}
          </Link>
        </TableCell>
      </Hidden>
      <TableCell align="center">
        <StatusButton source={source} />
      </TableCell>

      <Hidden mdDown>
        <TableCell align="center">
          <DeleteButton
            source={source}
            onSuccess={() => {
              setMessageInfo("The web syndication source was deleted");
            }}
            onError={message => setMessageInfo(message)}
          />
        </TableCell>
      </Hidden>
    </TableRow>
  );
});
