import React, { useContext } from "react";
import PropTypes from "prop-types";
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
  query: PropTypes.Validator<object>;
}

export default React.memo(function FeedItem({ bookmark, query }: Props) {
  const setMessageInfo = useContext(MessageContext);
  return (
    <Item query={query} item={bookmark}>
      {({ remove }) => (
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
              <UnbookmarkButton
                bookmark={bookmark}
                onSuccess={() => {
                  setMessageInfo(
                    "The document was removed from your bookmarks"
                  );
                  remove();
                }}
              />
              <UnfavoriteButton
                bookmark={bookmark}
                onSuccess={() => {
                  setMessageInfo(
                    "The bookmark was added back to your reading list"
                  );
                  remove();
                }}
              />
              <RefreshButton bookmark={bookmark} onSuccess={() => {}} />
            </CardActions>
          </ItemFooter>
        </>
      )}
    </Item>
  );
});
