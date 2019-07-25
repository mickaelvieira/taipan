import React, { PropsWithChildren } from "react";
import { useTheme } from "@material-ui/core/styles";
import useMediaQuery from "@material-ui/core/useMediaQuery";
import Modal from "./Modal";
import Panel from "./Panel";

interface Props {
  isOpen: boolean;
  close: () => void;
}

export default function Wrapper({
  isOpen,
  close,
  children
}: PropsWithChildren<Props>): JSX.Element {
  const theme = useTheme();
  const matches = useMediaQuery(theme.breakpoints.up("md"));
  const Wrapper = matches ? Modal : Panel;

  return (
    <Wrapper isOpen={isOpen} close={close}>
      {children}
    </Wrapper>
  );
}
