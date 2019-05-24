import React from "react";
import moment from "moment";
import { makeStyles } from "@material-ui/core/styles";
import Card from "@material-ui/core/Card";
import CardContent from "@material-ui/core/CardContent";
import CardActions from "@material-ui/core/CardActions";
import Link from "@material-ui/core/Link";
import Typography from "@material-ui/core/Typography";
import { Document } from "../../../types/document";
import { truncate } from "../../../helpers/string";
import { EagerLoadingImage, LazyLoadingImage } from "../../ui/Feed/Image";
import { BookmarkButton } from "../../ui/Feed/Button";
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
  document: Document;
}

export default React.memo(function FeedItem({ index, document }: Props) {
  const classes = useStyles();
  const ImageComp = index < 5 ? EagerLoadingImage : LazyLoadingImage;

  return (
    <Card className={classes.card}>
      {document.image && (
        <Link
          underline="none"
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
      )}
      <CardContent className={classes.content}>
        <Link
          underline="none"
          href={document.url}
          title={document.title}
          target="_blank"
          rel="noopener"
        >
          <Typography variant="h6" component="h6">
            {document.title}
          </Typography>
        </Link>
        {document.description && (
          <Typography component="p" gutterBottom>
            {truncate(document.description)}
          </Typography>
        )}
        <Typography variant="body2">
          Created: {moment(document.createdAt).fromNow()}
        </Typography>
        <Typography variant="body2">
          Updated: {moment(document.updatedAt).fromNow()}
        </Typography>
      </CardContent>
      <ItemFooter>
        <CardActions disableSpacing>
          <Domain item={document} />
        </CardActions>
        <CardActions disableSpacing>
          <BookmarkButton document={document} />
        </CardActions>
      </ItemFooter>
    </Card>
  );
});
