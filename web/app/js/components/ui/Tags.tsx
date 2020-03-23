import React from "react";
import { useQuery } from "@apollo/react-hooks";
import { makeStyles } from "@material-ui/core/styles";
import Chip from "@material-ui/core/Chip";
import { query, Data } from "../apollo/Query/SubscriptionTags";
import Loader from "./Loader";
import { sort } from "../../helpers/tags";

const useStyles = makeStyles(({ palette, spacing }) => ({
  container: {
    margin: `0 ${spacing(1)}px`,
  },
  chip: {
    margin: spacing(1),
  },
  active: {
    color: palette.common.white,
    backgroundColor: palette.primary.main,
  },
}));

interface Props {
  ids: string[];
  onChange: (tags: string[]) => void;
}

export default function Tags({ ids, onChange }: Props): JSX.Element | null {
  const classes = useStyles();
  const { data, loading, error } = useQuery<Data, {}>(query);

  if (error) {
    return <span>{error.message}</span>;
  }

  if (loading) {
    return <Loader />;
  }

  if (!data) {
    return null;
  }

  const {
    subscriptions: { tags },
  } = data;
  const list = sort(tags.results);

  return (
    <div className={classes.container}>
      {list.map((tag) => {
        const active = ids.includes(tag.id);
        return (
          <Chip
            key={tag.id}
            label={tag.label}
            size="small"
            color={active ? "primary" : "default"}
            className={classes.chip}
            onClick={() => {
              const sub = active
                ? ids.filter((t) => t != tag.id)
                : [tag.id, ...ids];
              onChange(sub);
            }}
          />
        );
      })}
    </div>
  );
}
