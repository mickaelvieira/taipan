import React, { useContext } from "react";
import CardContent from "@material-ui/core/CardContent";
import CardActions from "@material-ui/core/CardActions";
import { Document } from "../../../../types/document";
import {
  ShareButton,
  BookmarkButton,
  BookmarkAndFavoriteButton,
} from "../Button";
import Domain from "./Domain";
import ItemTitle from "./Title";
import ItemDescription from "./Description";
import ItemImage from "../Image";
import ItemFooter from "./Footer";
import { MessageContext } from "../../../context";

interface Props {
  document: Document;
}

export default React.memo(function FeedItemDocument({
  document,
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
            onSucceed={(message) => {
              setMessageInfo({ message });
            }}
            onFail={(message) => setMessageInfo({ message })}
          />
          <BookmarkAndFavoriteButton
            iconOnly
            document={document}
            onSucceed={({ updateCache, undo }) => {
              setMessageInfo({
                message: "The document was added to your favorites",
                action: undo,
                label: "undo",
              });
              updateCache();
            }}
            onFail={(message) => setMessageInfo({ message })}
          />
          <BookmarkButton
            iconOnly
            document={document}
            onSucceed={({ updateCache, undo }) => {
              setMessageInfo({
                message: "The document was added to your reading list",
                action: undo,
                label: "undo",
              });
              updateCache();
            }}
            onFail={(message) => setMessageInfo({ message })}
          />
        </CardActions>
      </ItemFooter>
    </>
  );
});
