import { ReduxActionWithPayload, UserActions } from "store/actions/types";
import { parseCollection } from "collection";
import { User } from "types/users";
import { CollectionResponse } from "collection/types";
import db from "services/repository/user";

export type UserState = User | null;

export const defaultState = null;

export default function(
  state: UserState = defaultState,
  action: ReduxActionWithPayload
): UserState {
  switch (action.type) {
    case UserActions.UPDATE: {
      const payload = action.payload as CollectionResponse;
      const collection = parseCollection(payload.collection);
      const user = collection.items[0];

      db.upsert(user);

      return { ...user };
    }
  }

  return state;
}
