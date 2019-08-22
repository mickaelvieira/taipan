import React from "react";
import { useTheme } from "@material-ui/core/styles";
import useMediaQuery from "@material-ui/core/useMediaQuery";
import { ExternalLink } from "../../ui/Link";
import Typography from "@material-ui/core/Typography";
import { Subscription } from "../../../types/subscription";
import { truncate } from "../../../helpers/string";
import { getDomain } from "../../../helpers/syndication";

interface Props {
  item: Subscription;
  shouldTruncate?: boolean;
  className?: string;
}

export default React.memo(function SubscriptionTitle({
  item,
  shouldTruncate = false,
  className
}: Props): JSX.Element {
  const theme = useTheme();
  const lg = useMediaQuery(theme.breakpoints.up("lg"));
  const md = useMediaQuery(theme.breakpoints.up("md"));
  const sm = useMediaQuery(theme.breakpoints.up("sm"));
  const url = getDomain(item);

  let chars = 30;
  if (lg) {
    chars = 20;
  } else if (md) {
    chars = 90;
  } else if (sm) {
    chars = 60;
  }

  return (
    <ExternalLink
      href={`${url.protocol}//${url.host}`}
      title={item.title ? item.title : `${item.url}`}
      className={className ? className : ""}
    >
      <Typography component="span">
        {item.title
          ? shouldTruncate
            ? truncate(item.title, chars)
            : item.title
          : "no title"}
      </Typography>
    </ExternalLink>
  );
});
