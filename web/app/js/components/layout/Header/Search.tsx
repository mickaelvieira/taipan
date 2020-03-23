import React, { useState, useCallback, useEffect, useRef } from "react";
import { useTheme } from "@material-ui/core/styles";
import useMediaQuery from "@material-ui/core/useMediaQuery";
import { debounce } from "lodash";
import { withRouter } from "react-router";
import { makeStyles } from "@material-ui/core/styles";
import InputBase from "@material-ui/core/InputBase";
import ButtonBase from "@material-ui/core/ButtonBase";
import SearchIcon from "@material-ui/icons/Search";
import CloseIcon from "@material-ui/icons/Close";
import { RoutesProps } from "../../../types/routes";

const useStyles = makeStyles(({ palette }) => ({
  search: {
    flexGrow: 1,
    display: "flex",
    alignItems: "center",
  },
  searchLabel: {
    flexGrow: 1,
    margin: "0 24px",
  },
  searchButton: {
    display: "flex",
    alignItems: "center",
    justifyContent: "center",
  },
  clearButton: {
    display: "flex",
    alignItems: "center",
    justifyContent: "center",
    visibility: "hidden",
    "&.active": {
      visibility: "visible",
    },
  },
  inputRoot: {
    width: "100%",
  },
  inputInput: {
    color: palette.common.white,
  },
}));

interface Props extends RoutesProps {
  type: string;
  terms: string[];
  className?: string;
}

export default withRouter(function Search({
  type,
  terms,
  history,
  className,
}: Props): JSX.Element {
  const classes = useStyles();
  const inputRef = useRef<HTMLInputElement>(null);
  const theme = useTheme();
  const md = useMediaQuery(theme.breakpoints.up("md"));
  const [search, setSearch] = useState<string[]>(terms);
  const [value, setValue] = useState<string>(search.join(" "));
  const redirect = useCallback(debounce(history.push, 1000), []);

  useEffect(() => {
    if (search.length > 0) {
      redirect(
        `/search?type=${type}&terms=${encodeURIComponent(search.join(" "))}`
      );
    }
  }, [redirect, search, type]);

  return (
    <form
      className={`${classes.search} ${className ? className : ""}`}
      action={`/search`}
    >
      <ButtonBase
        type="submit"
        className={classes.searchButton}
        aria-label="Search"
      >
        <SearchIcon />
      </ButtonBase>
      <input type="hidden" name="type" value={type} />
      <label htmlFor="search-field" className={classes.searchLabel}>
        <InputBase
          autoFocus={terms.length > 0}
          id="search-field"
          placeholder={`Search ${
            type === "document" ? "the news" : "your bookmarks"
          }`}
          name="terms"
          autoComplete="off"
          value={value}
          inputRef={inputRef}
          aria-label="Look up your bookmarks"
          onChange={(event) => {
            const value = event.target.value;
            setValue(value);
            setSearch(value.split(/\s/).filter((word) => word !== ""));
          }}
          classes={{
            root: classes.inputRoot,
            input: classes.inputInput,
          }}
        />
      </label>
      {md && (
        <ButtonBase
          aria-label="Clear search"
          className={`${classes.clearButton} ${
            search.length > 0 ? "active" : ""
          }`}
          onClick={() => {
            setValue("");
            setSearch([]);
            if (inputRef.current) {
              inputRef.current.focus();
            }
          }}
        >
          <CloseIcon />
        </ButtonBase>
      )}
    </form>
  );
});
