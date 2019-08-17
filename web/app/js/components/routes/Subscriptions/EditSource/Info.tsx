import React from "react";
import { useQuery } from "@apollo/react-hooks";
import { makeStyles } from "@material-ui/core/styles";
import List from "@material-ui/core/List";
import ListItem from "@material-ui/core/ListItem";
import ListItemText from "@material-ui/core/ListItemText";
import ListItemSecondaryAction from "@material-ui/core/ListItemSecondaryAction";
import { query, Data, Variables } from "../../../apollo/Query/Source";
import ButtonDeletedStatus from "../../../ui/Syndication/Button/DeletedStatus";
import ButtonPausedStatus from "../../../ui/Syndication/Button/PausedStatus";
import Loader from "../../../ui/Loader";
import Datetime from "../../../ui/Datetime";
import EditTitle from "./EditTitle";
import Domain from "../Domain";
import Link from "../Link";

const useStyles = makeStyles(({ typography }) => ({
  title: {
    paddingTop: 12,
    paddingBottom: 12
  },
  list: {
    width: "50%"
  },
  date: {
    fontWeight: typography.fontWeightMedium
  }
}));

interface Props {
  url: string;
}

export default React.memo(function Info({ url }: Props): JSX.Element | null {
  const classes = useStyles();
  const { data, loading, error } = useQuery<Data, Variables>(query, {
    variables: { url }
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
    syndication: { source }
  } = data;

  if (!source) {
    return null;
  }

  return (
    <List className={classes.list} dense>
      <ListItem>
        <ListItemText>
          <EditTitle source={source} />
        </ListItemText>
      </ListItem>
      <ListItem>
        <ListItemText>Status</ListItemText>
        <ListItemSecondaryAction>
          <ButtonPausedStatus source={source} />
        </ListItemSecondaryAction>
      </ListItem>
      <ListItem>
        <ListItemText>Visibility</ListItemText>
        <ListItemSecondaryAction>
          <ButtonDeletedStatus source={source} />
        </ListItemSecondaryAction>
      </ListItem>
      <ListItem>
        <ListItemText>
          <Domain item={source} />
        </ListItemText>
      </ListItem>
      <ListItem>
        <ListItemText>
          <Link item={source} />
        </ListItemText>
      </ListItem>
      <ListItem>
        <ListItemText>{source.type}</ListItemText>
      </ListItem>
      <ListItem>
        <ListItemText>Updated {source.frequency}</ListItemText>
      </ListItem>
      <ListItem>
        <ListItemText>
          <span>Last parsed </span>
          {source.parsedAt && (
            <Datetime className={classes.date} value={source.parsedAt} />
          )}
        </ListItemText>
      </ListItem>
      <ListItem>
        <ListItemText>
          <span>Created </span>
          <Datetime className={classes.date} value={source.createdAt} />
        </ListItemText>
      </ListItem>
      <ListItem>
        <ListItemText>
          <span>Updated </span>
          <Datetime className={classes.date} value={source.updatedAt} />
        </ListItemText>
      </ListItem>
    </List>
  );
});
