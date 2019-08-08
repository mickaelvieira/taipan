import React from "react";
import { useQuery } from "@apollo/react-hooks";
import { createMuiTheme, MuiThemeProvider } from "@material-ui/core/styles";
import CssBaseline from "@material-ui/core/CssBaseline";
import Layout from "../Layout";
import Routes from "./Routes";
import getThemeOptions from "../ui/themes";
import Loader from "../ui/Loader";
import { InitQueryData, query } from "../apollo/Query/Init";
import { AppContext } from "../context";

export default function Environment(): JSX.Element {
  const { loading, error, data } = useQuery<InitQueryData, {}>(query);

  if (loading) {
    return <Loader />;
  }

  let user = null;
  let appInfo = { info: { name: "Taipan", version: "0.0.0" } };
  if (data && !error) {
    let { users, app } = data;
    if (users.loggedIn) {
      user = users.loggedIn;
    }
    if (app) {
      appInfo = app;
    }
  }

  const theme = createMuiTheme(getThemeOptions(user ? user.theme : null));

  console.log(theme);

  return (
    <>
      <CssBaseline />
      <MuiThemeProvider theme={theme}>
        <AppContext.Provider value={appInfo.info}>
          <Layout user={user}>
            <Routes />
          </Layout>
        </AppContext.Provider>
      </MuiThemeProvider>
    </>
  );
}
