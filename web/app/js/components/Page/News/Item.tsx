import React, { useContext } from "react";
import PropTypes from "prop-types";
import CardContent from "@material-ui/core/CardContent";
import CardActions from "@material-ui/core/CardActions";
import { Document } from "../../../types/document";
import {
  BookmarkButton,
  BookmarkAndFavoriteButton
} from "../../ui/Feed/Button";
import Domain from "../../ui/Domain";
import Item from "../../ui/Feed/Item/Item";
import ItemTitle from "../../ui/Feed/Item/Title";
import ItemDescription from "../../ui/Feed/Item/Description";
import ItemImage from "../../ui/Feed/Image";
import ItemFooter from "../../ui/Feed/Item/Footer";
import { MessageContext } from "../../context";

interface Props {
  document: Document;
  query: PropTypes.Validator<object>;
}

export default React.memo(function FeedItem({
  document,
  query
}: Props): JSX.Element {
  const setMessageInfo = useContext(MessageContext);

  return (
    <Item query={query} item={document}>
      {({ remove }) => (
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
              <BookmarkAndFavoriteButton
                document={document}
                onSuccess={() => {
                  setMessageInfo("The document was added to your favorites");
                  remove();
                }}
              />
              <BookmarkButton
                document={document}
                onSuccess={() => {
                  setMessageInfo("The document was added to your reading list");
                  remove();
                }}
              />
            </CardActions>
          </ItemFooter>
        </>
      )}
    </Item>
  );
});
