import React, { useCallback, useState, useContext } from "react";
import { debounce } from "lodash";
import { makeStyles } from "@material-ui/core/styles";
import IconButton from "@material-ui/core/IconButton";
import InputBase from "@material-ui/core/InputBase";
import SearchIcon from "@material-ui/icons/Search";
import CloseIcon from "@material-ui/icons/Close";
import { UserContext } from "../../context";
import Results from "./Results";
import Empty from "./Empty";
import Tags from "../../ui/Tags";

const useStyles = makeStyles(({ palette }) => ({
  search: {
    display: "flex",
    margin: 16,
    borderBottom: `1px solid  ${palette.grey[500]}`
  },
  options: {
    margin: 16,
    alignItems: "center"
  },
  radioLabel: {
    marginRight: 16
  },
  radioGroup: {
    flexDirection: "row"
  }
}));

export default function Search(): JSX.Element | null {
  const classes = useStyles();
  const user = useContext(UserContext);
  const [terms, setTerms] = useState<string[]>([]);
  const [tags, setTags] = useState<string[]>([]);
  const [value, setValue] = useState("");
  const debouncedDispatch = useCallback(debounce(setTerms, 400), []);
  const onChange = useCallback(
    (input: string, debounced = true) => {
      const terms = input.split(/\s/).filter(term => term !== "");
      setValue(input);
      if (debounced) {
        debouncedDispatch(terms);
      } else {
        setTerms(terms);
      }
    },
    [debouncedDispatch, setTerms]
  );

  return !user ? null : (
    <>
      <form onSubmit={event => event.preventDefault()}>
        <div className={classes.search}>
          <InputBase
            aria-label="Look up RSS feeds available"
            placeholder="Search..."
            fullWidth
            value={value}
            onChange={event => onChange(event.target.value)}
            inputProps={{ "aria-label": "Search" }}
          />
          <IconButton type="submit" aria-label="Search">
            <SearchIcon />
          </IconButton>
          <IconButton
            aria-label="Clear search"
            disabled={terms.length === 0}
            onClick={() => onChange("", false)}
          >
            <CloseIcon />
          </IconButton>
        </div>
      </form>

      <Tags ids={tags} onChange={setTags} />

      {terms.length === 0 &&
        tags.length === 0 &&
        user.stats &&
        user.stats.subscriptions === 0 && (
          <Empty>
            Use the search field above to find new sources of excitement.
          </Empty>
        )}

      {(terms.length > 0 ||
        tags.length > 0 ||
        (user.stats && user.stats.subscriptions > 0)) && (
        <Results terms={terms} tags={tags} />
      )}
    </>
  );
}
