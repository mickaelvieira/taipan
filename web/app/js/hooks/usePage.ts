import { Page } from "../helpers/navigation";

export default function useNavigation(): Page {
  return new Page(`${window.location}`);
}
