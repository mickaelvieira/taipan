import { RouteComponentProps } from "react-router";

export type RouteLoginProps = RouteComponentProps<{}>;
export type RouteSearchProps = RouteComponentProps<{ type?: string }>;
export type RouteFeedProps = RouteComponentProps<{}>;
export type RouteSubscriptionProps = RouteComponentProps<{}>;
export type RouteAccountProps = RouteComponentProps<{}>;

export type RoutesProps = RouteLoginProps &
  RouteFeedProps &
  RouteSearchProps &
  RouteSubscriptionProps &
  RouteAccountProps;
