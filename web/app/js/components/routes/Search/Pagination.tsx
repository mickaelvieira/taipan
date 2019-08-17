import React from "react";
import { makeStyles } from "@material-ui/core/styles";
import Terms from "./Terms";
import { SearchType } from "../../../types/search";

const useStyles = makeStyles(({ spacing }) => ({
  container: {
    display: "flex",
    justifyContent: "flex-end",
    padding: spacing(2)
  },
  type: {
    fontWeight: 500
  }
}));

interface Props {
  count: number;
  total: number;
  terms: string[];
  type: SearchType;
  withCount?: boolean;
}

export default function Pagination({
  count,
  total,
  terms,
  type,
  withCount = false
}: Props): JSX.Element | null {
  const classes = useStyles();
  if (count === 0) {
    return null;
  }

  return (
    <div className={classes.container}>
      {!withCount && (
        <span>
          I found {total}{" "}
          <span className={classes.type}>
            {type}
            {count > 1 ? "s" : ""}
          </span>{" "}
          matching the term{terms.length > 1 ? "s" : ""} <Terms terms={terms} />
        </span>
      )}
      {withCount && (
        <span>
          {count} results of {total}
        </span>
      )}
    </div>
  );
}
