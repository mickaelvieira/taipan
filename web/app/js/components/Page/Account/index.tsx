import React, { useContext } from "react";
import { UserContext } from "../../context";
import Layout from "../../Layout/Account";
import ScrollToTop from "../../ui/ScrollToTop";
import Prolile from "./Profile";

export default function Account(): JSX.Element | null {
  const user = useContext(UserContext);
  if (!user) {
    return null;
  }

  return (
    <Layout>
      <ScrollToTop>
        <Prolile user={user} />
      </ScrollToTop>
    </Layout>
  );
}
