import React from "react";
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
import MoreVertIcon from "@material-ui/icons/MoreVert";

import { UserBookmark } from "../../../types/bookmark";

const styles = () =>
  createStyles({
    card: {
      marginBottom: 24,
      marginLeft: 12,
      marginRight: 12,
      flex: 1,
      display: "flex",
      flexDirection: "column"
    },
    media: {
      height: 0,
      paddingTop: "56.25%" // 16:9
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

export default withStyles(styles)(function FeedItem({
  bookmark,
  classes
}: Props) {
  const image = bookmark.image.startsWith("https")
    ? bookmark.image
    : "https://placekitten.com/g/400/225";
  return (
    <Card className={classes.card} raised>
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
          <Typography gutterBottom variant="h6" component="h2" noWrap>
            {bookmark.title}
          </Typography>
        </Link>
        <Typography component="p">{bookmark.description}</Typography>
      </CardContent>
      <CardActions className={classes.actions} disableActionSpacing>
        <IconButton aria-label="Add to favorites">
          <FavoriteIcon />
        </IconButton>
        <IconButton aria-label="Share">
          <ShareIcon />
        </IconButton>
        <IconButton>
          <MoreVertIcon />
        </IconButton>
      </CardActions>
    </Card>
  );
});
