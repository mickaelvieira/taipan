import React, { useContext } from "react";
import CardContent from "@material-ui/core/CardContent";
import CardActions from "@material-ui/core/CardActions";
import { Document } from "../../../types/document";
import { BookmarkButton } from "../../ui/Feed/Button";
import Domain from "../../ui/Domain";
import Item from "../../ui/Feed/Item/Item";
import ItemTitle from "../../ui/Feed/Item/Title";
import ItemDescription from "../../ui/Feed/Item/Description";
import ItemImage from "../../ui/Feed/Image";
import ItemFooter from "../../ui/Feed/Item/Footer";
import { MessageContext } from "../../context";

interface Props {
  index: number;
  document: Document;
}

export default React.memo(function FeedItem({ index, document }: Props) {
  const setMessageInfo = useContext(MessageContext);

  return (
    <Item>
      {({ remove }) => (
        <>
          <ItemImage index={index} item={document} />
          <CardContent>
            <ItemTitle item={document} />
            <ItemDescription item={document} />
          </CardContent>
          <ItemFooter>
            <CardActions disableSpacing>
              <Domain item={document} />
            </CardActions>
            <CardActions disableSpacing>
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
