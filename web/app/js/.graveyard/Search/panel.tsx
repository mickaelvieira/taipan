import Observer from "./Panel/Observer";
import withWindow from "components/HoC/withWindow";
import withDispatch from "components/HoC/withDispatch";

let observer;

const Panel = ({ window, dispatch, ...state }) => {
  if (!observer) {
    observer = new Observer(window, dispatch);
  }

  observer.onStateChange(state);
};

Panel.storeMapper = state => ({ ...state.search });

export default withDispatch(withWindow(Panel));
