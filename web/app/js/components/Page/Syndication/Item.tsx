import React, { useContext } from "react";
import { makeStyles } from "@material-ui/core/styles";
import Link from "@material-ui/core/Link";
import ListItem from "@material-ui/core/ListItem";
import ListItemIcon from "@material-ui/core/ListItemIcon";
import ListItemSecondaryAction from "@material-ui/core/ListItemSecondaryAction";
import ListItemText from "@material-ui/core/ListItemText";
import RssIcon from "@material-ui/icons/RssFeedOutlined";
import LinkIcon from "@material-ui/icons/OpenInBrowserOutlined";
import { Source } from "../../../types/syndication";
import Domain from "../../ui/Domain";
import { ToggleStatusButton, DeleteButton } from "../../ui/Syndication/Button";
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

export default React.memo(function SyndicationSource({
  source
}: Props): JSX.Element {
  const classes = useStyles();
  const setMessageInfo = useContext(MessageContext);
  const { title } = source;
  const url = new URL(source.url);

  return (
    <ListItem>
      <ListItemIcon>
        <RssIcon />
      </ListItemIcon>
      <ListItemText>
        {title ? <SourceTitle item={source} /> : <Domain item={source} />}
      </ListItemText>
      <ListItemSecondaryAction>
        <Link
          underline="none"
          href={`${url.protocol}//${url.host}`}
          title={source.title ? source.title : source.url}
          target="_blank"
          rel="noopener"
          className={classes.link}
        >
          <LinkIcon />
        </Link>
        <ToggleStatusButton source={source} />
        <DeleteButton
          source={source}
          onSuccess={() => {
            setMessageInfo("The web syndication source was deleted");
          }}
        />
      </ListItemSecondaryAction>
    </ListItem>
  );
});
