import Observer from "./Form/Observer";
import withWindow from "components/HoC/withWindow";
import withDispatch from "components/HoC/withDispatch";

let observer;

const FormSearch = ({ window, dispatch, ...state }) => {
  if (!observer) {
    observer = new Observer(window, dispatch);
  }

  observer.onStateChange(state);
};

FormSearch.storeMapper = state => ({ ...state.search });

export default withDispatch(withWindow(FormSearch));
