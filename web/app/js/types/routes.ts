import { RouteComponentProps } from "react-router";

export interface BooknarkParams {
  id?: string;
  section?: string;
}

export type RouteBookmarkProps = RouteComponentProps<BooknarkParams>;
