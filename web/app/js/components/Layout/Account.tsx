import React, { PropsWithChildren } from "react";
import { makeStyles } from "@material-ui/core/styles";
import MainLayout from "./Layout";
import MainContent from "./Content";

const useStyles = makeStyles(() => ({
  content: {
    paddingLeft: 12,
    paddingRight: 12
  }
}));

export default function AccountLayout({
  children
}: PropsWithChildren<{}>): JSX.Element {
  const classes = useStyles();
  return (
    <MainLayout>
      {() => <MainContent className={classes.content}>{children}</MainContent>}
    </MainLayout>
  );
}
