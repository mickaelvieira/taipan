import React, { useState } from "react";
import Checkbox from "@material-ui/core/Checkbox";
import ChangeSourceStatusMutation, {
  enableSourceMutation,
  disableSourceMutation
} from "../../../apollo/Mutation/Syndication/Status";
import { Source } from "../../../../types/syndication";

interface Props {
  source: Source;
}

export default React.memo(function StatusCheckbox({ source }: Props): JSX.Element {
  const { isPaused } = source;
  const [isChecked, setIsChecked] = useState(!isPaused);

  return (
    <ChangeSourceStatusMutation
      mutation={isPaused ? enableSourceMutation : disableSourceMutation}
    >
      {mutate => {
        return (
          <Checkbox
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
