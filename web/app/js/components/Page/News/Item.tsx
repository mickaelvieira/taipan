import React, { useContext } from "react";
import CardContent from "@material-ui/core/CardContent";
import CardActions from "@material-ui/core/CardActions";
import { Document } from "../../../types/document";
import {
  ShareButton,
  BookmarkButton,
  BookmarkAndFavoriteButton
} from "../../ui/Feed/Button";
import Domain from "../../ui/Domain";
import ItemTitle from "../../ui/Feed/Item/Title";
import ItemDescription from "../../ui/Feed/Item/Description";
import ItemImage from "../../ui/Feed/Image";
import ItemFooter from "../../ui/Feed/Item/Footer";
import { MessageContext } from "../../context";
import { CacheUpdater } from "../../../types";

interface Props {
  remove: (cb: CacheUpdater) => void;
  document: Document;
}

export default React.memo(function FeedItem({
  document,
  remove
}: Props): JSX.Element {
  const setMessageInfo = useContext(MessageContext);

  return (
    <>
      <ItemImage item={document} />
      <CardContent>
        <ItemTitle item={document} />
        <ItemDescription item={document} />
      </CardContent>
      <ItemFooter>
        <CardActions disableSpacing>
          <Domain item={document} />
        </CardActions>
        <CardActions disableSpacing>
          <ShareButton
            item={document}
            onSuccess={message => {
              setMessageInfo({ message });
            }}
            onError={message => setMessageInfo({ message })}
          />
          <BookmarkAndFavoriteButton
            document={document}
            onSuccess={(update, undo) => {
              setMessageInfo({
                message: "The document was added to your favorites",
                action: undo,
                label: "undo"
              });
              remove(update);
            }}
            onError={message => setMessageInfo({ message })}
          />
          <BookmarkButton
            document={document}
            onSuccess={(update, undo) => {
              setMessageInfo({
                message: "The document was added to your reading list",
                action: undo,
                label: "undo"
              });
              remove(update);
            }}
            onError={message => setMessageInfo({ message })}
          />
        </CardActions>
      </ItemFooter>
    </>
  );
});
