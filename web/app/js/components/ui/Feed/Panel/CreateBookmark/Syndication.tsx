import React from "react";
import { makeStyles } from "@material-ui/core/styles";
import Link from "@material-ui/core/Link";
import Checkbox from "@material-ui/core/Checkbox";
import Typography from "@material-ui/core/Typography";
import List from "@material-ui/core/List";
import ListItem from "@material-ui/core/ListItem";
import { Source } from "../../../../../types/syndication";

const useStyles = makeStyles(() => ({
  item: {
    padding: 0
  },
  link: {
    paddingTop: 10,
    paddingLeft: 9,
    paddingBottom: 9
  }
}));

interface Props {
  syndication: Source[];
  subscriptions: string[];
  onChange: (subscriptions: string[]) => void;
}

export default function Syndication({
  syndication,
  subscriptions,
  onChange
}: Props): JSX.Element | null {
  const classes = useStyles();

  return syndication.length === 0 ? null : (
    <>
      <Typography>
        We also found the following RSS feeds. Do you want to subscribe to them?
      </Typography>
      <List>
        {syndication.map(source => (
          <ListItem
            className={classes.item}
            alignItems="flex-start"
            key={source.id}
          >
            <Checkbox
              checked={subscriptions.includes(source.url)}
              onClick={() => {
                const sub = subscriptions.includes(source.url)
                  ? subscriptions.filter(url => url != source.url)
                  : [source.url, ...subscriptions];
                onChange(sub);
              }}
            />
            <Link
              underline="none"
              href={source.url}
              title={source.title}
              target="_blank"
              rel="noopener"
              className={classes.link}
            >
              {source.title && source.title != "wordpress feed"
                ? source.title
                : source.url}
            </Link>
          </ListItem>
        ))}
      </List>
    </>
  );
}
