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

export type Undoer = () => void;
export type CacheUpdater = () => void;

export interface MessageInfo {
  message: string;
  label?: string;
  action?: Undoer;
}
