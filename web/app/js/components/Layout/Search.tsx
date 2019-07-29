import React, { PropsWithChildren } from "react";
import MainLayout from "./Layout";
import MainContent from "./Content";

export default function SearchLayout({
  children
}: PropsWithChildren<{}>): JSX.Element {
  return (
    <MainLayout>
      {() => (
        <>
          <MainContent>{children}</MainContent>
        </>
      )}
    </MainLayout>
  );
}
