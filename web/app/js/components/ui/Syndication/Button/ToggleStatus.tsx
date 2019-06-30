import React, { useState } from "react";
import Switch from "@material-ui/core/Switch";
import ChangeSourceStatusMutation, {
  enableSourceMutation,
  disableSourceMutation
} from "../../../apollo/Mutation/Syndication/ChangeStatus";
import { Source } from "../../../../types/syndication";

interface Props {
  source: Source;
}

export default React.memo(function ToggleStatus({
  source
}: Props): JSX.Element {
  const { isPaused } = source;
  const [isChecked, setIsChecked] = useState(!isPaused);

  return (
    <ChangeSourceStatusMutation
      mutation={isPaused ? enableSourceMutation : disableSourceMutation}
    >
      {mutate => {
        return (
          <Switch
            edge="end"
            onChange={() => {
              setIsChecked(!isChecked);
              mutate({
                variables: {
                  url: source.url
                }
              });
            }}
            checked={isChecked}
            inputProps={{
              "aria-labelledby": "switch-list-label-source"
            }}
          />
        );
      }}
    </ChangeSourceStatusMutation>
  );
});
