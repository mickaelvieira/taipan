import React from "react";
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import classNames from "classnames";
import Button from "components/ui/Buttons";
import { Pagination } from "./Pagination";

interface Props {
  pagination: Pagination;
  onClickInfo: () => void;
}

export default function FeedNavigation({
  pagination: { prev, next, toPrev, toNext },
  onClickInfo
}: Props) {
  return (
    <footer className="app-footer">
      <nav className="btn-container">
        <Button
          onClick={() => toPrev()}
          className={classNames("btn-next", { disabled: !prev })}
        >
          <FontAwesomeIcon icon="angle-left" />
        </Button>
        <Button onClick={onClickInfo} className="btn-info">
          <FontAwesomeIcon icon="info-circle" />
        </Button>
        <Button
          onClick={() => toNext()}
          className={classNames("btn-next", { disabled: !next })}
        >
          <FontAwesomeIcon icon="angle-right" />
        </Button>
      </nav>
    </footer>
  );
}
