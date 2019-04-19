import React, { PureComponent, FormEvent } from "react";
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import { addFeedItems } from "store/actions/feed";
import { connect } from "react-redux";
import Panel from "components/ui/Panel";
import API from "lib/api";
import { ReduxAction } from "store/actions/types";
import { Dispatch } from "redux";
import { CollectionResponse, Template } from "collection/types";
import { PanelProps } from "components/ui/Panel";
import { RootState } from "store/reducer/default";
import { IndexLinks } from "types";

interface TemplateData {
  template: Template;
}

const getTemplate = (value: string, name: string = "url"): TemplateData => ({
  template: {
    data: [
      {
        name,
        value
      }
    ]
  }
});

interface Props {
  title: string;
  links: IndexLinks;
  onClickClose: () => void;
  addWallItems: (data: CollectionResponse) => any;
}

type FormProps = Props & Partial<PanelProps>;

class PanelAddBookmark extends PureComponent<FormProps> {
  input = React.createRef<HTMLInputElement>();

  static defaultProps = {
    title: "Add"
  };

  onSubmit = async (event: FormEvent) => {
    event.preventDefault();

    const { edit: createUrl } = this.props.links;

    if (!createUrl) {
      throw new Error("create link is not defined");
    }

    if (this.input.current && this.input.current.value.indexOf("http") === 0) {
      const value = this.input.current.value;

      this.input.current.value = "";
      this.props.onClickClose();

      const data = await API.post(createUrl, getTemplate(value));

      this.props.addWallItems(data);
    }
  };

  render() {
    console.log("render add bookmark panel");
    return (
      <Panel {...this.props}>
        <div className="add-bookmark-form-container">
          <form
            method="post"
            className="add-bookmark-form"
            onSubmit={this.onSubmit}
          >
            <input
              placeholder="https://"
              id="bookmark"
              name="bookmark"
              tabIndex={-1}
              className="add-bookmark-form-input-url"
              type="text"
              autoComplete="off"
              autoFocus
              ref={this.input}
            />
            <button className="add-bookmark-form-btn-submit">
              <FontAwesomeIcon icon="smile" />
            </button>
          </form>
        </div>
      </Panel>
    );
  }
}

const mapStateToProps = (state: RootState) => ({
  links: state.index.links
});

const mapDispatchToProps = (dispatch: Dispatch<ReduxAction>) => ({
  addWallItems: (data: CollectionResponse) => dispatch(addFeedItems(data))
});

export default connect(
  mapStateToProps,
  mapDispatchToProps
)(PanelAddBookmark);
