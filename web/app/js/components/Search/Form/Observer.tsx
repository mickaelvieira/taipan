import {
  hideSearchResults,
  launchSearch,
  prepareSearch,
  selectSearchResult,
  showSearchResults,
  updateSearchTerms
} from "../../../actions";

/**
 * @param {Array}  results
 * @param {Object} selected
 *
 * @returns {Object}
 */
function prev(results, selected) {
  if (!selected) {
    return results[results.length - 1];
  }

  const index = results.indexOf(selected);
  let prev = index - 1;

  if (prev < 0) {
    prev = results.length - 1;
  }

  return results[prev];
}

/**
 * @param {Array}  results
 * @param {Object} selected
 *
 * @returns {Object}
 */
function next(results, selected) {
  if (!selected) {
    return results[0];
  }

  const index = results.indexOf(selected);
  let next = index + 1;

  if (next >= results.length) {
    next = 0;
  }

  return results[next];
}

const words = (function() {
  let char,
    str = "";
  const words = [];
  const chars = "https://".split("");

  for (let key in chars) {
    char = chars[key];
    str += char;
    words.push(str);
  }

  return words;
})();

/**
 * @param {Event} event
 * @param {Array} terms
 *
 * @returns {Boolean}
 */
function shouldTriggerSearch({ keyCode }, terms) {
  /**
   * @param {String} term
   *
   * @returns {Boolean}
   */
  function isValidTerm(term) {
    return words.indexOf(term) === -1;
  }

  // 9  - tab
  // 13 - enter
  // 37 - left arrow
  // 38 - up arrow
  // 39 - right arrow
  // 40 - down arrow

  /**
   * @param {Number} code
   *
   * @returns {Boolean}
   */
  function isValidKey(code) {
    return [9, 13, 37, 38, 39, 40].indexOf(code) === -1;
  }

  return (
    keyCode && isValidKey(keyCode) && terms.length > 0 && isValidTerm(terms[0])
  );
}

/**
 * @param {Object} selected
 * @param {Array}  terms
 *
 * @returns {String}
 */
function getInputTextValue(selected, terms) {
  if (selected) {
    return selected.title;
  }

  if (terms.length > 0) {
    return terms.join(" ");
  }

  return "";
}

function parseInputTextValue(value) {
  return value
    .split(/\s/)
    .filter(term => term !== "")
    .map(term => term.replace(/^\s+/, "").replace(/\s+$/, ""));
}

export default class Observer {
  constructor(window, dispatch) {
    this.state = null;
    this.window = window;
    this.dispatch = dispatch;

    this.updateButtonState = this.updateButtonState.bind(this);

    this.form = this.window.document.querySelector(".search-bookmark-form");
    this.text = this.form.querySelector(".search-bookmark-form-input-terms");
    this.btn = this.form.querySelector(".search-bookmark-form-btn-submit");

    this.form.addEventListener("submit", function(event) {
      event.preventDefault();
    });

    // this.text.addEventListener("click", function(event) {
    //   // dispatch(prepareSearch());
    //   // dispatch(launchSearch());
    // });

    let timer;
    this.text.addEventListener("keyup", function(event) {
      const duration = 300;
      const terms = parseInputTextValue(this.value);

      if (shouldTriggerSearch(event, terms)) {
        // @TODO debounce this function
        if (timer) {
          window.clearTimeout(timer);
        }

        timer = window.setTimeout(function() {
          dispatch(updateSearchTerms(terms));
          dispatch(prepareSearch());
          dispatch(launchSearch());
        }, duration);
      }
    });

    this.text.addEventListener("keydown", ({ keyCode }) => {
      let actions;
      const { visible, results, selected } = this.state;

      if (results.length > 0) {
        switch (keyCode) {
          case 38:
            actions = visible
              ? selectSearchResult(prev(results, selected))
              : showSearchResults();
            break;
          case 40:
            actions = visible
              ? selectSearchResult(next(results, selected))
              : showSearchResults();
            break;
          case 13:
            actions = hideSearchResults();
            break;
          default:
        }

        if (actions) {
          dispatch(actions);
        }
      }
    });
  }

  updateButtonState() {
    const { terms } = this.state;
    this.btn.disabled = terms.length === 0;
  }

  onStateChange(state) {
    this.state = state;
    this.updateButtonState();
  }
}
