import React, { PropsWithChildren } from "react";
import { SvgIconProps } from "@material-ui/core/SvgIcon";
import { makeStyles } from "@material-ui/core/styles";
import { useTheme } from "@material-ui/core/styles";
import useMediaQuery from "@material-ui/core/useMediaQuery";
import IconButton from "@material-ui/core/IconButton";
import NavigateBeforeIcon from "@material-ui/icons/NavigateBefore";
import Typography from "@material-ui/core/Typography";
import Modal from "./Modal";
import Panel from "./Panel";

const useStyles = makeStyles(({ palette }) => ({
  header: {
    display: "flex",
    flexDirection: "row",
    justifyContent: "start",
    margin: 0,
    padding: 0,
    color: palette.primary.contrastText,
    backgroundColor: palette.primary.main,
  },
  prev: {
    color: palette.primary.contrastText,
  },
  title: {
    paddingTop: 12,
    paddingBottom: 12,
  },
  container: {
    padding: 16,
    display: "flex",
    flexDirection: "column",
  },
}));

interface Props {
  title: string;
  isOpen: boolean;
  prev: () => void;
  BackButton?: React.ComponentType<SvgIconProps>;
}

export default function Wrapper({
  title,
  isOpen,
  prev,
  children,
  BackButton,
}: PropsWithChildren<Props>): JSX.Element {
  const classes = useStyles();
  const theme = useTheme();
  const md = useMediaQuery(theme.breakpoints.up("md"));
  const Wrapper = md ? Modal : Panel;
  const Button = BackButton ? BackButton : NavigateBeforeIcon;

  return (
    <Wrapper isOpen={isOpen} prev={close}>
      <header className={classes.header}>
        <IconButton onClick={prev} className={classes.prev}>
          <Button />
        </IconButton>
        <Typography component="h5" variant="h6" className={classes.title}>
          {title}
        </Typography>
      </header>
      <section className={classes.container}>{children}</section>
    </Wrapper>
  );
}
