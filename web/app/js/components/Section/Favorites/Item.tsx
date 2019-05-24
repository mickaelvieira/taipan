import React from "react";
import moment from "moment";
import { makeStyles } from "@material-ui/core/styles";
import Card from "@material-ui/core/Card";
import CardContent from "@material-ui/core/CardContent";
import CardActions from "@material-ui/core/CardActions";
import IconButton from "@material-ui/core/IconButton";
import Link from "@material-ui/core/Link";
import Typography from "@material-ui/core/Typography";
import ShareIcon from "@material-ui/icons/Share";
import { Bookmark } from "../../../types/bookmark";
import { truncate } from "../../../helpers/string";
import { EagerLoadingImage, LazyLoadingImage } from "../../ui/Feed/Image";
import { FavoriteButton, RefreshButton } from "../../ui/Feed/Button";
import Domain from "../../ui/Domain";
import ItemFooter from "../../ui/Feed/Item/Footer";

const useStyles = makeStyles({
  card: {
    marginBottom: 24,
    display: "flex",
    flexDirection: "column"
  },
  media: {
    backgroundSize: "cover",
    minHeight: 200
  },
  content: {
    flex: 1
  }
});

interface Props {
  index: number;
  bookmark: Bookmark;
}

export default React.memo(function FeedItem({ index, bookmark }: Props) {
  const classes = useStyles();
  const ImageComp = index < 5 ? EagerLoadingImage : LazyLoadingImage;

  return (
    <Card className={classes.card}>
      <Link
        underline="none"
        href={bookmark.url}
        title={bookmark.title}
        target="_blank"
        rel="noopener"
      >
        <ImageComp
          className={classes.media}
          media={bookmark.image}
          title={bookmark.title}
        />
      </Link>
      <CardContent className={classes.content}>
        <Link
          underline="none"
          href={bookmark.url}
          title={bookmark.title}
          target="_blank"
          rel="noopener"
        >
          <Typography gutterBottom variant="h6" component="h2">
            {bookmark.title}
          </Typography>
        </Link>
        <Typography component="p" gutterBottom>
          {truncate(bookmark.description)}
        </Typography>
        <Typography variant="body2">
          Added: {moment(bookmark.addedAt).fromNow()}
        </Typography>
        <Typography variant="body2">
          Updated: {moment(bookmark.updatedAt).fromNow()}
        </Typography>
      </CardContent>
      <ItemFooter>
        <CardActions disableSpacing>
          <Domain item={bookmark} />
        </CardActions>
        <CardActions disableSpacing>
          <FavoriteButton bookmark={bookmark} />
          <IconButton aria-label="Share">
            <ShareIcon />
          </IconButton>
          <RefreshButton bookmmark={bookmark} />
        </CardActions>
      </ItemFooter>
    </Card>
  );
});
