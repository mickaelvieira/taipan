import React from "react";
import { User } from "../../types/users";
import { AppInfo } from "../../types/app";

const ClientContext = React.createContext<string>("");
const AppContext = React.createContext<AppInfo | null>(null);
const UserContext = React.createContext<User | null>(null);
const MessageContext = React.createContext((_: string) => {});

export { ClientContext, AppContext, UserContext, MessageContext };
