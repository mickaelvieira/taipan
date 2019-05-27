import React from "react";
import CardContent from "@material-ui/core/CardContent";
import CardActions from "@material-ui/core/CardActions";
import { Bookmark } from "../../../types/bookmark";
import {
  FavoriteButton,
  RefreshButton,
  UnbookmarkButton
} from "../../ui/Feed/Button";
import Domain from "../../ui/Domain";
import Item from "../../ui/Feed/Item/Item";
import ItemTitle from "../../ui/Feed/Item/Title";
import ItemDescription from "../../ui/Feed/Item/Description";
import ItemImage from "../../ui/Feed/Image";
import ItemFooter from "../../ui/Feed/Item/Footer";

interface Props {
  index: number;
  bookmark: Bookmark;
}

export default React.memo(function FeedItem({ index, bookmark }: Props) {
  return (
    <Item>
      <ItemImage index={index} item={bookmark} />
      <CardContent>
        <ItemTitle item={bookmark} />
        <ItemDescription item={bookmark} />
      </CardContent>
      <ItemFooter>
        <CardActions disableSpacing>
          <Domain item={bookmark} />
        </CardActions>
        <CardActions disableSpacing>
          <UnbookmarkButton bookmark={bookmark} />
          <FavoriteButton bookmark={bookmark} />
          <RefreshButton bookmark={bookmark} />
        </CardActions>
      </ItemFooter>
    </Item>
  );
});
