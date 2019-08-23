import React, { useCallback, useState, useContext } from "react";
import { debounce } from "lodash";
import { useTheme } from "@material-ui/core/styles";
import useMediaQuery from "@material-ui/core/useMediaQuery";
import FormGroup from "@material-ui/core/FormGroup";
import FormControlLabel from "@material-ui/core/FormControlLabel";
import RadioGroup from "@material-ui/core/RadioGroup";
import Radio from "@material-ui/core/Radio";
import Checkbox from "@material-ui/core/Checkbox";
import { makeStyles } from "@material-ui/core/styles";
import IconButton from "@material-ui/core/IconButton";
import InputBase from "@material-ui/core/InputBase";
import SearchIcon from "@material-ui/icons/Search";
import CloseIcon from "@material-ui/icons/Close";
import { UserContext } from "../../context";
import Results from "./Results";
import useSearchReducer, { Action } from "./useSearchReducer";

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

interface Props {
  editSource: (url: URL) => void;
}

export default function Search({ editSource }: Props): JSX.Element | null {
  const classes = useStyles();
  const user = useContext(UserContext);
  const theme = useTheme();
  const md = useMediaQuery(theme.breakpoints.up("md"));
  const [state, dispatch] = useSearchReducer();
  const [value, setValue] = useState("");
  const debouncedDispatch = useCallback(debounce(dispatch, 400), []);
  const onChange = useCallback(
    (input: string, debounced = true) => {
      const terms = input.split(/\s/).filter(term => term !== "");
      setValue(input);
      if (debounced) {
        debouncedDispatch([Action.TERMS, terms]);
      } else {
        dispatch([Action.TERMS, terms]);
      }
    },
    [debouncedDispatch, dispatch]
  );

  const { terms, showDeleted, pausedOnly } = state;

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
        {md && (
          <FormGroup row className={classes.options}>
            <RadioGroup
              aria-label="deleted"
              name="deleted"
              className={classes.radioGroup}
              value={showDeleted ? "1" : "0"}
              onChange={(_, value) =>
                dispatch([Action.DELETED, value === "1" ? true : false])
              }
            >
              <FormControlLabel value="0" control={<Radio />} label="Enabled" />
              <FormControlLabel
                value="1"
                control={<Radio />}
                label="Disabled"
              />
            </RadioGroup>
            <FormControlLabel
              control={
                <Checkbox
                  checked={pausedOnly}
                  onChange={() => dispatch([Action.PAUSED, !pausedOnly])}
                  value="Paused"
                />
              }
              labelPlacement="start"
              label="Paused only"
            />
          </FormGroup>
        )}
      </form>
      <Results
        terms={terms}
        showDeleted={showDeleted}
        pausedOnly={pausedOnly}
        editSource={editSource}
      />
    </>
  );
}
