import React, { useContext } from "react";
import CardContent from "@material-ui/core/CardContent";
import CardActions from "@material-ui/core/CardActions";
import { Bookmark } from "../../../types/bookmark";
import {
  UnfavoriteButton,
  RefreshButton,
  UnbookmarkButton
} from "../../ui/Feed/Button";
import Domain from "../../ui/Domain";
import Item from "../../ui/Feed/Item/Item";
import ItemTitle from "../../ui/Feed/Item/Title";
import ItemDescription from "../../ui/Feed/Item/Description";
import ItemImage from "../../ui/Feed/Image";
import ItemFooter from "../../ui/Feed/Item/Footer";
import { MessageContext } from "../../context";

interface Props {
  bookmark: Bookmark;
}

export default React.memo(function FeedItem({ bookmark }: Props): JSX.Element {
  const setMessageInfo = useContext(MessageContext);
  return (
    <Item>
      <ItemImage item={bookmark} />
      <CardContent>
        <ItemTitle item={bookmark} />
        <ItemDescription item={bookmark} />
      </CardContent>
      <ItemFooter>
        <CardActions disableSpacing>
          <Domain item={bookmark} />
        </CardActions>
        <CardActions disableSpacing>
          <UnbookmarkButton
            bookmark={bookmark}
            onSuccess={() => {
              setMessageInfo("The document was removed from your bookmarks");
            }}
            onError={message => setMessageInfo(message)}
          />
          <UnfavoriteButton
            bookmark={bookmark}
            onSuccess={() => {
              setMessageInfo(
                "The bookmark was added back to your reading list"
              );
            }}
            onError={message => setMessageInfo(message)}
          />
          <RefreshButton
            bookmark={bookmark}
            onSuccess={() => {}}
            onError={message => setMessageInfo(message)}
          />
        </CardActions>
      </ItemFooter>
    </Item>
  );
});
