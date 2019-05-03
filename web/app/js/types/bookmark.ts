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
  id: string;
  url: string;
  lang: string;
  charset: string;
  title: string;
  description: string;
  image: string;
  addedAt: string;
  updatedAt: string;
  isRead: boolean;
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
