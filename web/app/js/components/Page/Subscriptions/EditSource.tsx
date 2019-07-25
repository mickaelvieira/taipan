import React from "react";
import { makeStyles } from "@material-ui/core/styles";
import IconButton from "@material-ui/core/IconButton";
import CloseIcon from "@material-ui/icons/Close";
import Divider from "@material-ui/core/Divider";
import Typography from "@material-ui/core/Typography";
import List from "@material-ui/core/List";
import ListItem from "@material-ui/core/ListItem";
import ListItemText from "@material-ui/core/ListItemText";
import ListSubheader from "@material-ui/core/ListSubheader";
import ListItemSecondaryAction from "@material-ui/core/ListItemSecondaryAction";
import SourceQuery from "../../apollo/Query/Source";
import ButtonDeletedStatus from "../../ui/Syndication/Button/DeletedStatus";
import ButtonPausedStatus from "../../ui/Syndication/Button/PausedStatus";
import Panel from "../../ui/Panel";
import Loader from "../../ui/Loader";
import Datetime from "../../ui/Datetime";
import EditTitle from "./EditTitle";
import Link from "./Link";

const useStyles = makeStyles(({ palette, typography }) => ({
  dialog: {},
  header: {
    display: "flex",
    flexDirection: "row",
    justifyContent: "start",
    margin: 0,
    padding: 0,
    backgroundColor: palette.grey[200]
  },
  title: {
    paddingTop: 12,
    paddingBottom: 12
  },
  container: {
    padding: 16,
    display: "flex",
    flexDirection: "column"
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
  isOpen: boolean;
  close: () => void;
}

export default function EditSource({ url, isOpen, close }: Props): JSX.Element {
  const classes = useStyles();

  console.log(url);

  return (
    <Panel isOpen={isOpen} close={close}>
      <header className={classes.header}>
        <IconButton onClick={() => close()}>
          <CloseIcon />
        </IconButton>
        <Typography component="h5" variant="h6" className={classes.title}>
          Edit web syndication source
        </Typography>
      </header>
      <section className={classes.container}>
        <SourceQuery variables={{ url }}>
          {({ data, loading, error }) => {
            console.log(loading);
            console.log(error);

            if (loading) {
              return <Loader />;
            }

            if (!data) {
              return null;
            }

            const {
              syndication: { source }
            } = data;

            return (
              <List className={classes.list}>
                <ListItem>
                  <ListItemText>
                    <EditTitle source={source} />
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
                    <Datetime
                      className={classes.date}
                      value={source.parsedAt}
                    />
                  </ListItemText>
                </ListItem>
                <ListItem>
                  <ListItemText>
                    <span>Created </span>
                    <Datetime
                      className={classes.date}
                      value={source.createdAt}
                    />
                  </ListItemText>
                </ListItem>
                <ListItem>
                  <ListItemText>
                    <span>Updated </span>
                    <Datetime
                      className={classes.date}
                      value={source.updatedAt}
                    />
                  </ListItemText>
                </ListItem>
                <Divider />
                <ListSubheader>Actions</ListSubheader>
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
              </List>
            );
          }}
        </SourceQuery>
      </section>
    </Panel>
  );
}
