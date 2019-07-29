import { ThemeOptions } from "@material-ui/core/styles/createMuiTheme";
import { SimplePaletteColorOptions } from "@material-ui/core/styles/createPalette";
import { CSSProperties } from "@material-ui/core/styles/withStyles";
import amber from "./amber";
import blue from "./blue";
import cyan from "./cyan";
import green from "./green";
import indigo from "./indigo";
import orange from "./orange";
import pink from "./pink";
import purple from "./purple";

export type ThemeName =
  | "amber"
  | "blue"
  | "cyan"
  | "green"
  | "indigo"
  | "orange"
  | "pink"
  | "purple";

type Themes = {
  [k in ThemeName]: ThemeOptions;
};

export const themes: Themes = {
  amber,
  blue,
  cyan,
  green,
  indigo,
  orange,
  pink,
  purple
};

export const defaultTheme = "pink";

type ThemeClasseNames = {
  [k in ThemeName]: CSSProperties;
};

export function getThemeClasses(): ThemeClasseNames {
  const classes = Object.keys(themes).reduce(
    (acc: { [k: string]: CSSProperties }, name) => {
      const { palette } = themes[name as ThemeName];
      if (palette && palette.primary) {
        const primary = palette.primary as SimplePaletteColorOptions;
        acc[name] = {
          backgroundColor: primary.main
        };
      }
      return acc;
    },
    {}
  );
  return classes as ThemeClasseNames;
}

export default function getThemeOptions(
  name: ThemeName | null = defaultTheme
): Partial<ThemeOptions> {
  return themes[name === null ? defaultTheme : name];
}
