import React from "react";
import { useTheme } from "@material-ui/core/styles";
import useMediaQuery from "@material-ui/core/useMediaQuery";
import Link from "@material-ui/core/Link";
import Typography from "@material-ui/core/Typography";
import { Subscription } from "../../../types/subscription";
import { truncate } from "../../../helpers/string";
import { getDomain } from "../../../helpers/syndication";

interface Props {
  subscription: Subscription;
}

export default React.memo(function SubscriptionTitle({
  subscription
}: Props): JSX.Element {
  const theme = useTheme();
  const lg = useMediaQuery(theme.breakpoints.up("lg"));
  const md = useMediaQuery(theme.breakpoints.up("md"));
  const sm = useMediaQuery(theme.breakpoints.up("sm"));
  const url = getDomain(subscription);

  let chars = 30;
  if (lg) {
    chars = 20;
  } else if (md) {
    chars = 90;
  } else if (sm) {
    chars = 60;
  }

  return (
    <Link
      underline="none"
      href={`${url.protocol}//${url.host}`}
      title={subscription.title ? subscription.title : subscription.url}
      target="_blank"
      rel="noopener"
    >
      <Typography component="span">
        {truncate(subscription.title, chars)}
      </Typography>
    </Link>
  );
});
