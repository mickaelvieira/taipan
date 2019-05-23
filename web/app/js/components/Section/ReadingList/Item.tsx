import React from "react";
import moment from "moment";
import { withStyles, WithStyles, createStyles } from "@material-ui/core/styles";
import Card from "@material-ui/core/Card";
import CardContent from "@material-ui/core/CardContent";
import CardActions from "@material-ui/core/CardActions";
import IconButton from "@material-ui/core/IconButton";
import Link from "@material-ui/core/Link";
import Typography from "@material-ui/core/Typography";
import ShareIcon from "@material-ui/icons/Share";
import CachedIcon from "@material-ui/icons/Cached";
import MoreVertIcon from "@material-ui/icons/MoreVert";
import CircularProgress from "@material-ui/core/CircularProgress";
import { Bookmark } from "../../../types/bookmark";
import BookmarkMutation, { mutation } from "../../apollo/Mutation/Bookmark";
import { truncate } from "../../../helpers/string";
import { EagerLoadingImage, LazyLoadingImage } from "../../ui/Feed/Image";
import { FavoriteButton } from "../../ui/Feed/Button";

const styles = () =>
  createStyles({
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
    },
    actions: {
      display: "flex",
      alignSelf: "flex-end"
    }
  });

interface Props extends WithStyles<typeof styles> {
  index: number;
  bookmark: Bookmark;
}

export default withStyles(styles)(
  React.memo(function FeedItem({ index, bookmark, classes }: Props) {
    const ImageComp = index < 5 ? EagerLoadingImage : LazyLoadingImage;

    return (
      <Card className={classes.card}>
        <Link
          underline="none"
          block
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
        <CardActions className={classes.actions} disableActionSpacing>
          <FavoriteButton bookmark={bookmark} />
        </CardActions>
      </Card>
    );
  })
);
