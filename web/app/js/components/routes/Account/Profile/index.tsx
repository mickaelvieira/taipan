import React, { useRef, useContext } from "react";
import { useMutation } from "@apollo/react-hooks";
import { makeStyles } from "@material-ui/core/styles";
import { useTheme } from "@material-ui/core/styles";
import useMediaQuery from "@material-ui/core/useMediaQuery";
import Card from "../Card";
import CardContent from "@material-ui/core/CardContent";
import CardActions from "../CardActions";
import Button from "@material-ui/core/Button";
import Slider from "@material-ui/core/Slider";
import AvatarEditor from "react-avatar-editor";
import FormHelperText from "@material-ui/core/FormHelperText";
import { MessageContext } from "../../../context";
import {
  mutation,
  Data,
  Variables
} from "../../../apollo/Mutation/User/Profile";
import { User } from "../../../../types/users";
import Title from "../Title";
import Theme from "./Theme";
import Group from "../../../ui/Form/Group";
import { InputBase } from "../../../ui/Form/Input";
import Label from "../../../ui/Form/Label";
import useProfileReducer, { Action } from "../useProfileReducer";

const useStyles = makeStyles(({ palette, breakpoints }) => ({
  avatar: {
    display: "flex",
    flexDirection: "column",
    width: "100%",
    [breakpoints.up("md")]: {
      width: 350
    },
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
    return `${user.image.url}`;
  }
  return "";
}

interface Props {
  user: User;
}

export default function UserProfile({ user }: Props): JSX.Element | null {
  const editor = useRef<AvatarEditor | null>();
  const theme = useTheme();
  const md = useMediaQuery(theme.breakpoints.up("md"));
  const setMessageInfo = useContext(MessageContext);
  const classes = useStyles();
  const [state, dispatch] = useProfileReducer(user);
  const [mutate, { loading, error }] = useMutation<Data, Variables>(mutation, {
    onCompleted: () => {
      dispatch([Action.AVATAR, null]);
      dispatch([Action.SCALE, 1]);
      setMessageInfo({ message: "You profile has been saved" });
    }
  });

  const { firstname, lastname, scale, file } = state;

  return (
    <Card>
      <form onSubmit={event => event.preventDefault()}>
        <Title value="Profile" />
        <CardContent>
          <div className={classes.avatar}>
            <AvatarEditor
              ref={editor}
              image={getInputFile(file, user)}
              width={250}
              height={250}
              border={md ? 50 : 10}
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
                  Action.SCALE,
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
                  dispatch([Action.AVATAR, target.files[0]]);
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
          <Theme user={user} />
          <Group>
            <Label htmlFor="firtname">Firstname</Label>
            <InputBase
              id="firtname"
              value={firstname}
              onChange={event =>
                dispatch([Action.FIRSTNAME, event.target.value])
              }
            />
          </Group>
          <Group>
            <Label htmlFor="lastname">Firstname</Label>
            <InputBase
              id="lastname"
              value={lastname}
              onChange={event =>
                dispatch([Action.LASTNAME, event.target.value])
              }
            />
          </Group>
        </CardContent>
        <CardActions>
          <Button
            type="submit"
            disabled={loading}
            variant="contained"
            color="primary"
            onClick={() => {
              if (user) {
                const image = getAvatar(file, editor.current);
                mutate({
                  variables: {
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
          {error && <FormHelperText error>{error.message}</FormHelperText>}
        </CardActions>
      </form>
    </Card>
  );
}
