import React from "react";
import { useQuery } from "@apollo/react-hooks";
import { makeStyles } from "@material-ui/core/styles";
import CardMedia from "@material-ui/core/CardMedia";
import NoImage from "../../../NoImage";
import { Document } from "../../../../../types/document";
import { Bookmark } from "../../../../../types/bookmark";
import Title from "../../Item/Title";
import { getImageWidth } from "../../../../../helpers/image";
import { query, Data, Variables } from "../../../../apollo/Query/Bookmark";
import Loader from "../../../Loader";
import NewBookmark from "./FormBookmarkActions/NewBookmark";
import ExistingBookmark from "./FormBookmarkActions/ExistingBookmark";

const useStyles = makeStyles(({ breakpoints }) => ({
  form: {
    display: "flex",
    flexDirection: "column"
  },
  media: {
    marginTop: 16,
    marginBottom: 16,
    backgroundSize: "cover",
    minHeight: `calc(${getImageWidth("sm")}px * 9 / 16)`,
    [breakpoints.up("md")]: {
      minHeight: `calc(${getImageWidth("sm")}px * 9 / 16)`
    }
  },
  nomedia: {
    marginTop: 16,
    marginBottom: 16
  },
  item: {
    padding: 0
  },
  link: {
    paddingTop: 10,
    paddingLeft: 9,
    paddingBottom: 9
  },
  actions: {
    display: "flex",
    flexDirection: "row",
    justifyContent: "flex-end",
    alignItems: "center"
  },
  addto: {
    padding: "8px 6px"
  }
}));

interface Props {
  document: Document;
  onFinish: (bookmark: Bookmark) => void;
}

export default function FormBookmark({
  document,
  onFinish
}: Props): JSX.Element {
  const classes = useStyles();
  const { data, loading, error } = useQuery<Data, Variables>(query, {
    variables: { url: document.url }
  });

  return (
    <form className={classes.form}>
      <Title item={document} />
      {document.image && (
        <CardMedia
          className={classes.media}
          image={document.image.url}
          title={document.title}
        />
      )}
      {!document.image && <NoImage className={classes.nomedia} />}
      {loading && <Loader />}
      {error && <NewBookmark document={document} onFinish={onFinish} />}
      {data && data.bookmarks.bookmark && (
        <ExistingBookmark
          bookmark={data.bookmarks.bookmark}
          onFinish={onFinish}
        />
      )}
    </form>
  );
}
