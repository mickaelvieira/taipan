import React, { useRef, PropsWithChildren } from "react";
import { makeStyles } from "@material-ui/core/styles";
import useFeed from "../../../hooks/useFeed";
import { ListProps } from "./Feed";
import PointerEvents from "./PointerEvents";
import { FeedItem } from "../../../types/feed";

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
  results: FeedItem[];
}

export default React.memo(function FeedContainer({
  List,
  results,
  ...rest
}: PropsWithChildren<Props>): JSX.Element {
  const classes = useStyles();
  const ref = useRef<HTMLElement>();
  const { padding, items } = useFeed(ref, results);

  return (
    <PointerEvents>
      <section
        id="feed"
        ref={ref}
        style={{
          paddingTop: `${padding.top}px`,
          paddingBottom: `${padding.bottom}px`
        }}
        className={`${classes.container}`}
      >
        <List results={items} {...rest} />
      </section>
    </PointerEvents>
  );
});
