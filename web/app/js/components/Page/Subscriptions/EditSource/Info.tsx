import React from "react";
import { makeStyles } from "@material-ui/core/styles";
import List from "@material-ui/core/List";
import ListItem from "@material-ui/core/ListItem";
import ListItemText from "@material-ui/core/ListItemText";
import ListItemSecondaryAction from "@material-ui/core/ListItemSecondaryAction";
import SourceQuery from "../../../apollo/Query/Source";
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

export default React.memo(function Info({ url }: Props): JSX.Element {
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
      }}
    </SourceQuery>
  );
});