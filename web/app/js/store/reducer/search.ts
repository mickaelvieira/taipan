// case HISTORY_PUSH_STATE: {
//   const { data = null, title, url } = payload;

//   history.pushState(data, title, url);

//   return state;
// }
// case HISTORY_REPLACE_STATE: {
//   const { data = null, title, url } = payload;

//   history.replaceState(data, title, url);

//   return state;
// }
// case CONNECTION_STATUS_CHANGE: {
//   return {
//     ...state,
//     online: payload.online
//   };
// }
// case UPDATE_USER: {
//   const collection = parseCollection(payload.user.collection);
//   const user = collection.items[0];

//   return { ...state, user };
// }
// case PREPARE_SEARCH:
// case HIDE_SEARCH_RESULTS:
// case SHOW_SEARCH_RESULTS:
// case SEARCH_TERMS_UPDATE:
// case SELECT_SEARCH_RESULT: {
//   const search = { ...state.search, ...payload };
//   return { ...state, search };
// }
// case SEARCH_RESULTS_UPDATE: {
//   const items = parseItems(payload.results);
//   const re = new RegExp(`(${state.search.terms.join("|")})`, "gi");
//   let search = {
//     ...state.search,
//     ...payload,
//     results: items
//   };

//   return { ...state, search };
// }
// case WALL_ITEMS_PAGE: {
//   let page = state.feed.page;
//   page += 1;
//
//   return { ...state, {
//     feed: { ...state.feed, { page })
//   });
// }
// case UPDATE_FILTERS: {
//   const { filter } = payload;

//   const filters = { ...state.feed.filters, [filter.name]: filter.value };

//   const feed = {
//     ...state.feed,
//     filters
//   };

//   return { ...state, feed };
// }
// case WALL_ITEMS_FETCHING: {
//   const { fetching } = payload;
//   const feed = {
//     ...state.feed,
//     fetching
//   };

//   return { ...state, feed };
// }

// case UPDATE_BOOKMARK: {
//   const { html, info, history: collection } = payload;
//   const item = parseItem(info.collection.items[0]);
//   const history = parseCollection(collection.collection);

//   const bookmark = {
//     ...state.bookmark,
//     ...item,
//     history,
//     html
//   };

//   return { ...state, bookmark };
// }
// case SHOW_BOOKMARK: {
//   const bookmark = {
//     ...state.bookmark,
//     visible: payload.visible
//   };
//   return { ...state, bookmark };
// }
// case HIDE_BOOKMARK: {
//   const bookmark = {
//     ...state.bookmark,
//     visible: payload.visible,
//     html: "",
//     data: null,
//     history: []
//   };
//   return { ...state, bookmark };
// }
