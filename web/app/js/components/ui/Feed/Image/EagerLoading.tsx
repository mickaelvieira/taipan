import React from "react";
import { makeStyles } from "@material-ui/core/styles";
import CardMedia, { CardMediaProps } from "@material-ui/core/CardMedia";
import { IMAGE_PLACEHOLDER } from "../../../../constant/image";
import { Image } from "../../../../types/image";

interface Props extends CardMediaProps {
  media: Image | null;
}

const useStyles = makeStyles({
  media: {
    backgroundSize: "cover",
    minHeight: 200
  }
});

export default function EagerLoadingImage({ media, title }: Props) {
  const classes = useStyles();
  const src = media ? media.url : IMAGE_PLACEHOLDER;
  return <CardMedia className={classes.media} image={src} title={title} />;
}
