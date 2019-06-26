export type Noop = () => void;

export interface CursorPagination {
  from?: string;
  to?: string;
  limit?: number;
}

export interface OffsetPagination {
  offset?: string;
  limit?: number;
}
