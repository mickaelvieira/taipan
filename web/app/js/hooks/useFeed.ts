import { useEffect, useState } from "react";
import {
  Padding,
  calculateCursor,
  calculateHeightUpToFirstIndex,
  calculateHeightFromIndex,
  calculateInterval
} from "../helpers/feed";
import { FeedItem } from "../types/feed";

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
      const position = document.documentElement.getBoundingClientRect().top;
      this.previous = this.position;
      this.position = position;
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
  // // first index in the feed
  private first = 0;

  // // last index in the feed
  private last = 0;

  // Heigts of the HTML elements present in the feed
  private heights: number[] = [];

  collectHeights(items: Element[]): void {
    let j = this.first;
    // console.log("collect");
    // console.log(j);
    // console.log(items);
    for (let i = 0, l = items.length; i < l; i++) {
      if (!this.heights[j]) {
        const rect = items[i].getBoundingClientRect();
        this.heights[j] = rect.height + 24;
      }
      j++;
    }
  }

  adjust(scroll: Scroll, results: FeedItem[]): Result | null {
    console.log(this.heights);
    const gap = Math.abs(scroll.getPosition()) + 70;
    const cursor = calculateCursor(gap, this.heights);
    const [first, last] = calculateInterval(
      cursor,
      results.length
      // scroll.isDown()
    );

    if (first === this.first && last === this.last) {
      return null;
    }

    this.first = first;
    this.last = last;

    // console.log(this.previous)
    // console.log(this.interval)
    // console.log(`UP: ${scroll.isUp()}`)
    // console.log(`DOWN: ${scroll.isDown()}`)
    // console.log(`IDLE ${scroll.isIdle()}`)

    // console.log(`TOTAL ${results.length}`)
    // console.log(`FIRST ${this.first}`)
    // console.log(`LAST ${this.last}`)

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
  elements: Element[],
  results: FeedItem[]
): Result {
  console.log("=================== useFeedScrolling ===================");
  const [result, setResult] = useState<Result>({
    padding: { top: 0, bottom: 0 },
    items: results
  });

  useEffect(() => {
    let timeout: number | undefined = undefined;

    function clearTimer(): void {
      if (timeout) {
        window.clearTimeout(timeout);
      }
    }

    function onScrollStop(): void {
      scroll.record(false);
      feed.collectHeights(elements);
      const result = feed.adjust(scroll, results);
      if (result) {
        setResult(result);
      }
    }

    function onScrollHandler(): void {
      clearTimer();
      scroll.record(true);
      timeout = window.setTimeout(onScrollStop, 400);
    }

    window.addEventListener("scroll", onScrollHandler);

    return () => {
      clearTimer();
      window.removeEventListener("scroll", onScrollHandler);
    };
  }, [elements, results]);

  return result;
}
