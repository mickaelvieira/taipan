import React, { FunctionComponent } from "react";
import { Link as RouterLink } from "react-router-dom";
import {
  withStyles,
  WithStyles,
  createStyles,
  Theme
} from "@material-ui/core/styles";
import Drawer from "@material-ui/core/Drawer";
import List from "@material-ui/core/List";
import ListItem from "@material-ui/core/ListItem";
import Divider from "@material-ui/core/Divider";
import Link from "@material-ui/core/Link";
import ListItemIcon from "@material-ui/core/ListItemIcon";
import ListItemText from "@material-ui/core/ListItemText";
import BookmarkIcon from "@material-ui/icons/BookmarkBorderOutlined";
import AccountIcon from "@material-ui/icons/AccountCircleRounded";
import HomeIcon from "@material-ui/icons/HomeOutlined";

const styles = ({ spacing }: Theme) =>
  createStyles({
    list: {
      width: 220
    },
    icon: {
      margin: spacing.unit
    },
    link: {
      display: "block"
    }
  });

interface Props extends WithStyles<typeof styles> {
  isOpen: boolean;
  toggleDrawer: (status: boolean) => void;
}

const TemporaryDrawer: FunctionComponent<Props> = ({
  isOpen,
  toggleDrawer,
  classes
}) => (
  <div>
    <Drawer anchor="left" open={isOpen} onClose={() => toggleDrawer(false)}>
      <div className={classes.list}>
        <List>
          <Link
            to="/"
            className={classes.link}
            component={RouterLink}
            underline="none"
            onClick={() => toggleDrawer(false)}
          >
            <ListItem button key="Home">
              <ListItemIcon>
                <HomeIcon color="primary" className={classes.icon} />
              </ListItemIcon>
              <ListItemText primary="Home" />
            </ListItem>
          </Link>
          <Link
            to="/feed"
            className={classes.link}
            component={RouterLink}
            underline="none"
            onClick={() => toggleDrawer(false)}
          >
            <ListItem button key="Latest">
              <ListItemIcon>
                <BookmarkIcon color="primary" className={classes.icon} />
              </ListItemIcon>
              <ListItemText primary="Latest" />
            </ListItem>
          </Link>
        </List>
        <Divider />
        <List>
          <Link
            to="/"
            className={classes.link}
            component={RouterLink}
            underline="none"
            onClick={() => toggleDrawer(false)}
          >
            <ListItem button key="Account">
              <ListItemIcon>
                <AccountIcon color="primary" className={classes.icon} />
              </ListItemIcon>
              <ListItemText primary="Account" />
            </ListItem>
          </Link>
        </List>
      </div>
    </Drawer>
  </div>
);

export default withStyles(styles)(TemporaryDrawer);
