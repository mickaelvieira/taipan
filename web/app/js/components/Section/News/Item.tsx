import React from "react";
import moment from "moment";
import { withStyles, WithStyles, createStyles } from "@material-ui/core/styles";
import Card from "@material-ui/core/Card";
import CardContent from "@material-ui/core/CardContent";
import CardActions from "@material-ui/core/CardActions";
import IconButton from "@material-ui/core/IconButton";
import Link from "@material-ui/core/Link";
import Typography from "@material-ui/core/Typography";
import { Document } from "../../../types/document";
import { truncate } from "../../../helpers/string";
import { EagerLoadingImage, LazyLoadingImage } from "../../ui/Feed/Image";
import { BookmarkButton } from "../../ui/Feed/Button";

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
  document: Document;
}

export default withStyles(styles)(
  React.memo(function FeedItem({ index, document, classes }: Props) {
    const ImageComp = index < 5 ? EagerLoadingImage : LazyLoadingImage;

    return (
      <Card className={classes.card}>
        <Link
          underline="none"
          block
          href={document.url}
          title={document.title}
          target="_blank"
          rel="noopener"
        >
          <ImageComp
            className={classes.media}
            media={document.image}
            title={document.title}
          />
        </Link>
        <CardContent className={classes.content}>
          <Link
            underline="none"
            href={document.url}
            title={document.title}
            target="_blank"
            rel="noopener"
          >
            <Typography gutterBottom variant="h6" component="h2">
              {document.title}
            </Typography>
          </Link>
          <Typography component="p" gutterBottom>
            {truncate(document.description)}
          </Typography>
          <Typography variant="body2">
            Created: {moment(document.createdAt).fromNow()}
          </Typography>
          <Typography variant="body2">
            Updated: {moment(document.updatedAt).fromNow()}
          </Typography>
        </CardContent>
        <CardActions className={classes.actions} disableActionSpacing>
          <BookmarkButton document={document} />
        </CardActions>
      </Card>
    );
  })
);
