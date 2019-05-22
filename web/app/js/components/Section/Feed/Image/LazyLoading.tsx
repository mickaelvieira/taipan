import React, { useRef } from "react";
import RootRef from "@material-ui/core/RootRef";
import CardMedia, { CardMediaProps } from "@material-ui/core/CardMedia";
import useIsVisible from "../../../../hooks/is-visible";
import { IMAGE_PLACEHOLDER } from "../../../../constant/image";
import { Image } from "../../../../types/image";

interface Props extends CardMediaProps {
  media: Image | undefined;
}

export default function LazyLoadingImage({ media, title, className }: Props) {
  const divRef = useRef(null);
  const visible = useIsVisible(divRef);
  const src = visible && media ? media.url : IMAGE_PLACEHOLDER;

  return (
    <RootRef rootRef={divRef}>
      <CardMedia className={className} image={src} title={title} />
    </RootRef>
  );
}
