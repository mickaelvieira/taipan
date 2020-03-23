import React from "react";
import { useMutation } from "@apollo/react-hooks";
import IconButton from "@material-ui/core/IconButton";
import PausedIcon from "@material-ui/icons/Pause";
import PlayIcon from "@material-ui/icons/PlayArrow";
import {
  Data,
  Variables,
  pauseMutation,
  resumeMutation,
} from "../../../apollo/Mutation/Syndication/PausedStatus";
import { Source } from "../../../../types/syndication";

interface Props {
  source: Source;
}

export default React.memo(function StatusButton({
  source,
}: Props): JSX.Element {
  const { isPaused } = source;
  const [mutate] = useMutation<Data, Variables>(
    isPaused ? resumeMutation : pauseMutation
  );

  return (
    <IconButton
      onClick={() => {
        mutate({
          variables: {
            url: source.url,
          },
        });
      }}
    >
      {isPaused ? (
        <PausedIcon color="secondary" />
      ) : (
        <PlayIcon color="primary" />
      )}
    </IconButton>
  );
});
