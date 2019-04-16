import React from "react";
import Button from "./Button";
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";

const ButtonClose = ({ onClick }) => (
  <Button onClick={onClick} className="btn-close">
    <FontAwesomeIcon icon="angle-left" />
  </Button>
);

export default ButtonClose;
