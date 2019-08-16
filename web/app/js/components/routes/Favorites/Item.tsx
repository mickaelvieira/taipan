import React, { useContext } from "react";
import CardContent from "@material-ui/core/CardContent";
import CardActions from "@material-ui/core/CardActions";
import { Bookmark } from "../../../types/bookmark";
import {
  UnfavoriteButton,
  UnbookmarkButton,
  ShareButton
} from "../../ui/Feed/Button";
import Domain from "../../ui/Feed/Item/Domain";
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
    <>
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
          <ShareButton
            item={bookmark}
            onSucceed={message => {
              setMessageInfo({ message });
            }}
            onFail={message => setMessageInfo({ message })}
          />
          <UnbookmarkButton
            iconOnly
            bookmark={bookmark}
            onSucceed={({ updateCache, undo }) => {
              setMessageInfo({
                message: "The document was removed from your bookmarks",
                action: undo,
                label: "undo"
              });
              updateCache();
            }}
            onFail={message => setMessageInfo({ message })}
          />
          <UnfavoriteButton
            iconOnly
            bookmark={bookmark}
            onSucceed={({ updateCache, undo }) => {
              setMessageInfo({
                message: "The bookmark was added back to your reading list",
                action: undo,
                label: "undo"
              });
              updateCache();
            }}
            onFail={message => setMessageInfo({ message })}
          />
        </CardActions>
      </ItemFooter>
    </>
  );
});
