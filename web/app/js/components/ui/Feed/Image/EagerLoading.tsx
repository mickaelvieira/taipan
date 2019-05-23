import React from "react";
import CardMedia, { CardMediaProps } from "@material-ui/core/CardMedia";
import { IMAGE_PLACEHOLDER } from "../../../../constant/image";
import { Image } from "../../../../types/image";

interface Props extends CardMediaProps {
  media: Image | null;
}

export default function EagerLoadingImage({ media, title, className }: Props) {
  const src = media ? media.url : IMAGE_PLACEHOLDER;
  return <CardMedia className={className} image={src} title={title} />;
}
