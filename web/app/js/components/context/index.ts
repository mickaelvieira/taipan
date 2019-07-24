import React from "react";
import { User } from "../../types/users";
import { AppInfo } from "../../types/app";
import { MessageInfo } from "../../types";
import FeedsUpdater from "../apollo/helpers/feeds-updater";
import FeedsMutator from "../apollo/helpers/feeds-mutator";

const ClientContext = React.createContext<string>("");
const AppContext = React.createContext<AppInfo | null>(null);
const UserContext = React.createContext<User | null>(null);
const MessageContext = React.createContext((_: MessageInfo | null) => {});
const FeedsCacheContext = React.createContext<FeedsUpdater | null>(null);
const FeedsContext = React.createContext<FeedsMutator | null>(null);

export {
  ClientContext,
  AppContext,
  UserContext,
  MessageContext,
  FeedsCacheContext,
  FeedsContext
};
