import React, { useState } from "react";
import Checkbox from "@material-ui/core/Checkbox";
import ChangeStatusMutation, {
  subscribeMutation,
  unsubscribeMutation
} from "../../../apollo/Mutation/Subscriptions/Status";
import { Subscription } from "../../../../types/subscription";

interface Props {
  subscription: Subscription;
}

export default React.memo(function StatusCheckbox({
  subscription
}: Props): JSX.Element {
  const { isSubscribed } = subscription;
  const [isChecked, setIsChecked] = useState(isSubscribed);

  return (
    <ChangeStatusMutation
      mutation={isSubscribed ? unsubscribeMutation : subscribeMutation}
    >
      {mutate => {
        return (
          <Checkbox
            onChange={() => {
              setIsChecked(!isChecked);
              mutate({
                variables: {
                  url: subscription.url
                }
              });
            }}
            checked={isChecked}
            inputProps={{
              "aria-labelledby": "switch-list-label-subscription"
            }}
          />
        );
      }}
    </ChangeStatusMutation>
  );
});
