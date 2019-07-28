import React, {
  useReducer,
  useCallback,
  useState,
  useContext,
  Reducer
} from "react";
import { debounce } from "lodash";
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
import Table from "./Table";
import EditSource from "./EditSource";
import { isAdmin } from "../../../helpers/users";

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

interface State {
  terms: string[];
  showDeleted: boolean;
  pausedOnly: boolean;
}

type Payload = string[] | boolean;

enum SearchActions {
  TERMS = "terms",
  DELETED = "deleted",
  PAUSED = "paused"
}

function reducer(
  state: State,
  [type, payload]: [SearchActions, Payload]
): State {
  switch (type) {
    case SearchActions.TERMS:
      return {
        ...state,
        terms: payload as string[]
      };
    case SearchActions.DELETED:
      return {
        ...state,
        showDeleted: payload as boolean
      };
    case SearchActions.PAUSED:
      return {
        ...state,
        pausedOnly: payload as boolean
      };
    default:
      throw new Error(`Invalid action type '${type}'`);
  }
}

type SearchReducer = Reducer<State, [SearchActions, Payload]>;

export default function Search(): JSX.Element {
  const classes = useStyles();
  const user = useContext(UserContext);
  const canEdit = isAdmin(user);
  const [state, dispatch] = useReducer<SearchReducer>(reducer, {
    terms: [],
    showDeleted: false,
    pausedOnly: false
  });
  const [value, setValue] = useState("");
  const [editUrl, setEditURL] = useState("");
  const debouncedDispatch = useCallback(debounce(dispatch, 400), []);
  const onChange = useCallback(
    (terms: string[], debounced = true) => {
      setValue(terms.join(" "));
      if (debounced) {
        debouncedDispatch([SearchActions.TERMS, terms]);
      } else {
        dispatch([SearchActions.TERMS, terms]);
      }
    },
    [debouncedDispatch]
  );

  const { terms, showDeleted, pausedOnly } = state;

  return (
    <>
      <form onSubmit={event => event.preventDefault()}>
        <div className={classes.search}>
          <InputBase
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
        {canEdit && (
          <FormGroup row className={classes.options}>
            <RadioGroup
              aria-label="deleted"
              name="deleted"
              className={classes.radioGroup}
              value={showDeleted ? "1" : "0"}
              onChange={(_, value) =>
                dispatch([SearchActions.DELETED, value === "1" ? true : false])
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
                  onChange={() => dispatch([SearchActions.PAUSED, !pausedOnly])}
                  value="Paused"
                />
              }
              labelPlacement="start"
              label="Paused only"
            />
          </FormGroup>
        )}
      </form>
      <Table
        terms={terms}
        showDeleted={showDeleted}
        pausedOnly={pausedOnly}
        canEdit={canEdit}
        editSource={setEditURL}
      />
      {canEdit && (
        <EditSource
          url={editUrl}
          isOpen={editUrl !== ""}
          close={() => setEditURL("")}
        />
      )}
    </>
  );
}
