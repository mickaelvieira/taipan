import { attachFragment } from "lib/view/index";
import Observer from "./Observer";
import withWindow from "components/HoC/withWindow";
import withDispatch from "components/HoC/withDispatch";
import Preview from "./Preview";

let observer;
let container;

const Container = ({ window, dispatch, ...state }) => {
  if (!observer) {
    observer = new Observer(window, dispatch);
  }

  observer.onStateChange(state);

  if (!container) {
    container = window.document.querySelector(".bookmark");
  }

  const { visible } = state;

  if (visible) {
    attachFragment(container, Preview(state));
  }
};

Container.storeMapper = state => ({ ...state.bookmark });

export default withDispatch(withWindow(Container));
