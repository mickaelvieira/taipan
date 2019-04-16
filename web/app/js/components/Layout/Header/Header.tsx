import React from "react";
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import { Link } from "react-router-dom";
import Button from "components/ui/Buttons";

interface Props {
  onClickAddBookmark: () => void;
}

export default React.memo(function Header({ onClickAddBookmark }: Props) {
  console.log("render header");
  return (
    <header className="app-header">
      <Link to="/" className="btn-home">
        <FontAwesomeIcon icon="home" />
      </Link>
      <Button className="btn-add-bookmark" onClick={onClickAddBookmark}>
        <FontAwesomeIcon icon="plus" />
      </Button>
      <a className="btn-logout" href="#">
        <FontAwesomeIcon icon="sign-out-alt" />
      </a>
    </header>
  );
});
