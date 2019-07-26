import React from "react";
import IconButton from "@material-ui/core/IconButton";
import PausedIcon from "@material-ui/icons/Pause";
import PlayIcon from "@material-ui/icons/PlayArrow";
import ChangeStatusMutation, {
  pauseMutation,
  resumeMutation
} from "../../../apollo/Mutation/Syndication/PausedStatus";
import { Source } from "../../../../types/syndication";

interface Props {
  source: Source;
}

export default React.memo(function StatusButton({
  source
}: Props): JSX.Element {
  const { isPaused } = source;

  return (
    <ChangeStatusMutation mutation={isPaused ? resumeMutation : pauseMutation}>
      {mutate => {
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
            {isPaused ? (
              <PausedIcon color="secondary" />
            ) : (
              <PlayIcon color="primary" />
            )}
          </IconButton>
        );
      }}
    </ChangeStatusMutation>
  );
});
