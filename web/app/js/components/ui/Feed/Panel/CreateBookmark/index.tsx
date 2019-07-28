import React, { useState } from "react";
import { Document } from "../../../../../types/document";
import { Bookmark } from "../../../../../types/bookmark";
import FormDocument from "./FormDocument";
import FormBookmark from "./FormBookmark";
import Panel from "../../../Panel";

interface Props {
  isOpen: boolean;
  setIsPanelOpen: (isOpen: boolean) => void;
  onBookmarkCreated: (bookmark: Bookmark) => void;
}

export default function CreateBookmark({
  isOpen,
  setIsPanelOpen,
  onBookmarkCreated
}: Props): JSX.Element {
  const [document, setDocument] = useState<Document | null>(null);
  const prev = (): void => {
    if (document) {
      setDocument(null);
    } else {
      setIsPanelOpen(false);
    }
  };

  return (
    <Panel title="Bookmark a webpage" isOpen={isOpen} prev={prev}>
      {!document && <FormDocument onFetchDocument={setDocument} />}
      {document && (
        <FormBookmark
          document={document}
          onFinish={bookmark => {
            setDocument(null);
            onBookmarkCreated(bookmark);
          }}
        />
      )}
    </Panel>
  );
}
