import { SearchType } from "../types/search";
import { getSearchTerms, getSearchType } from "../helpers/search";

export default function useSearch(): [SearchType, string[]] {
  const url = new URL(`${document.location}`);
  const params = url.searchParams;
  const terms = getSearchTerms(params.get("terms"));
  const type = getSearchType(url.pathname, params.get("type") as SearchType);
  return [type, terms];
}
