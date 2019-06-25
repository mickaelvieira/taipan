import React, { useRef } from "react";
import { makeStyles } from "@material-ui/core/styles";
import RootRef from "@material-ui/core/RootRef";
import CardMedia, { CardMediaProps } from "@material-ui/core/CardMedia";
import useLazyLoadedImage from "../../../../hooks/lazy-loaded-image";
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

export default function LazyLoadingImage({ media, title }: Props): JSX.Element {
  const classes = useStyles();
  const divRef = useRef(null);
  const isVisible = useLazyLoadedImage(divRef);
  const src = isVisible && media ? media.url : IMAGE_PLACEHOLDER;

  return (
    <RootRef rootRef={divRef}>
      <CardMedia className={classes.media} image={src} title={title} />
    </RootRef>
  );
}
