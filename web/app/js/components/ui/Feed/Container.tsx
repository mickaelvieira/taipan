import React, { useRef, useState, useEffect } from "react";
import { makeStyles } from "@material-ui/core/styles";
import useFeedScrolling from "../../../hooks/useFeedScrolling";
import { FeedItem } from "../../../types/feed";
import {
  calculateFirstIndex,
  calculateTopGap,
  calculateBottomGap,
  calculateBoudaries
} from "../../../helpers/feed";

const useStyles = makeStyles({
  container: {
    display: "flex",
    flexDirection: "column",
    marginBottom: 60,
    minHeight: "100vh"
  }
});

interface RenderProps {
  results: FeedItem[];
}

interface Props {
  results: FeedItem[];
  children: (props: RenderProps) => JSX.Element;
}

export default function FeedContainer({
  children,
  results
}: // paddingTop,
// paddingBottom,
Props): JSX.Element {
  const classes = useStyles();
  const height = useFeedScrolling();
  const gap = Math.abs(height) + 70;
  const indices = useRef<number[]>([]);
  const [tracking, setTracking] = useState({
    start: 0,
    first: 0,
    last: 9,
    results,
    paddingTop: 0,
    paddingBottom: 0
  });
  // const list = useRef<HTMLElement | null>(null)

  useEffect(() => {
    const list = document.getElementById("feed");
    if (list !== null) {
      console.time("calc");

      const heights = indices.current;
      const items = Array.from(list.querySelectorAll(".feed-item"));
      let j = tracking.first;
      for (var i = 0, l = items.length; i < l; i++) {
        const item = items[i];
        const rect = item.getBoundingClientRect();
        heights[j] = rect.height + 24;
        j++;
      }

      // console.log("== heights ==");
      // console.log(j);
      // console.log(heights);
      // console.log(`gap to fill: ${gap}`);

      const start = calculateFirstIndex(gap, heights);
      const [first, last] = calculateBoudaries(start, results.length);
      console.log(`from ${first} to ${last}`);
      const paddingTop = calculateTopGap(first, heights);
      const paddingBottom = calculateBottomGap(last, heights);

      console.log(`Interval [${first}, ${last}]`);

      const r = [];
      for (let i = 0, l = results.length; i < l; i++) {
        if (i >= first && i <= last) {
          // console.log(`Including ${i}`)
          r.push(results[i]);
        } else {
          // console.log(`Excluding ${i}`)
        }
      }

      setTracking({
        start,
        first,
        last,
        results: r,
        paddingTop,
        paddingBottom
      });

      console.timeEnd("calc");
    }
    // console.log(list.current.querySelectorAll(".feed-item"))
  }, [gap, results, tracking.first]);

  console.log(tracking);

  return (
    <section
      id="feed"
      style={{
        paddingTop: `${tracking.paddingTop}px`,
        paddingBottom: `${tracking.paddingBottom}px`
      }}
      className={classes.container}
    >
      {children({ results: tracking.results })}
    </section>
  );
}
