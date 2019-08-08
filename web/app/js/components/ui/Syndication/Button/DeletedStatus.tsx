import React from "react";
import { useMutation } from "@apollo/react-hooks";
import IconButton from "@material-ui/core/IconButton";
import VisibleIcon from "@material-ui/icons/Visibility";
import HiddenIcon from "@material-ui/icons/VisibilityOff";
import {
  enableMutation,
  disableMutation,
  Data,
  Variables
} from "../../../apollo/Mutation/Syndication/DeletedStatus";
import { Source } from "../../../../types/syndication";

interface Props {
  source: Source;
}

export default React.memo(function VisibilityButton({
  source
}: Props): JSX.Element {
  const { isDeleted } = source;
  const [mutate] = useMutation<Data, Variables>(
    isDeleted ? enableMutation : disableMutation
  );

  return (
    <IconButton
      onClick={() => {
        mutate({
          variables: {
            url: source.url
          }
        });
      }}
    >
      {isDeleted ? (
        <HiddenIcon color="secondary" />
      ) : (
        <VisibleIcon color="primary" />
      )}
    </IconButton>
  );
});
