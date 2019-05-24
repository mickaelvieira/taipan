import React from "react";
import { Link as RouterLink } from "react-router-dom";
import { makeStyles } from "@material-ui/core/styles";
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
import FavoriteIcon from "@material-ui/icons/Favorite";

const useStyles = makeStyles(({ spacing }) => ({
  list: {
    width: 220
  },
  icon: {
    margin: spacing(1)
  },
  link: {
    display: "block"
  }
}));

interface Props {
  isOpen: boolean;
  toggleDrawer: (status: boolean) => void;
}

export default function Sidebar({ isOpen, toggleDrawer }: Props) {
  const classes = useStyles();
  return (
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
            to="/reading-list"
            className={classes.link}
            component={RouterLink}
            underline="none"
            onClick={() => toggleDrawer(false)}
          >
            <ListItem button key="Reading List">
              <ListItemIcon>
                <BookmarkIcon color="primary" className={classes.icon} />
              </ListItemIcon>
              <ListItemText primary="Reading List" />
            </ListItem>
          </Link>
          <Link
            to="/favorites"
            className={classes.link}
            component={RouterLink}
            underline="none"
            onClick={() => toggleDrawer(false)}
          >
            <ListItem button key="Favorites">
              <ListItemIcon>
                <FavoriteIcon color="primary" className={classes.icon} />
              </ListItemIcon>
              <ListItemText primary="Favorites" />
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
  );
}
