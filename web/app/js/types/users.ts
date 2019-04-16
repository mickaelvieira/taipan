export interface UserLinks {
  self: string;
}

export interface User {
  id: string;
  firstname: string;
  lastname: string;
  username: string;
  email: string;
  href: string;
  created_at: string;
  updated_at: string;
  links: UserLinks;
}
