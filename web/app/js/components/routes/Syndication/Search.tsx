import React, { useCallback, useState, useContext } from "react";
import { debounce } from "lodash";
import { useTheme } from "@material-ui/core/styles";
import useMediaQuery from "@material-ui/core/useMediaQuery";
import FormGroup from "@material-ui/core/FormGroup";
import FormControlLabel from "@material-ui/core/FormControlLabel";
import RadioGroup from "@material-ui/core/RadioGroup";
import Radio from "@material-ui/core/Radio";
import { makeStyles } from "@material-ui/core/styles";
import IconButton from "@material-ui/core/IconButton";
import InputBase from "@material-ui/core/InputBase";
import SearchIcon from "@material-ui/icons/Search";
import CloseIcon from "@material-ui/icons/Close";
import MenuItem from "@material-ui/core/MenuItem";
import Select from "@material-ui/core/Select";
import { UserContext } from "../../context";
import Results from "./Results";
import useSearchReducer, { Action } from "./useSearchReducer";
import Tags from "../../ui/Tags";

const useStyles = makeStyles(({ palette }) => ({
  search: {
    display: "flex",
    margin: 16,
    borderBottom: `1px solid  ${palette.grey[500]}`,
  },
  options: {
    margin: 16,
    alignItems: "center",
  },
  radioLabel: {
    marginRight: 16,
  },
  radioGroup: {
    flexDirection: "row",
  },
}));

interface Props {
  editSource: (url: URL) => void;
}

const sortingValues = [
  { by: "title", dir: "ASC", label: "Title" },
  { by: "updated_at", dir: "DESC", label: "Last Modified" },
];

export default function Search({ editSource }: Props): JSX.Element | null {
  const classes = useStyles();
  const user = useContext(UserContext);
  const theme = useTheme();
  const md = useMediaQuery(theme.breakpoints.up("md"));
  const [state, dispatch] = useSearchReducer();
  const [tags, setTags] = useState<string[]>([]);
  const [sorting, setSorting] = useState<string[]>(0);
  const [value, setValue] = useState("");
  const debouncedDispatch = useCallback(debounce(dispatch, 400), []);
  const onChange = useCallback(
    (input: string, debounced = true) => {
      const terms = input.split(/\s/).filter((term) => term !== "");
      setValue(input);
      if (debounced) {
        debouncedDispatch([Action.TERMS, terms]);
      } else {
        dispatch([Action.TERMS, terms]);
      }
    },
    [debouncedDispatch, dispatch]
  );

  const { terms, hidden, paused } = state;

  return !user ? null : (
    <>
      <form onSubmit={(event) => event.preventDefault()}>
        <div className={classes.search}>
          <InputBase
            aria-label="Look up RSS feeds available"
            placeholder="Search..."
            fullWidth
            value={value}
            onChange={(event) => onChange(event.target.value)}
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
              aria-label="Paused status"
              name="paused"
              className={classes.radioGroup}
              value={paused ? "1" : "0"}
              onChange={(_, value) =>
                dispatch([Action.PAUSED, value === "1" ? true : false])
              }
            >
              <FormControlLabel value="0" control={<Radio />} label="Parsed" />
              <FormControlLabel value="1" control={<Radio />} label="Paused" />
            </RadioGroup>
            <RadioGroup
              aria-label="visibility"
              name="hidden"
              className={classes.radioGroup}
              value={hidden ? "1" : "0"}
              onChange={(_, value) =>
                dispatch([Action.HIDDEN, value === "1" ? true : false])
              }
            >
              <FormControlLabel value="0" control={<Radio />} label="Visible" />
              <FormControlLabel value="1" control={<Radio />} label="Hidden" />
            </RadioGroup>

            <Select
              value={sorting}
              onChange={(event) => setSorting(event.target.value)}
            >
              {sortingValues.map(({ label }, index) => (
                <MenuItem key={index} value={index}>
                  {label}
                </MenuItem>
              ))}
            </Select>
          </FormGroup>
        )}
      </form>

      <Tags ids={tags} onChange={setTags} />

      <Results
        terms={terms}
        tags={tags}
        hidden={hidden}
        paused={paused}
        sort={sortingValues[sorting]}
        editSource={editSource}
      />
    </>
  );
}
