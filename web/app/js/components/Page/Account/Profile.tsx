import React, { useRef, useContext, useReducer, Reducer } from "react";
import { makeStyles } from "@material-ui/core/styles";
import Card from "@material-ui/core/Card";
import CardContent from "@material-ui/core/CardContent";
import CardActions from "@material-ui/core/CardActions";
import Button from "@material-ui/core/Button";
import TextField from "@material-ui/core/TextField";
import Slider from "@material-ui/core/Slider";
import AvatarEditor from "react-avatar-editor";
import FormHelperText from "@material-ui/core/FormHelperText";
import { MessageContext } from "../../context";
import UserProfileMutation from "../../apollo/Mutation/User/Profile";
import { UserContext } from "../../context";
import { User } from "../../../types/users";
import Title from "./Title";

const useStyles = makeStyles(({ palette }) => ({
  actions: {
    padding: 16,
    justifyContent: "flex-end"
  },
  avatar: {
    display: "flex",
    flexDirection: "column",
    width: 350,
    margin: "0 auto",
    alignItems: "center"
  },
  editor: {
    border: "1px solid",
    borderColor: palette.grey[400]
  },
  slider: {
    margin: "6px 0"
  },
  inputFile: {
    display: "none"
  }
}));

interface State {
  firstname: string;
  lastname: string;
  file: File | null;
  scale: number;
}

type Payload = string | number | File | null;

enum ProfileActions {
  FIRSTNAME = "firstname",
  LASTNAME = "lastname",
  AVATAR = "avatar",
  SCALE = "scale"
}

function reducer(
  state: State,
  [type, payload]: [ProfileActions, Payload]
): State {
  switch (type) {
    case ProfileActions.FIRSTNAME:
      return {
        ...state,
        firstname: payload as string
      };
    case ProfileActions.LASTNAME:
      return {
        ...state,
        lastname: payload as string
      };
    case ProfileActions.AVATAR:
      return {
        ...state,
        file: payload as File
      };
    case ProfileActions.SCALE:
      return {
        ...state,
        scale: payload as number
      };
    default:
      throw new Error(`Invalid action type '${type}'`);
  }
}

function getInitialState(user: User | null): State {
  const scale = 1;
  const file = null;
  let firstname = "";
  let lastname = "";

  if (user) {
    firstname = user.firstname;
    lastname = user.lastname;
  }

  return {
    firstname,
    lastname,
    scale,
    file
  };
}

type ProfileReducer = Reducer<State, [ProfileActions, Payload]>;

function getAvatar(file: File | null, editor: AvatarEditor | null): string {
  if (file && editor) {
    const canvas = editor.getImageScaledToCanvas();
    return canvas.toDataURL();
  }
  return "";
}

function getInputFile(file: File | null, user: User | null): File | string {
  if (file) {
    return file;
  }
  if (user && user.image) {
    return user.image.url;
  }
  return "";
}

export default function Profile(): JSX.Element | null {
  const editor = useRef<AvatarEditor | null>();
  const setMessageInfo = useContext(MessageContext);
  const classes = useStyles();
  const user = useContext(UserContext);
  const [state, dispatch] = useReducer<ProfileReducer>(
    reducer,
    getInitialState(user)
  );
  const { firstname, lastname, scale, file } = state;

  console.log(file);
  return (
    <Card>
      <form>
        <Title value="Profile" />
        <CardContent>
          <div className={classes.avatar}>
            <AvatarEditor
              ref={editor}
              image={getInputFile(file, user)}
              width={250}
              height={250}
              border={50}
              color={[255, 255, 255, 0.6]}
              scale={scale}
              className={classes.editor}
            />
            <Slider
              value={scale}
              className={classes.slider}
              disabled={!file}
              min={1}
              onChange={(_, value) =>
                dispatch([
                  ProfileActions.SCALE,
                  typeof value === "object" ? value[0] : value
                ])
              }
            />
            <input
              accept="image/*"
              className={classes.inputFile}
              id="contained-button-file"
              onChange={({ target }) => {
                if (target.files && target.files.length > 0) {
                  dispatch([ProfileActions.AVATAR, target.files[0]]);
                }
              }}
              multiple
              type="file"
            />
            <label htmlFor="contained-button-file">
              <Button variant="contained" component="span" color="primary">
                Add Avatar
              </Button>
            </label>
          </div>
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
            onCompleted={() => {
              console.log("completed");
              dispatch([ProfileActions.AVATAR, null]);
              dispatch([ProfileActions.SCALE, 1]);
              setMessageInfo("You profile has been saved");
            }}
          >
            {(mutate, { loading, error }) => (
              <>
                <Button
                  disabled={loading}
                  variant="contained"
                  color="primary"
                  onClick={() => {
                    if (user) {
                      const image = getAvatar(file, editor.current);
                      mutate({
                        variables: {
                          id: user.id,
                          user: {
                            firstname,
                            lastname,
                            image
                          }
                        }
                      });
                    }
                  }}
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
