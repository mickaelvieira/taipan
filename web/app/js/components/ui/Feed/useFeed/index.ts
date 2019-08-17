import {
  useEffect,
  useState,
  MutableRefObject,
  useCallback,
  useMemo
} from "react";
import { FeedName, FeedItem } from "../../../../types/feed";
import Scroll from "./Scroll";
import getFeed, { Result } from "./Feed";

const scroll = new Scroll();

export default function useFeed(
  name: FeedName,
  ref: MutableRefObject<HTMLElement | null>,
  results: FeedItem[]
): Result {
  const feed = useMemo(() => getFeed(name), [name]);
  const [result, setResult] = useState<Result>({
    padding: { top: 0, bottom: 0 },
    items: results
  });

  const adjust = useCallback(async () => {
    if (ref.current) {
      await feed.collectHeights(ref.current);
      const result = feed.adjust(scroll.getPosition(), results);
      if (result) {
        setResult(result);
      }
    }
  }, [feed, ref, results]);

  useEffect(() => {
    let scrollStopTimeout: number | undefined = undefined;
    let scrollTimeout: number | undefined = undefined;

    function clearScrollStopTimer(): void {
      window.clearTimeout(scrollStopTimeout);
      scrollStopTimeout = undefined;
    }

    function clearScrollTimer(): void {
      window.clearTimeout(scrollTimeout);
      scrollTimeout = undefined;
    }

    function onScroll(): void {
      clearScrollTimer();
      adjust();
    }

    function onScrollStop(): void {
      clearScrollTimer();
      scroll.record(false);
      adjust();
    }

    function startScrollTimer(): void {
      if (!scrollTimeout) {
        scrollTimeout = window.setTimeout(onScroll, 400);
      }
    }

    function startScrollStopTimer(): void {
      if (!scrollStopTimeout) {
        scrollStopTimeout = window.setTimeout(onScrollStop, 200);
      }
    }

    function onScrollHandler(): void {
      clearScrollStopTimer();
      startScrollTimer();
      startScrollStopTimer();
      scroll.record(true);
    }

    window.addEventListener("scroll", onScrollHandler);

    adjust();

    return () => {
      clearScrollTimer();
      clearScrollStopTimer();
      window.removeEventListener("scroll", onScrollHandler);
    };
  }, [adjust]);

  return result;
}
