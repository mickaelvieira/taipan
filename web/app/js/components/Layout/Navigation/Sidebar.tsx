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
import AccountIcon from "@material-ui/icons/AccountCircleOutlined";
import HomeIcon from "@material-ui/icons/HomeOutlined";
import FavoriteIcon from "@material-ui/icons/Favorite";
import UserQuery, { query } from "../../apollo/Query/User";

const useStyles = makeStyles(({ spacing, palette, typography }) => ({
  user: {
    fontSize: "1.2rem",
    fontWeight: 500,
    lineHeight: 1.33,
    letterSpacing: "0em",
    color: palette.grey[900],
    textAlign: "center",
    margin: 0,
    padding: "1.2rem 0",
    backgroundColor: palette.grey[100]
  },
  list: {
    width: 220
  },
  icon: {
    margin: spacing(1),
    color: palette.grey[900]
  },
  link: {
    display: "block",
    fontWeight: 500,
    fontSize: typography.fontSize,
    color: palette.grey[900]
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
      <UserQuery query={query}>
        {({ data }) => {
          return !data || !data.User ? null : (
            <p className={classes.user}>
              {data.User.firstname} {data.User.lastname}
            </p>
          );
        }}
      </UserQuery>
      <div className={classes.list}>
        <List>
          <Link
            to="/"
            classes={{
              root: classes.link
            }}
            component={RouterLink}
            underline="none"
            onClick={() => toggleDrawer(false)}
          >
            <ListItem button key="Home">
              <ListItemIcon>
                <HomeIcon className={classes.icon} />
              </ListItemIcon>
              <ListItemText disableTypography>Home</ListItemText>
            </ListItem>
          </Link>
          <Link
            to="/reading-list"
            classes={{
              root: classes.link
            }}
            component={RouterLink}
            underline="none"
            onClick={() => toggleDrawer(false)}
          >
            <ListItem button key="Reading List">
              <ListItemIcon>
                <BookmarkIcon className={classes.icon} />
              </ListItemIcon>
              <ListItemText disableTypography>Reading List</ListItemText>
            </ListItem>
          </Link>
          <Link
            to="/favorites"
            classes={{
              root: classes.link
            }}
            component={RouterLink}
            underline="none"
            onClick={() => toggleDrawer(false)}
          >
            <ListItem button key="Favorites">
              <ListItemIcon>
                <FavoriteIcon className={classes.icon} />
              </ListItemIcon>
              <ListItemText disableTypography>Favorites</ListItemText>
            </ListItem>
          </Link>
        </List>
        <Divider />
        <List>
          <Link
            to="/"
            classes={{
              root: classes.link
            }}
            component={RouterLink}
            underline="none"
            onClick={() => toggleDrawer(false)}
          >
            <ListItem button key="Account">
              <ListItemIcon>
                <AccountIcon className={classes.icon} />
              </ListItemIcon>
              <ListItemText disableTypography>Account</ListItemText>
            </ListItem>
          </Link>
        </List>
      </div>
    </Drawer>
  );
}
