import React, { useContext, useReducer, Reducer } from "react";
import { makeStyles } from "@material-ui/core/styles";
import Card from "@material-ui/core/Card";
import CardContent from "@material-ui/core/CardContent";
import CardActions from "@material-ui/core/CardActions";
import Button from "@material-ui/core/Button";
import TextField from "@material-ui/core/TextField";
import FormHelperText from "@material-ui/core/FormHelperText";
import { MessageContext } from "../../context";
import UserProfileMutation from "../../apollo/Mutation/User/Profile";
import { UserContext } from "../../context";
import { User } from "../../../types/users";
import Title from "./Title";

const useStyles = makeStyles(() => ({
  actions: {
    padding: 16,
    justifyContent: "flex-end"
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
  const setMessageInfo = useContext(MessageContext);
  const classes = useStyles();
  const user = useContext(UserContext);
  const [state, dispatch] = useReducer<ProfileReducer>(
    reducer,
    getInitialState(user)
  );

  const { firstname, lastname } = state;
  return (
    <Card>
      <form>
        <Title value="Profile" />
        <CardContent>
          <TextField
            autoFocus
            fullWidth
            margin="normal"
            id="firtname"
            label="Firstname"
            value={firstname}
            onChange={event =>
              dispatch([ProfileActions.FIRSTNAME, event.target.value])
            }
          />
          <TextField
            fullWidth
            margin="normal"
            id="lastname"
            label="Lastname"
            value={lastname}
            onChange={event =>
              dispatch([ProfileActions.LASTNAME, event.target.value])
            }
          />
        </CardContent>
        <CardActions className={classes.actions}>
          <UserProfileMutation
            onCompleted={() => setMessageInfo("You profile has been saved")}
          >
            {(mutate, { loading, error }) => (
              <>
                <Button
                  disabled={loading}
                  variant="contained"
                  color="primary"
                  onClick={() =>
                    mutate({
                      variables: {
                        id: user ? user.id : "",
                        user: {
                          firstname,
                          lastname
                        }
                      }
                    })
                  }
                >
                  Save
                </Button>
                {error && (
                  <FormHelperText error>{error.message}</FormHelperText>
                )}
              </>
            )}
          </UserProfileMutation>
        </CardActions>
      </form>
    </Card>
  );
}
