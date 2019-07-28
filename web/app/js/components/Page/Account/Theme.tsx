import React from "react";
import ButtonBase from "@material-ui/core/ButtonBase";
import { makeStyles } from "@material-ui/core/styles";
import {
  themes,
  defaultTheme,
  getThemeClasses,
  ThemeName
} from "../../ui/themes";
import UserThemeMutation from "../../apollo/Mutation/User/Theme";
import { User } from "../../../types/users";

const useStyles = makeStyles(({ palette, breakpoints }) => ({
  ...getThemeClasses(),
  buttons: {
    display: "flex",
    justifyContent: "center",
    margin: "16px 0"
  },
  button: {
    width: 25,
    height: 25,
    margin: 4,
    [breakpoints.up("md")]: {
      width: 50,
      height: 50
    }
  },
  active: {
    border: `1px solid ${palette.grey[900]}`
  }
}));

interface Props {
  user: User;
}

export default function UserTheme({ user }: Props): JSX.Element {
  const classes = useStyles();
  return (
    <section className={classes.buttons}>
      <UserThemeMutation>
        {mutate => {
          return Object.keys(themes).map((name: string): JSX.Element | null => {
            const { palette } = themes[name as ThemeName];
            if (!palette) {
              return null;
            }
            const { primary } = palette;
            if (!primary) {
              return null;
            }
            const isActive = user.theme
              ? user.theme === name
              : name === defaultTheme;

            return (
              <ButtonBase
                className={`${classes[name]} ${classes.button} ${
                  isActive ? classes.active : ""
                }`}
                key={name}
                onClick={() =>
                  mutate({
                    variables: {
                      id: user.id,
                      theme: name
                    }
                  })
                }
              >
                &nbsp;
              </ButtonBase>
            );
          });
        }}
      </UserThemeMutation>
    </section>
  );
}
