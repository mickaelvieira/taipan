import React, { useContext } from "react";
import { useTheme } from "@material-ui/core/styles";
import useMediaQuery from "@material-ui/core/useMediaQuery";
import { makeStyles } from "@material-ui/core/styles";
import Drawer from "@material-ui/core/Drawer";
import List from "@material-ui/core/List";
import ListItem from "@material-ui/core/ListItem";
import Divider from "@material-ui/core/Divider";
import ListItemText from "@material-ui/core/ListItemText";
import LibraryIcon from "@material-ui/icons/LocalLibrarySharp";
import ExitIcon from "@material-ui/icons/ExitToApp";
import AccountIcon from "@material-ui/icons/AccountCircleSharp";
import HomeIcon from "@material-ui/icons/HomeSharp";
import FavoriteIcon from "@material-ui/icons/FavoriteSharp";
import RssFeedIcon from "@material-ui/icons/RssFeedSharp";
import UserInfo from "./UserInfo";
import AppInfo from "./AppInfo";
import { SIDEBAR_WIDTH } from "../../../constant/sidebar";
import { getSectionTitle } from "../../../helpers/navigation";
import { logout } from "../../../helpers/app";
import MenuLink from "./MenuLink";
import Admin from "./Admin";
import { UserContext } from "../../context";
import { isAdmin } from "../../../helpers/users";

const useStyles = makeStyles(
  ({ breakpoints, spacing, palette, typography }) => ({
    drawer: {
      [breakpoints.up("md")]: {
        width: SIDEBAR_WIDTH,
        flexShrink: 0,
        backgroundColor: palette.grey[500],
      },
    },
    paper: {
      [breakpoints.up("md")]: {
        maxWidth: SIDEBAR_WIDTH + 1, // +1 for the border
        backgroundColor: palette.grey[900],
      },
    },
    divider: {
      margin: "0 1rem",
    },
    list: {
      width: SIDEBAR_WIDTH,
    },
    icon: {
      margin: spacing(1),
      marginRight: spacing(3),
    },
    link: {
      display: "block",
      fontSize: typography.fontSize,
      color: palette.grey[600],
      [breakpoints.up("md")]: {
        color: palette.grey[100],
      },
      "&.active": {
        color: palette.primary.main,
      },
    },
  })
);

const entries = [
  {
    path: "/",
    icon: HomeIcon,
  },
  {
    path: "/reading-list",
    icon: LibraryIcon,
  },
  {
    path: "/favorites",
    icon: FavoriteIcon,
  },
  {
    path: "/subscriptions",
    icon: RssFeedIcon,
  },
];

interface Props {
  isOpen: boolean;
  toggleDrawer: (status: boolean) => void;
}

export default function Sidebar({ isOpen, toggleDrawer }: Props): JSX.Element {
  const user = useContext(UserContext);
  const classes = useStyles();
  const theme = useTheme();
  const md = useMediaQuery(theme.breakpoints.up("md"));

  return (
    <nav className={classes.drawer}>
      <Drawer
        anchor="left"
        open={isOpen}
        variant={md ? "permanent" : "temporary"}
        onClose={() => toggleDrawer(false)}
        classes={{
          paper: classes.paper,
        }}
      >
        <UserInfo />
        <Divider className={classes.divider} />
        <List className={classes.list}>
          {entries.map((entry) => (
            <li key={entry.path}>
              <MenuLink to={entry.path} onClick={() => toggleDrawer(false)}>
                <ListItem button>
                  <entry.icon className={classes.icon} />
                  <ListItemText disableTypography>
                    {getSectionTitle(entry.path)}
                  </ListItemText>
                </ListItem>
              </MenuLink>
            </li>
          ))}
        </List>
        <Divider className={classes.divider} />
        <List>
          <li>
            <MenuLink to="/account" onClick={() => toggleDrawer(false)}>
              <ListItem button>
                <AccountIcon className={classes.icon} />
                <ListItemText disableTypography>Account</ListItemText>
              </ListItem>
            </MenuLink>
          </li>
          {user && isAdmin(user) && <Admin toggleDrawer={toggleDrawer} />}
          <li>
            <MenuLink
              to="/signin"
              onClick={(event: React.MouseEvent) => {
                event.preventDefault();
                logout()
                  .then(({ error, result }) => {
                    if (error) {
                      console.warn(error.error);
                    } else if (result) {
                      window.location.href = "/signin";
                    }
                  })
                  .catch((e) => {
                    console.warn(e.message);
                  });
              }}
            >
              <ListItem button>
                <ExitIcon className={classes.icon} />
                <ListItemText disableTypography>Sign out</ListItemText>
              </ListItem>
            </MenuLink>
          </li>
        </List>
        <Divider className={classes.divider} />
        <AppInfo />
      </Drawer>
    </nav>
  );
}
