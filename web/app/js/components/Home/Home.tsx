import React from "react";
import Button from "components/ui/Buttons";
import { Link } from "react-router-dom";

interface Props {
  onClickAddBookmark: () => void;
}
export default function Home(props: Props) {
  return (
    <div>
      <ul>
        <li>
          <Button className="btn-link" onClick={props.onClickAddBookmark}>
            Add a bookmark
          </Button>
        </li>
        <li>
          <Link to="/feed">Feed</Link>
        </li>
        <li>
          <Link to="/search">Search</Link>
        </li>
      </ul>
    </div>
  );
}
