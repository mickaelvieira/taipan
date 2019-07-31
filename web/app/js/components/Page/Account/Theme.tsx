import React from "react";
import ButtonBase from "@material-ui/core/ButtonBase";
import TickedIcom from "@material-ui/icons/Done";
import { makeStyles } from "@material-ui/core/styles";
import {
  themes,
  defaultTheme,
  getThemeClasses,
  ThemeName
} from "../../ui/themes";
import Empty from "../../ui/Empty";
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
    color: palette.common.white,
    [breakpoints.up("md")]: {
      width: 50,
      height: 50
    }
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
          return Object.keys(themes).map((name): JSX.Element | null => {
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
                className={`${classes[name as ThemeName]} ${classes.button}`}
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
                {isActive ? <TickedIcom /> : <Empty />}
              </ButtonBase>
            );
          });
        }}
      </UserThemeMutation>
    </section>
  );
}
