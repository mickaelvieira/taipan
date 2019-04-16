import { createStore, combineReducers, applyMiddleware, Store } from "redux";
import { createLogger } from "redux-logger";
import thunk from "redux-thunk";
import indexReducer from "store/reducer/index";
import feedReducer from "store/reducer/feed";
import userReducer from "store/reducer/user";
import { RootState } from "store/reducer/default";

export default function(state: RootState): Store<RootState> {
  const middlewares = [thunk, createLogger({ collapsed: true })];
  const reducers = combineReducers<RootState>({
    index: indexReducer,
    feed: feedReducer,
    user: userReducer
  });

  return createStore(reducers, state, applyMiddleware(...middlewares));
}
