import React, { useEffect, useState } from "react";
import { Transition } from "react-transition-group";
import { useTheme } from "@material-ui/core/styles";
import useMediaQuery from "@material-ui/core/useMediaQuery";
import { makeStyles } from "@material-ui/core/styles";
import AppBar from "@material-ui/core/AppBar";
import Toolbar from "@material-ui/core/Toolbar";
import IconButton from "@material-ui/core/IconButton";
import MenuIcon from "@material-ui/icons/Menu";
import SearchIcon from "@material-ui/icons/Search";
import CloseIcon from "@material-ui/icons/Close";
import Typography from "@material-ui/core/Typography";
import { SIDEBAR_WIDTH } from "../../../constant/sidebar";
import useSearch from "../../../hooks/useSearch";
import usePage from "../../../hooks/usePage";
import Search from "./Search";

const animDuration = 200;

const useStyles = makeStyles(({ breakpoints, palette }) => ({
  appBar: {
    marginLeft: SIDEBAR_WIDTH,
    [breakpoints.up("md")]: {
      width: `calc(100% - ${SIDEBAR_WIDTH}px)`
    },
    "& .animation-entering": {
      transform: "translate(0px, -36px)"
    },
    "& .animation-entered": {
      transform: "translate(0px, -36px)"
    },
    "& .animation-exiting": {
      transform: "translate(0px, 0px)"
    },
    "& .animation-exited": {
      transform: "translate(0px, 0px)"
    }
  },
  animated: {
    width: "100%",
    transition: `transform ${animDuration}ms ease-in-out`,
    transform: "translate(0px, 1px)"
  },
  toolbar: {
    display: "flex",
    alignItems: "center",
    justifyContent: "space-between"
  },
  titleBar: {
    display: "flex"
  },
  title: {
    flexGrow: 1,
    textAlign: "center",
    lineHeight: 1.5
  },
  menuButton: {
    paddingTop: 4,
    marginLeft: -12
  },
  searchButton: {
    marginRight: -12,
    color: palette.common.white
  },
  container: {
    display: "flex",
    flexGrow: 1,
    overflow: "hidden",
    height: 30
  }
}));

interface Props {
  toggleDrawer: (status: boolean) => void;
}

export default function Header({ toggleDrawer }: Props): JSX.Element {
  const classes = useStyles();
  const theme = useTheme();
  const page = usePage();
  const md = useMediaQuery(theme.breakpoints.up("md"));
  const title = page.getTitle();
  const [type, terms] = useSearch();
  const [isSearchOpen, setSearchStatus] = useState(terms.length > 0);

  // @TODO add search terms to the document title when looking up documents or articles
  useEffect(() => {
    document.title = title;
  }, [title]);

  return (
    <AppBar position="fixed" className={classes.appBar}>
      <Toolbar className={classes.toolbar}>
        {!md && (
          <>
            <div className={classes.container}>
              <Transition in={isSearchOpen} timeout={animDuration}>
                {state => (
                  <div className={`${classes.animated} animation-${state}`}>
                    <div className={classes.titleBar}>
                      <IconButton
                        className={classes.menuButton}
                        color="inherit"
                        aria-label="Menu"
                        onClick={() => toggleDrawer(true)}
                      >
                        <MenuIcon />
                      </IconButton>
                      <Typography
                        component="h6"
                        variant="h5"
                        className={classes.title}
                      >
                        {page.getSection()}
                      </Typography>
                    </div>
                    <Search type={type} terms={terms} />
                  </div>
                )}
              </Transition>
            </div>
            <IconButton
              className={classes.searchButton}
              onClick={() => setSearchStatus(!isSearchOpen)}
            >
              {isSearchOpen ? <CloseIcon /> : <SearchIcon />}
            </IconButton>
          </>
        )}
        {md && <Search type={type} terms={terms} />}
      </Toolbar>
    </AppBar>
  );
}
