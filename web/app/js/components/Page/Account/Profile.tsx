import React, { useContext, useReducer, Reducer } from "react";
import { makeStyles } from "@material-ui/core/styles";
import Paper from "@material-ui/core/Paper";
import TextField from "@material-ui/core/TextField";
import { UserContext } from "../../context";
import { User } from "../../../types/users";
import Title from "./Title";

const useStyles = makeStyles(() => ({
  paper: {
    padding: 24
  }
}));

interface State {
  firstname: string;
  lastname: string;
}

type Payload = string;

enum ProfileActions {
  FIRSTNAME = "firstname",
  LASTNAME = "lastname"
}

function reducer(
  state: State,
  [type, payload]: [ProfileActions, Payload]
): State {
  switch (type) {
    case ProfileActions.FIRSTNAME:
      return {
        ...state,
        firstname: payload
      };
    case ProfileActions.LASTNAME:
      return {
        ...state,
        lastname: payload
      };
    default:
      throw new Error(`Invalid action type '${type}'`);
  }
}

function getInitialState(user: User | null): State {
  if (!user) {
    return {
      firstname: "",
      lastname: ""
    };
  }

  const { firstname, lastname } = user;
  return {
    firstname,
    lastname
  };
}

type ProfileReducer = Reducer<State, [ProfileActions, Payload]>;

export default function Profile(): JSX.Element | null {
  const classes = useStyles();
  const user = useContext(UserContext);
  const [state, dispatch] = useReducer<ProfileReducer>(
    reducer,
    getInitialState(user)
  );

  const { firstname, lastname } = state;
  return (
    <Paper className={classes.paper}>
      <Title value="Profile" />
      <form>
        <TextField
          autoFocus
          margin="normal"
          id="firtname"
          label="Firstname"
          value={firstname}
          error={false}
          autoComplete="off"
          autoCapitalize="off"
          autoCorrect="off"
          helperText={""}
          onChange={event =>
            dispatch([ProfileActions.FIRSTNAME, event.target.value])
          }
          fullWidth
        />
        <TextField
          margin="normal"
          id="lastname"
          label="Lastname"
          value={lastname}
          error={false}
          autoComplete="off"
          autoCapitalize="off"
          autoCorrect="off"
          helperText={""}
          onChange={event =>
            dispatch([ProfileActions.LASTNAME, event.target.value])
          }
          fullWidth
        />
        {/* <FormGroup>
            <FormControlLabel
              control={
                <Checkbox
                  checked={withFeeds}
                  onChange={() => setWithFeeds(!withFeeds)}
                  value="1"
                  color="primary"
                />
              }
              label="Parse RSS feeds"
            />
          </FormGroup> */}
      </form>
    </Paper>
  );
}
