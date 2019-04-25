export interface BookmarkLinks {
  about: string;
  read: string;
  self: string;
  unread: string;
  image: string;
}

export interface BookmarksLinks {
  edit?: string;
  first?: string;
  last?: string;
  prev?: string;
  next?: string;
  search?: string;
}

export interface UserBookmark {
  href: string;
  id: string;
  url: string;
  title: string;
  hash: string;
  charset: string;
  description: string;
  image: string;
  canonical_url: string;
  added_at: string;
  accessed_at: string;
  created_at: string;
  updated_at: string;
  // is_read: boolean;
  // is_pending: boolean;
  // is_fetching: boolean;
  // is_fetched: boolean;
  // links: BookmarkLinks;
  // history?: BookmarkHistory;
}

export interface BookmarkHistoryEntry {
  href: string;
  created_at: string;
  request_uri: string;
  request_method: string;
  response_code: number;
  response_phrase: string;
}

export interface BookmarkHistory {
  items: BookmarkHistoryEntry[];
}

export interface Bookmarks {
  items: Bookmark[];
  links: BookmarksLinks;
}
