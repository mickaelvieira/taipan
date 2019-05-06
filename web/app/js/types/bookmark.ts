export interface BookmarkImage {
  name: string;
  url: string;
  width: number;
  height: number;
  format: string;
}

export interface UserBookmark {
  id: string;
  url: string;
  lang: string;
  charset: string;
  title: string;
  description: string;
  image?: BookmarkImage;
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

