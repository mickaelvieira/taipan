import { RouteComponentProps } from "react-router";

export type RouteSearchProps = RouteComponentProps<{ type?: string }>;
export type RouteFeedProps = RouteComponentProps<{}>;
export type RoutesProps = RouteFeedProps & RouteSearchProps;
