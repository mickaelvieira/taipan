export interface IndexLinks {
  edit?: string;
  first?: string;
  search?: string;
  user?: string;
}

export interface Index {
  href: string;
  links: IndexLinks;
}

export type Noop = () => void;
