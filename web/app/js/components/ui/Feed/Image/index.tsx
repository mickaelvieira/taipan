import React from "react";
import LazyLoadingImage from "./LazyLoading";
import Link from "@material-ui/core/Link";
import { Bookmark } from "../../../../types/bookmark";
import { Document } from "../../../../types/document";

interface Props {
  item: Bookmark | Document;
}

export default function ItemImage({ item }: Props): JSX.Element | null {
  return !item.image ? null : (
    <Link
      underline="none"
      href={item.url}
      title={item.title}
      target="_blank"
      rel="noopener"
    >
      <LazyLoadingImage media={item.image} title={item.title} />
    </Link>
  );
}
