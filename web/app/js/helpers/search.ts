import { SearchType } from "../types/search";
export function getSearchTerms(search: string | null): string[] {
  if (!search) {
    return [];
  }
  return decodeURIComponent(search)
    .split(/(\s)+/)
    .map(term => term.trim())
    .filter(term => term !== "");
}

export function getSearchType(
  pathname: string,
  type: SearchType | null
): SearchType {
  if (!type) {
    return pathname === "/" ? "document" : "bookmark";
  }
  return type;
}

export function getSearchParams(terms: string[]): string {
  return terms.map(term => `terms=${term}`).join(" ");
}
