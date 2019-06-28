import React from "react";
import { User } from "../../types/users";
import { AppInfo } from "../../types/app";

const AppContext = React.createContext<AppInfo | null>(null);
const UserContext = React.createContext<User | null>(null);
const MessageContext = React.createContext((_: string) => {});

export { AppContext, UserContext, MessageContext };
