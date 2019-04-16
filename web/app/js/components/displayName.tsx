import { ComponentType } from "react";

const getDisplayName = <P extends {}>(Component: ComponentType<P>): string =>
  Component.displayName || Component.name || "Component";

const formatDisplayName = <P extends {}>(
  wapper: string,
  WrappedComponent: ComponentType<P>
): string => `${wapper}(${getDisplayName(WrappedComponent)})`;

export { getDisplayName, formatDisplayName };
