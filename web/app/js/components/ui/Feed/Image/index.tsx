import React from "react";
import LazyLoadingImage from "./LazyLoading";
import ExternalLink from "../../../ui/Link/External";
import { Bookmark } from "../../../../types/bookmark";
import { Document } from "../../../../types/document";

interface Props {
  item: Bookmark | Document;
}

export default function ItemImage({ item }: Props): JSX.Element | null {
  return !item.image ? null : (
    <ExternalLink underline="none" href={item.url} title={item.title}>
      <LazyLoadingImage media={item.image} title={item.title} />
    </ExternalLink>
  );
}
