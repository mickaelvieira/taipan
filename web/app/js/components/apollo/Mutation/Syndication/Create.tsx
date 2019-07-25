import { DataProxy } from "apollo-cache";
import { Mutation } from "react-apollo";
import { Source } from "../../../../types/syndication";
import mutation from "../../graphql/mutation/syndication/create.graphql";
import { query, Data as QueryData } from "../../Query/Syndication";
import { addSource } from "../../helpers/syndication";

interface Data {
  syndication: {
    source: Source;
  };
}

interface Variables {
  url: string;
}

class CreateSourceMutation extends Mutation<Data, Variables> {
  static defaultProps = {
    mutation,
    update: (cache: DataProxy, { data }: { data: Data }) => {
      const { source } = data.syndication;
      const prev = cache.readQuery({ query }) as QueryData;
      const result = addSource(prev.syndication.sources, source);
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

export default CreateSourceMutation;
