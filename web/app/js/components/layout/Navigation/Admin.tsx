import React, { useState } from "react";
import { makeStyles } from "@material-ui/core/styles";
import Collapse from "@material-ui/core/Collapse";
import List from "@material-ui/core/List";
import ListItem from "@material-ui/core/ListItem";
import ListItemText from "@material-ui/core/ListItemText";
import ExpandLess from "@material-ui/icons/ExpandLess";
import ExpandMore from "@material-ui/icons/ExpandMore";
import SettingsIcon from "@material-ui/icons/Settings";
import RssFeedIcon from "@material-ui/icons/RssFeedSharp";
import MenuLink from "./MenuLink";

const useStyles = makeStyles(({ spacing, palette, breakpoints }) => ({
  nested: {
    [breakpoints.up("md")]: {
      backgroundColor: "#1d1d1d"
    }
  },
  toggle: {
    color: palette.grey[600],
    [breakpoints.up("md")]: {
      color: palette.grey[100]
    },
    "&.active": {
      color: palette.primary.main
    }
  },
  icon: {
    margin: spacing(1),
    marginRight: spacing(3)
  }
}));

interface Props {
  toggleDrawer: (status: boolean) => void;
}

export default function AdminSidebar({ toggleDrawer }: Props): JSX.Element {
  const classes = useStyles();
  const [isOpen, setIsOpen] = useState(false);

  return (
    <li>
      <ListItem
        button
        onClick={() => setIsOpen(!isOpen)}
        className={classes.toggle}
      >
        <SettingsIcon className={classes.icon} />
        <ListItemText disableTypography>Admin</ListItemText>
        {isOpen ? <ExpandLess /> : <ExpandMore />}
      </ListItem>
      <Collapse in={isOpen} timeout="auto" unmountOnExit>
        <List disablePadding>
          <li className={classes.nested}>
            <MenuLink to="/syndication" onClick={() => toggleDrawer(false)}>
              <ListItem button>
                <RssFeedIcon className={classes.icon} />
                <ListItemText disableTypography>Syndication</ListItemText>
              </ListItem>
            </MenuLink>
          </li>
        </List>
      </Collapse>
    </li>
  );
}
