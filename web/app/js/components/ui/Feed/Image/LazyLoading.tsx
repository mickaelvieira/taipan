import React, { useRef } from "react";
import { makeStyles } from "@material-ui/core/styles";
import RootRef from "@material-ui/core/RootRef";
import CardMedia, { CardMediaProps } from "@material-ui/core/CardMedia";
import useLazyLoadedImage from "../../../../hooks/lazy-loaded-image";
import { IMAGE_PLACEHOLDER } from "../../../../constant/image";
import { Image } from "../../../../types/image";
import { getImageWidth } from "../../../../helpers/image";

interface Props extends CardMediaProps {
  media: Image | null;
}

const useStyles = makeStyles(({ breakpoints }) => ({
  media: {
    backgroundSize: "cover",
    minHeight: `calc(${getImageWidth("sm")}px * 9 / 16)`,
    [breakpoints.up("md")]: {
      minHeight: `calc(${getImageWidth("sm")}px * 9 / 16)`,
    },
  },
}));

export default function LazyLoadingImage({ media, title }: Props): JSX.Element {
  const classes = useStyles();
  const divRef = useRef(null);
  const isVisible = useLazyLoadedImage(divRef);
  const src = isVisible && media ? `${media.url}` : IMAGE_PLACEHOLDER;

  return (
    <RootRef rootRef={divRef}>
      <CardMedia className={classes.media} image={src} title={title} />
    </RootRef>
  );
}
