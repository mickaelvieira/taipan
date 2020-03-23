import React from "react";
import { makeStyles } from "@material-ui/core/styles";
import Checkbox from "@material-ui/core/Checkbox";
import Typography from "@material-ui/core/Typography";
import List from "@material-ui/core/List";
import ListItem from "@material-ui/core/ListItem";
import InputLabel from "@material-ui/core/InputLabel";
import { Source } from "../../../../types/syndication";

const useStyles = makeStyles(() => ({
  item: {
    padding: 0,
  },
  label: {
    padding: "14px 8px",
  },
}));

interface Props {
  syndication: Source[];
  subscriptions: string[];
  onChange: (subscriptions: string[]) => void;
}

export default function Syndication({
  syndication,
  subscriptions,
  onChange,
}: Props): JSX.Element | null {
  const classes = useStyles();

  return syndication.length === 0 ? null : (
    <>
      <Typography>
        We also found the following RSS feeds. Do you want to subscribe to them?
      </Typography>
      <List>
        {syndication.map((source, index) => (
          <ListItem
            className={classes.item}
            alignItems="flex-start"
            key={source.id}
          >
            <Checkbox
              id={`feed-${index}`}
              checked={subscriptions.includes(`${source.url}`)}
              onClick={() => {
                const sub = subscriptions.includes(`${source.url}`)
                  ? subscriptions.filter((url) => url != `${source.url}`)
                  : [`${source.url}`, ...subscriptions];
                onChange(sub);
              }}
            />
            <InputLabel htmlFor={`feed-${index}`} className={classes.label}>
              {source.title && source.title != "wordpress feed"
                ? source.title
                : `${source.url}`}
            </InputLabel>
          </ListItem>
        ))}
      </List>
    </>
  );
}
