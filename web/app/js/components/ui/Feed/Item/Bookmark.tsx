import React, { useContext } from "react";
import CardContent from "@material-ui/core/CardContent";
import CardActions from "@material-ui/core/CardActions";
import { Bookmark } from "../../../../types/bookmark";
import {
  FavoriteButton,
  UnfavoriteButton,
  UnbookmarkButton,
  ShareButton
} from "../Button";
import Domain from "./Domain";
import ItemTitle from "./Title";
import ItemDescription from "./Description";
import ItemImage from "../Image";
import ItemFooter from "./Footer";
import { MessageContext } from "../../../context";

interface Props {
  bookmark: Bookmark;
}

export default React.memo(function FeedItemBookmark({
  bookmark
}: Props): JSX.Element {
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
          {bookmark.isFavorite && (
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
          )}
          {!bookmark.isFavorite && (
            <FavoriteButton
              iconOnly
              bookmark={bookmark}
              onSucceed={({ updateCache, undo }) => {
                setMessageInfo({
                  message: "The bookmark was added to your favorites",
                  action: undo,
                  label: "undo"
                });
                updateCache();
              }}
              onFail={message => setMessageInfo({ message })}
            />
          )}
        </CardActions>
      </ItemFooter>
    </>
  );
});
