import API, { ContentTypes } from "lib/api";
import Form from "./form";

export default function Login() {
  const form: HTMLFormElement | null = window.document.querySelector(
    ".form-login"
  );

  if (!form) {
    return;
  }

  new Form(form, async function(data) {
    const response = await API.post("/login", data, ContentTypes.HTML);
    if (response.status === 200) {
      window.location.href = "/";
    }
  });
}
