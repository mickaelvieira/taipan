import { SFC, useState } from "react";
import { connect } from "react-redux";
import { Bookmark } from "types/bookmark";
import { withRouter } from "react-router";
import { RouteBookmarkProps } from "types/routes";
import { RootState } from "store/reducer/default";
import { selectFeedBookmarks, getCurrentIndex } from "store/selector";

export interface Pagination {
  first?: Bookmark;
  prev?: Bookmark;
  current?: Bookmark;
  next?: Bookmark;
  last?: Bookmark;
  indices: Indices;
  toFirst: () => void;
  toPrev: () => void;
  toNext: () => void;
  toLast: () => void;
}

export interface PaginationProps {
  pagination: Pagination;
  isNearTheEnd: boolean;
}

interface StateProps {
  items: Bookmark[];
  current: number;
}

interface Props {
  children: (args: PaginationProps) => JSX.Element;
}

type ComponentProps = Props & StateProps & RouteBookmarkProps;

interface Indices {
  [key: string]: number;
}

interface Values {
  [key: string]: Bookmark;
}

const getIndices = (index: number, total: number): Indices => ({
  first: 0,
  prev: index - 1,
  current: index,
  next: index + 1,
  last: total - 1
});

const getValues = (indices: Indices, items: any[]): Values => ({
  first: items[indices.first],
  prev: items[indices.prev],
  current: items[indices.current],
  next: items[indices.next],
  last: items[indices.last]
});

interface Traps {
  [key: string]: () => void;
}

type Target = Values & { indices: Indices };
type SetIndex = (idx: number) => void;

function getPagination(data: Target, setIndex: SetIndex) {
  const traps: Traps = Object.freeze({
    toFirst: function() {
      setIndex(this.indices.first);
    },
    toPrev: function() {
      const { prev, first } = this.indices;
      setIndex(prev < first ? first : prev);
    },
    toNext: function() {
      const { next, last } = this.indices;
      setIndex(next > last ? last : next);
    },
    toLast: function() {
      setIndex(this.indices.last);
    }
  });

  const isTrapped = (name: string) => Reflect.ownKeys(traps).includes(name);

  return new Proxy(data, {
    get: function(target, prop: string, receiver) {
      if (isTrapped(prop)) {
        return function() {
          traps[prop].call(target);
          return receiver;
        };
      }

      return typeof target[prop] === "function"
        ? Reflect.get(target, prop)
        : target[prop];
    }
  });
}

const BookmarkPagination: SFC<ComponentProps> = ({ children, items }) => {
  const [index, setIndex] = useState(0);
  const indices = getIndices(index, items.length);
  const values = getValues(indices, items);
  const pagination = getPagination({ ...values, indices }, setIndex);

  console.log(`${index + 1}/${items.length}`);

  console.log(pagination);
  const isNearTheEnd = index > 0 && items.length - index <= 10;

  return children({
    isNearTheEnd,
    pagination
  });
};

const mapStateToProps = (state: RootState, ownProps: ComponentProps) => {
  const {
    params: { id }
  } = ownProps.match;

  const items = selectFeedBookmarks(state);
  const current = getCurrentIndex(id, items);

  return {
    items,
    current
  };
};

export default withRouter(connect(mapStateToProps)(BookmarkPagination));
