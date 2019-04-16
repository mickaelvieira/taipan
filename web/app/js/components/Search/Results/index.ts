import { wrap } from "dom-element-wrapper";
import Fetching from "../Result/Fetching";
import NoResult from "../Result/NoResult";
import Result from "../Result";

/**
 * @param {Array} terms
 *
 * @returns {Object}
 */
function highlighter(terms) {
  function highlight(text) {
    let str = text;
    const elements = [];
    while (str.length > 0) {
      for (let term of terms) {
        const pos = str.indexOf(term);
        const end = pos + term.length;

        if (pos >= 0) {
          const before = str.substring(0, pos);
          const highlighted = str.substring(pos, end);
          const after = str.substring(end);

          console.log(text);
          // console.log(str);
          // console.log(terms);
          console.log(before);
          console.log(highlighted);
          console.log(after);

          const span = document.createElement("span");
          const strong = document.createElement("strong");

          span.appendChild(document.createTextNode(before));
          strong.appendChild(document.createTextNode(highlighted));

          str = after;

          elements.push(document.createTextNode(before));
          elements.push(strong);
        }
      }

      if (str.length > 0) {
        elements.push(document.createTextNode(str));
        str = "";
      }
    }

    console.log(elements);

    return elements;

    // if (terms.length > 0) {
    //     const pos = str.indexOf(terms[0]);
    //     const end = pos + terms[0].length;
    //
    //     if (pos >= 0) {
    //       const before = str.substring(0, pos);
    //       const highlighted = str.substring(pos, end);
    //       const after = str.substring(end);
    //
    //       console.log(text);
    //       // console.log(str);
    //       // console.log(terms);
    //       console.log(before);
    //       console.log(highlighted);
    //       console.log(after);
    //     }
    // }
    // //

    // const re = new RegExp(`(${terms.join("|")})`, "gi");
    // return terms.length > 0 ? text.replace(re, "<strong>$1</strong>") : text;
  }

  return { highlight };
}

/**
 * @param {Object} state
 *
 * @returns {Node}
 */
const Results = ({ selected, fetching, visible, results, terms }) => {
  const element = wrap("ul", { className: "search-bookmark-results" });
  const hl = highlighter(new Set(terms));

  if (fetching) {
    element.append(Fetching());
  } else if (terms.length > 0 && results.length === 0) {
    element.append(NoResult());
  } else {
    const elements = results.map(({ data: { id, url, title, description } }) =>
      Result({
        id,
        url: new URL(url),
        title,
        description: hl.highlight(description)
      })
    );

    element.append(...elements);
  }

  return element.unwrap();
};

export default Results;
