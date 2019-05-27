import React from "react";
import EagerLoadingImage from "./EagerLoading";
import LazyLoadingImage from "./LazyLoading";
import Link from "@material-ui/core/Link";
import { Bookmark } from "../../../../types/bookmark";
import { Document } from "../../../../types/document";

interface Props {
  index: number;
  item: Bookmark | Document;
}

export default function ItemImage({ index, item }: Props) {
  const ImageComp = index < 5 ? EagerLoadingImage : LazyLoadingImage;

  return !item.image ? null : (
    <Link
      underline="none"
      href={item.url}
      title={item.title}
      target="_blank"
      rel="noopener"
    >
      <ImageComp media={item.image} title={item.title} />
    </Link>
  );
}
