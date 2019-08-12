import React from "react";
import EagerLoadingImage from "./EagerLoading";
import { ExternalLink } from "../../../ui/Link";
import { Bookmark } from "../../../../types/bookmark";
import { Document } from "../../../../types/document";

interface Props {
  item: Bookmark | Document;
}

export default function ItemImage({ item }: Props): JSX.Element | null {
  return !item.image ? null : (
    <ExternalLink underline="none" href={`${item.url}`} title={item.title}>
      <EagerLoadingImage media={item.image} title={item.title} />
    </ExternalLink>
  );
}
