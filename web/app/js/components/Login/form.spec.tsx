import Form from "./form";
import simulant from "simulant";

describe("Login Form", () => {
  document.body.innerHTML =
    '<form method="post" class="form-login" novalidate="">' +
    '<div class="form-row">' +
    '<label for="username">Username</label>' +
    '<input placeholder="Username or email" id="username" name="username" class="input-username" type="text" tabindex="20" autocomplete="off" autocorrect="off" spellcheck="false" autofocus="" required="">' +
    '<small data-for="username" class="form-feedback">' +
    "Please provide your username or email" +
    "</small>" +
    '<div style="adjuster">' +
    '<label for="password">Password</label>' +
    '<a href="/forgot-password" class="link-forgot-password" tabindex="23">Forgot password?</a>' +
    "</div>" +
    '<input placeholder="Your password" id="password" name="password" class="input-password" type="password" tabindex="21" required="">' +
    '<small data-for="password" class="form-feedback">' +
    "Please provide your password" +
    "</small>" +
    '<div class="form-actions">' +
    '<button class="btn-login" tabindex="22">Sign in</button>' +
    "</div>" +
    "</div>" +
    "</form>";

  let form, elem;

  beforeEach(() => {
    elem = document.querySelector(".form-login");
    form = new Form(elem);
  });

  test("is initialized", () => {
    expect(form).not.toBe(undefined);
  });
});
