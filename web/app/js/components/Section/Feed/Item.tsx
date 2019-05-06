import React from "react";
import moment from "moment";
import { withStyles, WithStyles, createStyles } from "@material-ui/core/styles";
import Card from "@material-ui/core/Card";
import CardMedia from "@material-ui/core/CardMedia";
import CardContent from "@material-ui/core/CardContent";
import CardActions from "@material-ui/core/CardActions";
import IconButton from "@material-ui/core/IconButton";
import Link from "@material-ui/core/Link";
import Typography from "@material-ui/core/Typography";
import FavoriteIcon from "@material-ui/icons/Favorite";
import ShareIcon from "@material-ui/icons/Share";
import CachedIcon from "@material-ui/icons/Cached";
import MoreVertIcon from "@material-ui/icons/MoreVert";
import CircularProgress from "@material-ui/core/CircularProgress";
import { UserBookmark } from "../../../types/bookmark";
import UpdateBookmarkMutation, {
  mutation
} from "../../apollo/Mutation/UpdateBookmark";
import { truncate } from "../../../helpers/string";

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
  bookmark: UserBookmark;
}

export default withStyles(styles)(
  React.memo(function FeedItem({ bookmark, classes }: Props) {
    const image = bookmark.image
      ? bookmark.image.url
      : "https://placekitten.com/g/400/225";
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
          <CardMedia
            className={classes.media}
            image={image}
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
          <IconButton aria-label="Add to favorites">
            <FavoriteIcon />
          </IconButton>
          <IconButton aria-label="Share">
            <ShareIcon />
          </IconButton>
          <UpdateBookmarkMutation mutation={mutation}>
            {(mutate, { loading }) => (
              <IconButton
                aria-label="Share"
                disabled={loading}
                onClick={() =>
                  mutate({
                    variables: { url: bookmark.url }
                  })
                }
              >
                {!loading && <CachedIcon />}
                {loading && <CircularProgress size={16} />}
              </IconButton>
            )}
          </UpdateBookmarkMutation>
          <IconButton>
            <MoreVertIcon />
          </IconButton>
        </CardActions>
      </Card>
    );
  })
);
