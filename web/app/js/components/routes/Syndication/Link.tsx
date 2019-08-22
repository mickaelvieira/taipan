import React from "react";
import { ExternalLink } from "../../ui/Link";
import { Source } from "../../../types/syndication";

interface Props {
  item: Source;
}

export default React.memo(function SourceLink({ item }: Props): JSX.Element {
  return (
    <ExternalLink
      href={`${item.url}`}
      title={item.title ? item.title : `${item.url}`}
    >
      {`${item.url}`}
    </ExternalLink>
  );
});
