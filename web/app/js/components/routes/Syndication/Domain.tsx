import React from "react";
import { ExternalLink } from "../../ui/Link";
import { Source } from "../../../types/syndication";
import { getDomain } from "../../../helpers/syndication";

interface Props {
  item: Source;
}

export default React.memo(function SourceDomain({ item }: Props): JSX.Element {
  const url = getDomain(item);

  return (
    <ExternalLink
      href={`${url.protocol}//${url.host}`}
      title={item.title ? item.title : `${item.url}`}
    >
      {`${url.protocol}//${url.host}`}
    </ExternalLink>
  );
});
