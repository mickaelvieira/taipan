type Errors = ReadonlyArray<string>;

interface Data {
  readonly [key: string]: string;
}

interface OnSubmit {
  (data: Data): Promise<void>;
}

export default class Form {
  form: HTMLFormElement;
  callback: OnSubmit;
  feedbacks: NodeListOf<HTMLElement>;

  constructor(form: HTMLFormElement, callback: OnSubmit) {
    this.form = form;
    this.callback = callback;
    this.feedbacks = form.querySelectorAll(".form-feedback");
    this.form.addEventListener("submit", this.onSubmit);

    const username = form.querySelector(".input-username");
    if (username) {
      username.addEventListener("keyup", this.onUserNameChange);
    }
  }

  onSubmit = async (event: Event) => {
    event.preventDefault();

    const [data, errors] = this.validate();

    this.handleMessages(errors);

    if (errors.length === 0) {
      this.disable();
      await this.callback(data);
      this.enable();
    } else {
      this.form.elements[errors[0]].focus();
    }
  };

  onUserNameChange = (event: Event) => {
    if (event.target && event.target.value) {
      event.target.value = event.target.value.toLowerCase();
    }
  };

  handleMessages(errors: Errors) {
    Array.from(this.feedbacks).forEach(element => {
      const fn =
        element.dataset.for && errors.includes(element.dataset.for)
          ? "add"
          : "remove";
      element.classList[fn]("visible");
    });
  }

  validate(): [Data, Errors] {
    const fields = Array.from(new FormData(this.form).entries());
    const data: Data = Object.freeze(
      fields.reduce(
        (carry, [name, value]) => ({ ...carry, [name]: `${value}`.trim() }),
        {}
      )
    );
    const errors: Errors = Object.freeze(
      Object.keys(data).filter(key => data[key] === "")
    );

    return [data, errors];
  }

  enable() {
    Array.from(this.form.elements).forEach(
      element => (element.disabled = false)
    );
  }

  disable() {
    Array.from(this.form.elements).forEach(
      element => (element.disabled = true)
    );
  }
}
