import React, { useState, useCallback } from "react";
import Item from "./Item";
import PanelInfo from "components/Panels/Info";
import Navigation from "./Navigation";
import Pagination, { PaginationProps } from "./Pagination";
import InfiniteFeed from "./InfiniteFeed";

export default React.memo(function Feed() {
  const [isPanelInfoOpen, setIsPanelInfoOpen] = useState(false);

  const openPanelInfo = useCallback(() => setIsPanelInfoOpen(true), []);
  const closePanelInfo = useCallback(() => setIsPanelInfoOpen(false), []);

  return (
    <Pagination>
      {({ pagination, isNearTheEnd }: PaginationProps) => (
        <>
          <InfiniteFeed isNearTheEnd={isNearTheEnd}>
            <section id="feed" className="feed">
              <div className="items">
                <ul>
                  {pagination.current && <Item item={pagination.current} />}
                </ul>
              </div>
            </section>
          </InfiniteFeed>
          {pagination.current && (
            <PanelInfo
              title="Info"
              item={pagination.current}
              isOpen={isPanelInfoOpen}
              onClickClose={closePanelInfo}
            />
          )}
          <Navigation pagination={pagination} onClickInfo={openPanelInfo} />
        </>
      )}
    </Pagination>
  );
});
