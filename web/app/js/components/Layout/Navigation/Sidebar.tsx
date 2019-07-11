import React from "react";
import { NavLink as RouterLink } from "react-router-dom";
import { useTheme } from "@material-ui/core/styles";
import useMediaQuery from "@material-ui/core/useMediaQuery";
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
import RssFeedIcon from "@material-ui/icons/RssFeedOutlined";
import UserInfo from "./UserInfo";
import AppInfo from "./AppInfo";
import { SIDEBAR_WIDTH } from "../../../constant/sidebar";
import { getSectionTitle } from "../helpers/navigation";

const useStyles = makeStyles(
  ({ breakpoints, spacing, palette, typography }) => ({
    drawer: {
      [breakpoints.up("md")]: {
        width: SIDEBAR_WIDTH,
        flexShrink: 0
      }
    },
    list: {
      width: SIDEBAR_WIDTH
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
  })
);

const entries = [
  {
    path: "/",
    icon: HomeIcon
  },
  {
    path: "/reading-list",
    icon: LibraryIcon
  },
  {
    path: "/favorites",
    icon: FavoriteIcon
  },
  {
    path: "/syndication",
    icon: RssFeedIcon
  }
];

interface Props {
  isOpen: boolean;
  toggleDrawer: (status: boolean) => void;
}

export default function Sidebar({ isOpen, toggleDrawer }: Props): JSX.Element {
  const classes = useStyles();
  const theme = useTheme();
  const matches = useMediaQuery(theme.breakpoints.up("md"));

  return (
    <nav className={classes.drawer}>
      <Drawer
        anchor="left"
        open={isOpen}
        variant={matches ? "permanent" : "temporary"}
        onClose={() => toggleDrawer(false)}
      >
        <UserInfo />
        <List className={classes.list}>
          {entries.map(entry => (
            <Link
              exact
              key={entry.path}
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
                <ListItemText disableTypography>
                  {getSectionTitle(entry.path)}
                </ListItemText>
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
        <AppInfo />
      </Drawer>
    </nav>
  );
}
