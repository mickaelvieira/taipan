import React, { useState } from "react";
import { useMutation } from "@apollo/react-hooks";
import Checkbox from "@material-ui/core/Checkbox";
import {
  Data,
  Variables,
  tagMutation,
  untagMutation
} from "../../../apollo/Mutation/Syndication/Tag";
import { Source, Tag } from "../../../../types/syndication";

interface Props {
  tag: Tag;
  ids: string[];
  source: Source;
  className?: string;
}

export default React.memo(function Tag({
  tag,
  ids,
  source,
  className
}: Props): JSX.Element {
  const hasTag = ids.includes(tag.id);
  const [isChecked, setIsChecked] = useState(hasTag);
  const [mutate] = useMutation<Data, Variables>(
    hasTag ? untagMutation : tagMutation
  );

  return (
    <Checkbox
      onChange={() => {
        setIsChecked(!isChecked);
        mutate({
          variables: {
            sourceId: source.id,
            tagId: tag.id
          }
        });
      }}
      checked={isChecked}
      inputProps={{ "aria-label": "Add or remove tag" }}
      classes={{ root: className }}
    />
  );
});
