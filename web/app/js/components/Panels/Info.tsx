import React, { useEffect, useState, useMemo } from "react";
import { Dispatch } from "redux";
import { connect } from "react-redux";
import Panel from "components/ui/Panel";
import { refreshBookmark, removeBookmark } from "store/actions/feed";
import { fetchBookmarkHistory } from "store/actions/bookmarks";
import { Bookmark } from "types/bookmark";
import { PanelProps } from "components/ui/Panel";
import { ReduxAction } from "store/actions/types";
import YesNo from "../ui/YesNo";
import Domain from "../ui/Domain";
import Datetime from "../ui/Datetime";
import Info from "../ui/Info";

function getEntries(item: Bookmark) {
  const { history } = item;
  const lastEntry = history && history.items.pop();
  const fetchedAt = lastEntry ? lastEntry.created_at : undefined;

  const entries = [
    <Datetime label="Added at" value={item.added_at} />,
    <Datetime label="Accessed at" value={item.accessed_at} />,
    <Datetime label="Created at" value={item.created_at} />,
    <Datetime label="Updated at" value={item.updated_at} />,
    <Datetime label="Last fetched at" value={fetchedAt} />,
    <Info label="Charset" value={item.charset} />,
    <YesNo label="Read" value={item.is_read} />,
    <YesNo label="Pending" value={item.is_pending} />,
    <YesNo label="Fetching" value={item.is_fetching} />
  ];

  return entries.map((entry, index) => <li key={index}>{entry}</li>);
}

interface Props {
  title: string;
  item: Bookmark;
  onClickClose: () => void;
  removeBookmark: (item: Bookmark) => any;
  refreshBookmark: (item: Bookmark) => any;
  fetchBookmarkHistory: (item: Bookmark) => any;
}

type ComponentProps = Props & Partial<PanelProps>;

function PanelInfo({
  item,
  removeBookmark,
  fetchBookmarkHistory,
  refreshBookmark,
  isOpen,
  ...panelProps
}: ComponentProps) {
  const [isLoading, setIsLoading] = useState(false);
  const url = new URL(item.url);

  useEffect(() => {
    if (item && !("history" in item) && !isLoading && isOpen) {
      setIsLoading(true);
      fetchBookmarkHistory(item).then(() => setIsLoading(false));
    }
  }, [item, isLoading, isOpen]);

  const entries = useMemo(() => getEntries(item), [item]);

  function onClickDelete() {
    if (window.confirm("Are you sure you want to remove this bookmark?")) {
      removeBookmark(item);
    }
  }

  function onClickRefresh() {
    refreshBookmark(item);
  }

  return (
    <Panel {...panelProps} isOpen={isOpen}>
      {item.title && <h5>{item.title}</h5>}
      <Domain value={url} />
      <ul className="bookmark-info">{entries}</ul>
      <div className="btn-actions-container">
        <button type="button" onClick={onClickRefresh}>
          <span>Refresh</span>
        </button>
        <button type="button" className="btn-delete" onClick={onClickDelete}>
          <span>Delete</span>
        </button>
      </div>
    </Panel>
  );
}

const mapDispatchToProps = (dispatch: Dispatch<ReduxAction>) => ({
  refreshBookmark: (item: Bookmark) => dispatch(refreshBookmark(item)),
  removeBookmark: (item: Bookmark) => dispatch(removeBookmark(item)),
  fetchBookmarkHistory: (item: Bookmark) => dispatch(fetchBookmarkHistory(item))
});

export default connect(
  undefined,
  mapDispatchToProps
)(PanelInfo);
