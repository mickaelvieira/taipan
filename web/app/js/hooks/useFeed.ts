import { useEffect, useState, MutableRefObject, useCallback } from "react";

import FastDomBase from "fastdom";
import fastdomPromised from "fastdom/extensions/fastdom-promised";
import {
  Padding,
  calculateCursor,
  calculateHeightUpToFirstIndex,
  calculateHeightFromIndex,
  calculateInterval
} from "../helpers/feed";
import { FeedItem } from "../types/feed";

const FastDom = FastDomBase.extend(fastdomPromised);

interface Result {
  padding: Padding;
  items: FeedItem[];
}

class Scroll {
  // is the scroll active?
  private active = false;

  // previous scroll position
  private previous = 0;

  // current scroll position
  private position = 0;

  // record the scroll status
  record(active: boolean): void {
    this.active = active;
    if (active) {
      FastDom.measure(() => window.scrollY).then(position => {
        this.previous = this.position;
        this.position = position;
      });
    }
  }

  isActive(): boolean {
    return this.active;
  }

  getPosition(): number {
    return this.position;
  }

  isIdle(): boolean {
    return this.previous === this.position;
  }

  isUp(): boolean {
    return this.previous < this.position;
  }

  isDown(): boolean {
    return this.previous > this.position;
  }
}

class Feed {
  // first index in the feed
  private first = 0;

  // last index in the feed
  private last = 0;

  // Heigts of the HTML elements present in the feed
  private heights: number[] = [];

  async collectHeights(container: HTMLElement): Promise<void> {
    let j = this.first;

    const elements = Array.from(container.querySelectorAll(".feed-item"));
    const tasks = [];
    for (let i = 0, l = elements.length; i < l; i++) {
      if (!this.heights[j]) {
        tasks.push(
          FastDom.measure(
            (function(index) {
              return () => {
                const rect = elements[i].getBoundingClientRect();
                return {
                  index,
                  height: rect.height + 24
                };
              };
            })(j)
          )
        );
      }
      j++;
    }

    const heights = await Promise.all(tasks);

    heights.forEach(({ index, height }) => {
      this.heights[index] = height;
    });
  }

  adjust(scroll: Scroll, results: FeedItem[]): Result | null {
    const gap = Math.abs(scroll.getPosition()) + 70;
    const cursor = calculateCursor(gap, this.heights);
    const [first, last] = calculateInterval(cursor, results.length);

    if (first === this.first && last === this.last) {
      return null;
    }

    this.first = first;
    this.last = last;

    const items = [];
    for (let index = 0, l = results.length; index < l; index++) {
      if (index >= this.first && index <= this.last) {
        items.push(results[index]);
      }
    }

    const top = this.getTopPadding();
    const bottom = this.getBottomPadding();

    return {
      padding: {
        top,
        bottom
      },
      items
    };
  }

  private getTopPadding(): number {
    return calculateHeightUpToFirstIndex(this.first, this.heights);
  }

  private getBottomPadding(): number {
    return calculateHeightFromIndex(this.last, this.heights);
  }
}

const feed = new Feed();
const scroll = new Scroll();

export default function useFeed(
  ref: MutableRefObject<HTMLElement | undefined>,
  results: FeedItem[]
): Result {
  const [result, setResult] = useState<Result>({
    padding: { top: 0, bottom: 0 },
    items: results
  });

  const adjust = useCallback(() => {
    if (ref.current) {
      feed.collectHeights(ref.current).then(() => {
        const result = feed.adjust(scroll, results);
        if (result) {
          setResult(result);
        }
      });
    }
  }, [ref, results]);

  useEffect(() => {
    let stopTimeout: number | undefined = undefined;
    let scrollTimeout: number | undefined = undefined;

    function clearTimer(): void {
      window.clearTimeout(stopTimeout);
      stopTimeout = undefined;
    }

    function clearLongTimer(): void {
      window.clearTimeout(scrollTimeout);
      scrollTimeout = undefined;
    }

    function onLongTimer(): void {
      clearLongTimer();
      adjust();
    }

    function startLongTime(): void {
      if (!scrollTimeout) {
        scrollTimeout = window.setTimeout(onLongTimer, 400);
      }
    }

    function onScrollStop(): void {
      clearLongTimer();
      scroll.record(false);
      adjust();
    }

    function onScrollHandler(): void {
      clearTimer();
      startLongTime();
      scroll.record(true);
      stopTimeout = window.setTimeout(onScrollStop, 200);
    }

    window.addEventListener("scroll", onScrollHandler);

    adjust();

    return () => {
      clearTimer();
      window.removeEventListener("scroll", onScrollHandler);
    };
  }, [adjust, results]);

  return result;
}
