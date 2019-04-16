import { library } from "@fortawesome/fontawesome-svg-core";
import { config } from "@fortawesome/fontawesome-svg-core";
import {
  faAngleLeft,
  faAngleRight,
  faInfoCircle,
  faSignOutAlt,
  faPlus,
  faHome,
  faSmile
} from "@fortawesome/free-solid-svg-icons";

export default function() {
  config.autoAddCss = false;
  library.add(
    faAngleLeft,
    faAngleRight,
    faInfoCircle,
    faSignOutAlt,
    faPlus,
    faHome,
    faSmile
  );
}
