import React from "react";
import IconButton from "@material-ui/core/IconButton";
import FavoriteIcon from "@material-ui/icons/Favorite";
import CircularProgress from "@material-ui/core/CircularProgress";
import { Bookmark } from "../../../../types/bookmark";
import FavoriteMutation from "../../../apollo/Mutation/Bookmarks/Favorite";
import { queryFavorites, variables } from "../../../apollo/Query/Feed";

interface Props {
  bookmark: Bookmark;
  onSuccess: (bookmark: Bookmark) => void;
  onError: (message: string) => void;
}

export default React.memo(function Favorite({
  bookmark,
  onSuccess,
  onError
}: Props): JSX.Element {
  return (
    <FavoriteMutation
      onCompleted={data => onSuccess(data.bookmarks.favorite)}
      onError={error => onError(error.message)}
    >
      {(mutate, { loading }) => (
        <>
          <IconButton
            aria-label="Mark as favorite"
            disabled={loading}
            onClick={() =>
              mutate({
                variables: {
                  url: bookmark.url
                },
                refetchQueries: [
                  {
                    query: queryFavorites,
                    variables
                  }
                ]
              })
            }
          >
            {!loading && <FavoriteIcon />}
            {loading && <CircularProgress size={16} />}
          </IconButton>
        </>
      )}
    </FavoriteMutation>
  );
});
