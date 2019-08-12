import React from "react";
import { makeStyles } from "@material-ui/core/styles";
import CardMedia from "@material-ui/core/CardMedia";
import { IMAGE_PLACEHOLDER } from "../../../../constant/image";
import { Image } from "../../../../types/image";
import { getImageWidth } from "../../../../helpers/image";

interface Props {
  title: string;
  media: Image | null;
}

const useStyles = makeStyles(({ breakpoints }) => ({
  media: {
    backgroundSize: "cover",
    minHeight: `calc(${getImageWidth("sm")}px * 9 / 16)`,
    [breakpoints.up("md")]: {
      minHeight: `calc(${getImageWidth("sm")}px * 9 / 16)`
    }
  }
}));

export default function EagerLoadingImage({
  media,
  title
}: Props): JSX.Element {
  const classes = useStyles();
  const src = media ? `${media.url}` : IMAGE_PLACEHOLDER;
  return <CardMedia className={classes.media} image={src} title={title} />;
}
