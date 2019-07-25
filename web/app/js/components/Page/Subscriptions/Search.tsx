import React, { useCallback, useState } from "react";
import { debounce } from "lodash";
import { makeStyles } from "@material-ui/core/styles";
import IconButton from "@material-ui/core/IconButton";
import InputBase from "@material-ui/core/InputBase";
import SearchIcon from "@material-ui/icons/Search";
import CloseIcon from "@material-ui/icons/Close";
import Table from "./Table";

const useStyles = makeStyles(({ palette }) => ({
  search: {
    display: "flex",
    margin: 16,
    borderBottom: `1px solid  ${palette.grey[500]}`
  }
}));

export default function Search(): JSX.Element {
  const classes = useStyles();
  const [value, setValue] = useState<string>("");
  const [terms, setTerms] = useState<string[]>([]);
  const setTermsDebounced = useCallback(debounce(setTerms, 400), []);
  const onChange = useCallback(
    (terms: string[], debounced = true) => {
      setValue(terms.join(" "));
      if (debounced) {
        setTermsDebounced(terms);
      } else {
        setTerms(terms);
      }
    },
    [setTermsDebounced]
  );

  return (
    <>
      <form onSubmit={event => event.preventDefault()}>
        <div className={classes.search}>
          <InputBase
            autoFocus
            placeholder="Search..."
            fullWidth
            value={value}
            onChange={event => onChange(event.target.value.split(/\s/))}
            inputProps={{ "aria-label": "Search" }}
          />
          <IconButton type="submit">
            <SearchIcon />
          </IconButton>
          <IconButton onClick={() => onChange([], false)}>
            <CloseIcon />
          </IconButton>
        </div>
      </form>
      <Table terms={terms} />
    </>
  );
}
