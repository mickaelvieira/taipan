import React from "react";
import { useQuery } from "@apollo/react-hooks";
import { makeStyles } from "@material-ui/core/styles";
import List from "@material-ui/core/List";
import ListItem from "@material-ui/core/ListItem";
import Checkbox from "@material-ui/core/Checkbox";
import InputLabel from "@material-ui/core/InputLabel";
import { Data, query } from "../../../apollo/Query/Tags";
import Loader from "../../Loader";

const useStyles = makeStyles(() => ({
  item: {
    padding: 0
  }
}));

interface Props {
  ids: string[];
  onChange: (ids: string[]) => void;
}

export default React.memo(function Tags({
  ids,
  onChange
}: Props): JSX.Element | null {
  const classes = useStyles();

  const { data, loading, error } = useQuery<Data, {}>(query);

  if (loading) {
    return <Loader />;
  }

  if (error) {
    return <span>{error.message}</span>;
  }

  if (!data) {
    return null;
  }

  const {
    syndication: { tags }
  } = data;

  return (
    <List>
      {tags.results.map(tag => (
        <ListItem key={tag.id} className={classes.item}>
          <Checkbox
            id={`tag-${tag.id}`}
            checked={ids.includes(tag.id)}
            onClick={() => {
              const sub = ids.includes(tag.id)
                ? ids.filter(t => t != tag.id)
                : [tag.id, ...ids];
              onChange(sub);
            }}
          />
          <InputLabel htmlFor={`tag-${tag.id}`}>{tag.label}</InputLabel>
        </ListItem>
      ))}
    </List>
  );
});
