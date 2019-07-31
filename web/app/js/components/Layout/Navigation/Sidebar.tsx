import React from "react";
import { NavLink as RouterLink, LinkProps as RouterLinkProps } from "react-router-dom";
import { useTheme } from "@material-ui/core/styles";
import useMediaQuery from "@material-ui/core/useMediaQuery";
import { makeStyles } from "@material-ui/core/styles";
import Drawer from "@material-ui/core/Drawer";
import List from "@material-ui/core/List";
import ListItem from "@material-ui/core/ListItem";
import Divider from "@material-ui/core/Divider";
import Link from "@material-ui/core/Link";
import ListItemText from "@material-ui/core/ListItemText";
import LibraryIcon from "@material-ui/icons/LocalLibrarySharp";
import AccountIcon from "@material-ui/icons/AccountCircleSharp";
import HomeIcon from "@material-ui/icons/HomeSharp";
import FavoriteIcon from "@material-ui/icons/FavoriteSharp";
import RssFeedIcon from "@material-ui/icons/RssFeedSharp";
import UserInfo from "./UserInfo";
import AppInfo from "./AppInfo";
import { SIDEBAR_WIDTH } from "../../../constant/sidebar";
import { getSectionTitle } from "../helpers/navigation";

const AdapterLink = React.forwardRef<HTMLAnchorElement, RouterLinkProps>((props, ref) => (
  <RouterLink exact innerRef={ref as any} {...props} />
));

const useStyles = makeStyles(
  ({ breakpoints, spacing, palette, typography }) => ({
    drawer: {
      [breakpoints.up("md")]: {
        width: SIDEBAR_WIDTH,
        flexShrink: 0,
        backgroundColor: palette.grey[500]
      }
    },
    paper: {
      [breakpoints.up("md")]: {
        backgroundColor: palette.grey[900]
      }
    },
    divider: {
      margin: "0 1rem"
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
      fontSize: typography.fontSize,
      color: palette.grey[600],
      [breakpoints.up("md")]: {
        color: palette.grey[100]
      },
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
  const md = useMediaQuery(theme.breakpoints.up("md"));

  return (
    <nav className={classes.drawer}>
      <Drawer
        anchor="left"
        open={isOpen}
        variant={md ? "permanent" : "temporary"}
        onClose={() => toggleDrawer(false)}
        classes={{
          paper: classes.paper
        }}
      >
        <UserInfo />
        <Divider className={classes.divider} />
        <List className={classes.list}>
          {entries.map(entry => (
            <Link
              key={entry.path}
              to={entry.path}
              classes={{
                root: classes.link
              }}
              component={AdapterLink}
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
        <Divider className={classes.divider} />
        <List>
          <Link
            to="/account"
            classes={{
              root: classes.link
            }}
            component={AdapterLink}
            underline="none"
            onClick={() => toggleDrawer(false)}
          >
            <ListItem button key="Account">
              <AccountIcon className={classes.icon} />
              <ListItemText disableTypography>Account</ListItemText>
            </ListItem>
          </Link>
        </List>
        <Divider className={classes.divider} />
        <AppInfo />
      </Drawer>
    </nav>
  );
}
