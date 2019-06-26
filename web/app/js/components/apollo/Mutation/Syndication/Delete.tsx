import { DataProxy } from "apollo-cache";
import { Mutation } from "react-apollo";
import { Source } from "../../../../types/syndication";
import mutation from "../../../../services/apollo/mutation/syndication/delete-source.graphql";
import { query, Data as QueryData } from "../../Query/Syndication";
import { removeSource } from "../../helpers/syndication";

interface Data {
  syndication: {
    delete: Source;
  };
}

interface Variables {
  url: string;
}

class ChangeStatusMutation extends Mutation<Data, Variables> {
  static defaultProps = {
    mutation,
    update: (cache: DataProxy, { data }: { data: Data }) => {
      const { delete: source } = data.syndication;
      const prev = cache.readQuery({ query }) as QueryData;
      const result = removeSource(prev.syndication.sources, source);
      cache.writeQuery({
        query,
        data: {
          syndication: {
            ...prev.syndication,
            sources: result
          }
        }
      });
    }
  };
}

export { mutation };

export default ChangeStatusMutation;
