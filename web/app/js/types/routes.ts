import { RouteComponentProps } from "react-router";

export type RouteForgotPasswordProps = RouteComponentProps<{}>;
export type RouteResetPasswordProps = RouteComponentProps<{}>;
export type RouteConfirmEmailProps = RouteComponentProps<{}>;
export type RouteSigninProps = RouteComponentProps<{}>;
export type RouteSignupProps = RouteComponentProps<{}>;
export type RouteSearchProps = RouteComponentProps<{ type?: string }>;
export type RouteFeedProps = RouteComponentProps<{}>;
export type RouteSubscriptionProps = RouteComponentProps<{}>;
export type RouteAccountProps = RouteComponentProps<{}>;

export type RoutesProps = RouteSigninProps &
  RouteSignupProps &
  RouteFeedProps &
  RouteSearchProps &
  RouteSubscriptionProps &
  RouteAccountProps &
  RouteConfirmEmailProps &
  RouteResetPasswordProps &
  RouteForgotPasswordProps;
