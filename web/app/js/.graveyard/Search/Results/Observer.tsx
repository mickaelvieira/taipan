import { hideSearchResults } from "../../../actions";

export default class Observer {
  constructor(window, dispatch) {
    this.state = null;
    this.window = window;
    this.dispatch = dispatch;

    this.window.document.addEventListener("click", event =>
      this.hideIfVisible(event)
    );
    this.window.addEventListener("scroll", event => this.hideIfVisible(event));
  }

  hideIfVisible() {
    if (!this.state || this.state.visible) {
      this.dispatch(hideSearchResults());
    }
  }

  onStateChange(state) {
    this.state = state;
    this.hideIfVisible();
  }
}
