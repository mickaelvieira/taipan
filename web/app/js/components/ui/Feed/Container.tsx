import React, { useRef, PropsWithChildren } from "react";
import { makeStyles } from "@material-ui/core/styles";
import useFeed from "./useFeed";
import { ListProps } from "./Feed";
import PointerEvents from "./PointerEvents";
import { FeedItem, FeedName } from "../../../types/feed";

const useStyles = makeStyles({
  container: {
    display: "flex",
    flexDirection: "column",
    marginBottom: 60,
    minHeight: "100vh",
    pointerEvents: "auto",
  },
});

interface Props extends ListProps {
  name: FeedName;
  List: React.FunctionComponent<ListProps>;
  results: FeedItem[];
}

export default React.memo(function FeedContainer({
  name,
  List,
  results,
  ...rest
}: PropsWithChildren<Props>): JSX.Element {
  const classes = useStyles();
  const ref = useRef<HTMLElement>(null);
  const { padding, items } = useFeed(name, ref, results);

  return (
    <PointerEvents>
      <section
        ref={ref}
        style={{
          paddingTop: `${padding.top}px`,
          paddingBottom: `${padding.bottom}px`,
        }}
        className={classes.container}
      >
        <List results={items} {...rest} />
      </section>
    </PointerEvents>
  );
});
