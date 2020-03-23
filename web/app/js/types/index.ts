export type Noop = () => void;

export interface CursorPagination {
  from?: string;
  to?: string;
  limit?: number;
}

export interface OffsetPagination {
  offset?: number;
  limit?: number;
}

export interface Sorting {
  by: string;
  dir: SortingDirection;
}

export enum SortingDirection {
  ASC = "ASC",
  DESC = "DESC",
}

export type Undoer = () => void;
export type CacheUpdater = () => void;

export interface MessageInfo {
  message: string;
  label?: string;
  action?: Undoer;
}
