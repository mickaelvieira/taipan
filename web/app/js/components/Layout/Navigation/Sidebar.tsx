import React, { FunctionComponent } from "react";
import { WithStyles, createStyles, Theme } from "@material-ui/core";
import { withStyles } from "@material-ui/core/styles";
import Drawer from "@material-ui/core/Drawer";
import List from "@material-ui/core/List";
import Divider from "@material-ui/core/Divider";
import ListItem from "@material-ui/core/ListItem";
import ListItemIcon from "@material-ui/core/ListItemIcon";
import ListItemText from "@material-ui/core/ListItemText";
import InboxIcon from "@material-ui/icons/MoveToInbox";
import MailIcon from "@material-ui/icons/Mail";

const styles = () =>
  createStyles({
    list: {
      width: 250
    },
    fullList: {
      width: "auto"
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
      <div
        tabIndex={0}
        role="button"
        onClick={() => toggleDrawer(false)}
        onKeyDown={() => toggleDrawer(false)}
      >
        <div className={classes.list}>
          <List>
            {["Inbox", "Starred", "Send email", "Drafts"].map((text, index) => (
              <ListItem button key={text}>
                <ListItemIcon>
                  {index % 2 === 0 ? <InboxIcon /> : <MailIcon />}
                </ListItemIcon>
                <ListItemText primary={text} />
              </ListItem>
            ))}
          </List>
          <Divider />
          <List>
            {["All mail", "Trash", "Spam"].map((text, index) => (
              <ListItem button key={text}>
                <ListItemIcon>
                  {index % 2 === 0 ? <InboxIcon /> : <MailIcon />}
                </ListItemIcon>
                <ListItemText primary={text} />
              </ListItem>
            ))}
          </List>
        </div>
      </div>
    </Drawer>
  </div>
);

export default withStyles(styles)(TemporaryDrawer);
