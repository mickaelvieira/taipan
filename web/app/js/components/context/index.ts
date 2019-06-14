import React from "react";
import { User } from "../../types/users";

const UserContext = React.createContext<User | null>(null);
const MessageContext = React.createContext((_: string) => {});
const NewsContext = React.createContext((_: string) => {});

export { UserContext, MessageContext, NewsContext };
