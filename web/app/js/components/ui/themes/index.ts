import { ThemeOptions } from "@material-ui/core/styles/createMuiTheme";
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

export function getThemeColors(): string[] {
  return Object.keys(themes)
    .map((name: string): string => {
      const { palette } = themes[name as ThemeName];
      if (!palette) {
        return "";
      }
      const { primary } = palette;
      if (!primary) {
        return "";
      }
      return primary.main;
    })
    .filter(v => v !== "");
}

interface ThemeClassName {
  border: string;
  backgroundColor: string;
}

type ThemeClasseNames = {
  [k in ThemeName]?: ThemeClassName;
};

export function getThemeClasses(): ThemeClasseNames {
  const classNames: ThemeClasseNames = {};
  Object.keys(themes).forEach(name => {
    const { palette } = themes[name as ThemeName];
    if (palette && palette.primary && palette.primary.main) {
      classNames[name as ThemeName] = {
        border: `1px solid ${palette.primary.main}`,
        backgroundColor: palette.primary.main
      };
    }
  });
  return classNames;
}

export default function getThemeOptions(
  name: ThemeName
): Partial<ThemeOptions> {
  return themes[name];
}
