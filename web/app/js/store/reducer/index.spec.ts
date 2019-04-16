import reducer from ".";
import defaultState from "./default";
import {
  updateConnectionStatus,
  selectSearchResult,
  hideSearchResults,
  showSearchResults,
  updateSearchTerms,
  updateSearchResults,
  prepareSearch
} from "../actions";

describe("reducer", () => {
  let state;

  beforeEach(() => {
    state = reducer(defaultState, { type: "@@redux/INIT" });
  });

  test("is initialized with a default state", () => {
    expect(reducer(state, { type: "DUMMY_ACTION" })).toBe(defaultState);
  });

  test("assigns the connection status to the store", () => {
    state = reducer(defaultState, updateConnectionStatus(true));
    expect(state.online).toBe(true);
    state = reducer(defaultState, updateConnectionStatus(false));
    expect(state.online).toBe(false);
  });

  test("assigns the state of the search result panel to the store", () => {
    expect(state.search.visible).toBe(false);
    state = reducer(state, showSearchResults());
    expect(state.search.visible).toBe(true);
    state = reducer(state, hideSearchResults());
    expect(state.search.visible).toBe(false);
  });

  test("assigns the search terms to the store", () => {
    const terms = ["term1", "term2"];
    expect(state.search.terms.length).toBe(0);
    state = reducer(state, updateSearchTerms(terms));
    expect(state.search.terms.length).toBe(2);
    expect(state.search.terms).toBe(terms);
  });

  test("prepares the search before fetching results", () => {
    state = reducer(state, prepareSearch());
    expect(state.search.visible).toBe(true);
    expect(state.search.fetching).toBe(true);
    expect(state.search.results).toEqual([]);
    expect(state.search.selected).toBe(null);
  });

  xtest("updates the store with the newly received results", () => {
    const results = [
      { title: "article 1", description: "" },
      { title: "article 2", description: "" }
    ];
    // state.search.terms = [];
    state = reducer(state, updateSearchResults(results));
    expect(state.search.visible).toBe(true);
    expect(state.search.fetching).toBe(false);
    expect(state.search.results).toEqual(results);
    expect(state.search.selected).toBe(null);
  });

  test("assigns the selected result to the store", () => {
    const results = [
      { title: "article 1", description: "" },
      { title: "article 2", description: "" }
    ];
    state = reducer(state, selectSearchResult(results[1]));
    expect(state.search.selected).toBe(results[1]);
  });
});
