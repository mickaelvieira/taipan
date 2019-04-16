import React, { Component, Fragment } from "react";
import API, { ContentTypes } from "lib/api";
import Loader from "components/ui/Loader";

interface Props {
  alt: string;
  source: string;
}

interface State {
  blob: string | null;
  source: string | null;
  isLoading: boolean;
}

class Image extends Component<Props, State> {
  state: State = {
    blob: null,
    source: null,
    isLoading: false
  };

  controller: AbortController | null = null;

  image = React.createRef<HTMLImageElement>();

  static getDerivedStateFromProps(props: Props, state: State) {
    if (props.source !== state.source) {
      return {
        blob: null,
        source: props.source
      };
    }
    return null;
  }

  showImage = () => {
    setTimeout(() => {
      if (this.image.current) {
        this.image.current.classList.add("active");
      }
    }, 100);
  };

  hideImage = () => {
    if (this.image.current) {
      this.image.current.classList.remove("active");
    }
  };

  async componentDidUpdate() {
    const { blob, isLoading } = this.state;
    const { source } = this.props;

    if (!blob && source && !isLoading) {
      this.hideImage();
      this.setState({
        isLoading: true
      });

      // if (this.controller) {
      //   this.controller.abort();
      // }

      this.controller = new AbortController();

      const signal = this.controller.signal;
      signal.addEventListener("abort", () => {
        console.log("aborted!");
      });

      try {
        const response = await API.get(source, ContentTypes.IMAGE, { signal });
        // this.controller.abort();
        const blob = response.ok ? await response.blob() : null;

        if (blob) {
          this.showImage();
          this.setState({
            isLoading: false,
            blob: URL.createObjectURL(blob)
          });
        }
        // } else {
        //   this.setState({
        //     isLoading: false,
        //     blob: null,
        //     source: null
        //   });
        // }
      } catch (err) {
        console.log(err);
        this.setState({
          isLoading: false
        });
      }
    }
  }

  componentWillUnmount() {
    // @TODO request cancelation
    if (this.controller) {
      this.controller.abort();
    }
  }

  render() {
    const { alt } = this.props;
    const { blob, isLoading } = this.state;

    return (
      <Fragment>
        {isLoading && <Loader />}
        {blob && (
          <section className="bookmark-image">
            <img ref={this.image} src={blob} alt={alt} />
          </section>
        )}
      </Fragment>
    );
  }
}

export default Image;
