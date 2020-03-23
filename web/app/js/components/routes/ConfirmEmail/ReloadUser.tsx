import { useQuery } from "@apollo/react-hooks";
import { Data, query } from "../../apollo/Query/LoggedInUser";

export default function ReloadUser(): null {
  useQuery<Data, {}>(query, {
    fetchPolicy: "network-only",
  });
  return null;
}
