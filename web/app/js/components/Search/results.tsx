import withDispatch from "components/HoC/withDispatch";
import withWindow from "components/HoC/withWindow";
import Observer from "./Results/Observer";
import { attachFragment } from "lib/view";
import Results from "./Results/index";

let observer;
let container;

const Container = ({ window, dispatch, ...state }) => {
  if (!observer) {
    observer = new Observer(window, dispatch);
  }

  observer.onStateChange(state);

  if (!container) {
    container = window.document.querySelector(
      ".search-bookmark-results-container"
    );
  }

  attachFragment(container, Results(state));
};

Container.storeMapper = state => ({ ...state.search });

export default withDispatch(withWindow(Container));
