import React, { useState } from "react";
import { useMutation } from "@apollo/react-hooks";
import Checkbox from "@material-ui/core/Checkbox";
import {
  Data,
  Variables,
  subscribeMutation,
  unsubscribeMutation,
} from "../../../apollo/Mutation/Subscriptions/Status";
import { Subscription } from "../../../../types/subscription";

interface Props {
  subscription: Subscription;
}

export default React.memo(function StatusCheckbox({
  subscription,
}: Props): JSX.Element {
  const { isSubscribed } = subscription;
  const [isChecked, setIsChecked] = useState(isSubscribed);
  const [mutate] = useMutation<Data, Variables>(
    isSubscribed ? unsubscribeMutation : subscribeMutation
  );

  return (
    <Checkbox
      onChange={() => {
        setIsChecked(!isChecked);
        mutate({
          variables: {
            url: subscription.url,
          },
        });
      }}
      checked={isChecked}
      inputProps={{ "aria-label": "Subscribe or unsubscribe" }}
    />
  );
});
