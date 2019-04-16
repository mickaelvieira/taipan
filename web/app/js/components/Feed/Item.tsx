import React from "react";
import classNames from "classnames";
import { isEqual } from "lodash";
import Image from "./Image";
import { Bookmark } from "../../types/bookmark";
import Domain from "../ui/Domain";

interface Props {
  item: Bookmark;
}

export default React.memo(function Item({ item }: Props) {
  const {
    title,
    description,
    is_read,
    links: { image }
  } = item;

  const url = new URL(item.url);

  return (
    <li className={is_read ? "read" : "unread"}>
      <div className="bookmark-block">
        <header className="bookmark-header">
          <a
            className={classNames("bookmark-title", {
              truncate: !title
            })}
            href={`${url}`}
            title={`${url}`}
            target="_blank"
          >
            {title ? title : `${url}`}
          </a>
          <br />
          <Domain value={url} />
        </header>
        <article className="bookmark-entry">
          <Image source={image} alt={title} />
          <section className="bookmark-description">{description}</section>
        </article>
      </div>
    </li>
  );
}, isEqual);
