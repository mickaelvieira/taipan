import React, { useContext } from "react";
import ListItem from "@material-ui/core/ListItem";
import ListItemIcon from "@material-ui/core/ListItemIcon";
import ListItemSecondaryAction from "@material-ui/core/ListItemSecondaryAction";
import ListItemText from "@material-ui/core/ListItemText";
import RssIcon from "@material-ui/icons/RssFeedOutlined";
import { Source } from "../../../types/syndication";
import Domain from "../../ui/Domain";
import { ToggleStatusButton, DeleteButton } from "../../ui/Syndication/Button";
import SourceTitle from "./Title";
import { MessageContext } from "../../context";

interface Props {
  source: Source;
}

export default React.memo(function SyndicationSource({
  source
}: Props): JSX.Element {
  const setMessageInfo = useContext(MessageContext);
  const { title } = source;

  return (
    <ListItem>
      <ListItemIcon>
        <RssIcon />
      </ListItemIcon>
      <ListItemText>
        {title ? <SourceTitle item={source} /> : <Domain item={source} />}
      </ListItemText>
      <ListItemSecondaryAction>
        <ToggleStatusButton source={source} />
        <DeleteButton
          source={source}
          onSuccess={() => {
            setMessageInfo("The web syndication source was deleted");
          }}
        />
      </ListItemSecondaryAction>
    </ListItem>
  );
});
