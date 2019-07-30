import React, { useState, useCallback, useEffect } from "react";
import { debounce } from "lodash";
import { withRouter } from "react-router";
import { makeStyles } from "@material-ui/core/styles";
import InputBase from "@material-ui/core/InputBase";
import ButtonBase from "@material-ui/core/ButtonBase";
import SearchIcon from "@material-ui/icons/Search";
import CloseIcon from "@material-ui/icons/Close";
import { RoutesProps } from "../../../types/routes";
import { getSearchTerms, getSearchType } from "../../../helpers/search";

const useStyles = makeStyles(({ palette }) => ({
  search: {
    flexGrow: 1,
    display: "flex",
    alignItems: "center"
  },
  searchLabel: {
    flexGrow: 1,
    margin: "0 24px"
  },
  searchIcon: {
    display: "flex",
    alignItems: "center",
    justifyContent: "center"
  },
  inputRoot: {
    width: "100%"
  },
  inputInput: {
    color: palette.common.white
  }
}));

export default withRouter(function Search({
  history
}: RoutesProps): JSX.Element {
  const classes = useStyles();
  const url = new URL(`${document.location}`);
  const terms = getSearchTerms(url.searchParams.get("terms"));
  const [search, setSearch] = useState<string[]>(terms);
  const [value, setValue] = useState<string>(search.join(" "));
  const redirect = useCallback(debounce(history.push, 1000), []);
  const type = getSearchType(url.pathname, url.searchParams.get("type"));

  useEffect(() => {
    if (search.length > 0) {
      redirect(
        `/search?type=${type}&terms=${encodeURIComponent(search.join(" "))}`
      );
    }
  }, [redirect, search, type]);

  return (
    <form className={classes.search} action={`/search`}>
      <ButtonBase type="submit" className={classes.searchIcon}>
        <SearchIcon />
      </ButtonBase>
      <label htmlFor="search-field" className={classes.searchLabel}>
        <input type="hidden" name="type" value={type} />
        <InputBase
          id="search-field"
          placeholder="Search..."
          name="terms"
          autoComplete="off"
          value={value}
          onChange={event => {
            const value = event.target.value;
            setValue(value);
            setSearch(value.split(/\s/).filter(word => word !== ""));
          }}
          classes={{
            root: classes.inputRoot,
            input: classes.inputInput
          }}
        />
      </label>
      <ButtonBase
        className={classes.searchIcon}
        onClick={() => {
          setValue("");
          setSearch([]);
        }}
      >
        <CloseIcon />
      </ButtonBase>
    </form>
  );
});
