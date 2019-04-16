import { all, call, put, takeEvery, takeLatest } from "redux-saga/effects";
import {
  fetchBookmark,
  fetchIndex,
  fetchItems,
  historyPushState,
  markWallAsFetching,
  showBookmark
} from "../actions";
// import Api from '...'

// store.dispatch(markWallAsFetching());
// store.dispatch(fetchIndex()).then(function() {
//   store.dispatch(fetchUser());
//   store.dispatch(markWallAsFetching());
//   store.dispatch(fetchItems());
// });

// worker Saga: will be fired on USER_FETCH_REQUESTED actions
// function* fetchUser(actions) {
//   try {
//     const user = yield call(Api.fetchUser, actions.payload.userId);
//     yield put({ type: "USER_FETCH_SUCCEEDED", user: user });
//   } catch (e) {
//     yield put({ type: "USER_FETCH_FAILED", message: e.message });
//   }
// }

/*
  Starts fetchUser on each dispatched `USER_FETCH_REQUESTED` actions.
  Allows concurrent fetches of user.
*/
// function* mySaga() {
//   yield takeEvery("USER_FETCH_REQUESTED", fetchUser);
// }

/*
  Alternatively you may use takeLatest.

  Does not allow concurrent fetches of user. If "USER_FETCH_REQUESTED" gets
  dispatched while a fetch is already pending, that pending fetch is cancelled
  and only the latest one will be run.
*/
// function* searchSaga() {
//   yield markWallAsFetching();
//   yield fetchIndex();
// }

function* showBookmarkPreview(actions) {
  const { payload } = actions;
  const { item } = payload;

  console.log("showBookmarkPreview");
  console.log(item);
  yield historyPushState(null, item.title, "/bookmark/" + item.id);
  yield showBookmark();
  const [response, info, history] = yield fetchBookmark(item.id);

  console.log(response);
  console.log(info);
  console.log(history);

  yield put({
    type: "BOOKMARK_PREVIEW_SHOWN",
    payload: { response, info, history }
  });
}

function* watchBookmarkPreview() {
  console.log("watchBookmarkPreview");
  yield takeEvery("SHOW_BOOKMARK_PREVIEW", showBookmarkPreview);
}

// notice how we now only export the rootSaga
// single entry point to start all Sagas at once
export default function* rootSaga() {
  yield all([watchBookmarkPreview()]);
}
