export default class Observer {
  constructor(window, dispatch) {
    this.state = null;
    this.window = window;
    this.dispatch = dispatch;

    this.panel = this.window.document.querySelector(".search-bookmark-panel");
    this.results = this.panel.querySelector(".search-bookmark-results");
    this.btnClose = this.panel.querySelector(".btn-close");

    this.panel.addEventListener("click", event => this.show(event));
    this.btnClose.addEventListener("click", event => this.hide(event));
  }

  isActive() {
    return this.panel.classList.contains("active");
  }

  show(event) {
    event.preventDefault();

    if (!event.target.classList.contains("search-bookmark-panel")) {
      return;
    }

    if (this.isActive()) {
      return;
    }

    const contained = () => {
      this.window.document.documentElement.classList.add("contained");
      this.panel.querySelector(".search-bookmark-form-input-terms").focus();
      this.panel.removeEventListener("transitionend", contained);
    };

    this.panel.addEventListener("transitionend", contained);
    this.panel.classList.add("active");
  }

  hide(event) {
    event.preventDefault();

    this.window.document.documentElement.classList.remove("contained");
    this.panel.classList.remove("active");
  }

  onStateChange(state) {
    this.state = state;
  }
}
