import React, { useRef, PropsWithChildren } from "react";
import { makeStyles } from "@material-ui/core/styles";
import useFeed from "../../../hooks/useFeed";
import { ListProps } from "./Feed";

const useStyles = makeStyles({
  container: {
    display: "flex",
    flexDirection: "column",
    marginBottom: 60,
    minHeight: "100vh",
    pointerEvents: "auto"
  },
  scrolling: {
    pointerEvents: "none"
  }
});

interface Props extends ListProps {
  List: React.FunctionComponent<ListProps>;
}

export default React.memo(function FeedContainer({
  List,
  results,
  ...rest
}: PropsWithChildren<Props>): JSX.Element {
  const classes = useStyles();
  const div = useRef<HTMLElement>(null);
  const items = div.current
    ? Array.from(div.current.querySelectorAll(".feed-item"))
    : [];
  const result = useFeed(items, results);

  return (
    <section
      id="feed"
      ref={div}
      style={{
        paddingTop: `${result.padding.top}px`,
        paddingBottom: `${result.padding.bottom}px`
      }}
      className={`${classes.container} ${false ? classes.scrolling : ""}`}
    >
      <List results={result.items} {...rest} />
    </section>
  );
});
