import React, { PropsWithChildren } from "react";
import { makeStyles } from "@material-ui/core/styles";
import Link from "@material-ui/core/Link";
import { RouterLink } from "../../ui/Link";

const useStyles = makeStyles(({ breakpoints, palette, typography }) => ({
  link: {
    display: "block",
    fontSize: typography.fontSize,
    color: palette.grey[600],
    [breakpoints.up("md")]: {
      color: palette.grey[100]
    },
    "&.active": {
      color: palette.primary.main
    }
  }
}));

interface Props {
  to: string;
  onClick: (event: React.MouseEvent) => void;
}

export default function MenuLink({
  to,
  children,
  onClick
}: PropsWithChildren<Props>): JSX.Element {
  const classes = useStyles();

  return (
    <Link
      to={to}
      classes={{
        root: classes.link
      }}
      component={RouterLink}
      underline="none"
      onClick={onClick}
    >
      {children}
    </Link>
  );
}
