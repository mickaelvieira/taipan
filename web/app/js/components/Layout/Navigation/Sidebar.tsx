import React, { useContext } from "react";
import { NavLink as RouterLink } from "react-router-dom";
import { makeStyles } from "@material-ui/core/styles";
import Drawer from "@material-ui/core/Drawer";
import List from "@material-ui/core/List";
import ListItem from "@material-ui/core/ListItem";
import Divider from "@material-ui/core/Divider";
import Link from "@material-ui/core/Link";
import ListItemText from "@material-ui/core/ListItemText";
import LibraryIcon from "@material-ui/icons/LocalLibraryOutlined";
import AccountIcon from "@material-ui/icons/AccountCircleOutlined";
import HomeIcon from "@material-ui/icons/HomeOutlined";
import FavoriteIcon from "@material-ui/icons/FavoriteBorderOutlined";
import { UserContext } from "../../context";

const useStyles = makeStyles(({ spacing, palette, typography }) => ({
  userInfo: {
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
  appInfo: {
    lineHeight: 1.33,
    color: palette.grey[500],
    textAlign: "center",
    margin: 0,
    padding: "1.2rem 0"
  },
  list: {
    width: 220
  },
  icon: {
    margin: spacing(1),
    marginRight: spacing(3)
  },
  link: {
    display: "block",
    fontWeight: 500,
    fontSize: typography.fontSize,
    color: palette.grey[900],
    "&.active": {
      color: palette.primary.main
    }
  }
}));

const entries = [
  {
    path: "/",
    label: "News",
    icon: HomeIcon
  },
  {
    path: "/reading-list",
    label: "Reading List",
    icon: LibraryIcon
  },
  {
    path: "/favorites",
    label: "Favorites",
    icon: FavoriteIcon
  }
];

interface Props {
  isOpen: boolean;
  toggleDrawer: (status: boolean) => void;
}

export default function Sidebar({ isOpen, toggleDrawer }: Props): JSX.Element {
  const user = useContext(UserContext);
  const classes = useStyles();
  return (
    <Drawer anchor="left" open={isOpen} onClose={() => toggleDrawer(false)}>
      {user && (
        <p className={classes.userInfo}>
          {user.firstname} {user.lastname}
        </p>
      )}
      <List className={classes.list}>
        {entries.map(entry => (
          <Link
            exact
            key={entry.label}
            to={entry.path}
            classes={{
              root: classes.link
            }}
            component={RouterLink}
            underline="none"
            onClick={() => toggleDrawer(false)}
          >
            <ListItem button>
              <entry.icon className={classes.icon} />
              <ListItemText disableTypography>{entry.label}</ListItemText>
            </ListItem>
          </Link>
        ))}
      </List>
      <Divider />
      <List>
        <Link
          to="/account"
          exact
          classes={{
            root: classes.link
          }}
          component={RouterLink}
          underline="none"
          onClick={() => toggleDrawer(false)}
        >
          <ListItem button key="Account">
            <AccountIcon className={classes.icon} />
            <ListItemText disableTypography>Account</ListItemText>
          </ListItem>
        </Link>
      </List>
      <Divider />
      <p className={classes.appInfo}>Taipan {process.env.APP_VERSION}</p>
    </Drawer>
  );
}
