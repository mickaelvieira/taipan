import React, { Component, ErrorInfo } from "react";
import App from "./App";
import ErrorMessage from "./Error";

interface State {
  error: Error | null;
  info: ErrorInfo | null;
}

export default class Bootstrap extends Component<{}, State> {
  state = {
    error: null,
    info: null,
  };

  componentDidCatch(error: Error, info: ErrorInfo): void {
    this.setState({
      error: error,
      info: info,
    });
  }

  render(): JSX.Element {
    const { info, error } = this.state;
    return info ? <ErrorMessage info={info} error={error} /> : <App />;
  }
}
