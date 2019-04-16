import React, { Component } from "react";
import ButtonClose from "components/ui/Buttons/ButtonClose";
import classNames from "classnames";

export interface PanelProps {
  title: string;
  isOpen: boolean;
  onClickClose: () => void;
  children: any;
}

class Panel extends Component<PanelProps, {}> {
  static defaultProps = {
    isOpen: false,
    onClickClose: () => {}
  };

  render() {
    const { children, title, isOpen, onClickClose } = this.props;

    return (
      <div className={classNames("app-panel", { active: isOpen })}>
        <header className="app-panel-header">
          <h3>{title}</h3>
          <ButtonClose onClick={onClickClose} />
        </header>
        <section>{children}</section>
      </div>
    );
  }
}

export default Panel;
