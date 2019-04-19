import API, { ContentTypes } from "lib/api";
import withWindow from "components/HoC/withWindow";

const Logout = ({ window }) => {
  const btn = window.document.querySelector(".btn-logout");

  btn.addEventListener("click", event => {
    event.preventDefault();
    API.get("/logout", ContentTypes.HTML).then(response => {
      console.log(response);
      if (response.status === 200) {
        window.location = "/login";
      }
    });
  });
};

export default withWindow(Logout);
