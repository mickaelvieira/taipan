import React from "react";
import { makeStyles } from "@material-ui/core/styles";
import { useMutation } from "@apollo/react-hooks";
import IconButton from "@material-ui/core/IconButton";
import DeleteIcon from "@material-ui/icons/DeleteSharp";
import {
  Data,
  Variables,
  mutation,
} from "../../../apollo/Mutation/Syndication/Tags/Delete";
import { Tag } from "../../../../types/syndication";
import { query } from "../../../apollo/Query/Tags";

const useStyles = makeStyles(() => ({
  button: {
    paddingTop: 0,
    paddingBottom: 0,
  },
}));

interface Props {
  tag: Tag;
}

export default React.memo(function DeleteTagButton({
  tag,
}: Props): JSX.Element {
  const classes = useStyles();
  const [mutate] = useMutation<Data, Variables>(mutation, {
    refetchQueries: [
      {
        query,
      },
    ],
  });

  return (
    <IconButton
      className={classes.button}
      onClick={() => {
        if (window.confirm("Are you sure you want to delete this tag?")) {
          mutate({
            variables: {
              id: tag.id,
            },
          });
        }
      }}
    >
      <DeleteIcon color="primary" />
    </IconButton>
  );
});
