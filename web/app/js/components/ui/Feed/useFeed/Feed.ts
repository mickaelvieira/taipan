import FastDomBase from "fastdom";
import fastdomPromised from "fastdom/extensions/fastdom-promised";
import {
  calculateCursor,
  calculateHeightUpToFirstIndex,
  calculateHeightFromIndex,
  calculateInterval
} from "./helpers";
import { FeedName, FeedItem } from "../../../../types/feed";

const FastDom = FastDomBase.extend(fastdomPromised);

interface Padding {
  top: number;
  bottom: number;
}

export interface Result {
  padding: Padding;
  items: FeedItem[];
}

const ElementsMarginBottom = 24;

class Feed {
  // feed's name
  name: string;

  // first index in the feed
  private first = 0;

  // last index in the feed
  private last = 0;

  // Heights of the HTML elements present in the feed
  private heights: number[] = [];

  constructor(name: string) {
    this.name = name;
  }

  async collectHeights(container: HTMLElement): Promise<void> {
    let j = this.first;

    const tasks = Array.from(container.querySelectorAll(".feed-item")).map(
      element =>
        FastDom.measure(
          (function(index) {
            return () => {
              const rect = element.getBoundingClientRect();
              const height = rect.height + ElementsMarginBottom;
              return {
                index,
                height
              };
            };
          })(j++)
        )
    );

    const heights = await Promise.all(tasks);
    heights.forEach(({ index, height }) => {
      this.heights[index] = height;
    });
  }

  adjust(scrollPosition: number, results: FeedItem[]): Result | null {
    const cursor = calculateCursor(scrollPosition, this.heights);
    const [first, last] = calculateInterval(cursor, results.length);

    if (first === this.first && last === this.last) {
      return null;
    }

    this.first = first;
    this.last = last;

    const items = results.filter(
      (_, index) => index >= this.first && index <= this.last
    );
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

const cache = new Map<FeedName, Feed>();

export default function(name: FeedName): Feed {
  if (!cache.has(name)) {
    cache.set(name, new Feed(name));
  }
  return cache.get(name) as Feed;
}
