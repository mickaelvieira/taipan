import React from "react";
import { useQuery } from "@apollo/react-hooks";
import { makeStyles } from "@material-ui/core/styles";
import Table from "@material-ui/core/Table";
import TableBody from "@material-ui/core/TableBody";
import TableCell from "@material-ui/core/TableCell";
import TableHead from "@material-ui/core/TableHead";
import TableRow from "@material-ui/core/TableRow";
import red from "@material-ui/core/colors/red";
import {
  query as querySource,
  Data as SourceData,
  Variables as SourceVariables
} from "../../../apollo/Query/Source";
import {
  Data as TagsData,
  query as queryTags
} from "../../../apollo/Query/Tags";
import Loader from "../../Loader";
import { TagButton } from "../../Syndication/Button";
import FormTag from "./FormTag";
import EditTag from "./EditTag";
import DeleteTagButton from "./DeleteTag";

const useStyles = makeStyles(({ breakpoints, typography }) => ({
  container: {
    width: "100%",
    overflowX: "hidden",
    overflowY: "auto",
    height: 380,
    maxHeight: 380
  },
  table: {
    [breakpoints.up("md")]: {
      minWidth: 650
    }
  },
  error: {
    color: red[500],
    fontWeight: typography.fontWeightBold
  },
  checkbox: {
    paddingTop: 0,
    paddingBottom: 0
  }
}));

interface Props {
  url: URL;
}

export default React.memo(function Tags({ url }: Props): JSX.Element | null {
  const classes = useStyles();
  const {
    data: sourceData,
    loading: isLoadingSource,
    error: errorSource
  } = useQuery<SourceData, SourceVariables>(querySource, {
    variables: { url }
  });
  const { data: tagsData, loading: isLoadingTags, error: errorTags } = useQuery<
    TagsData,
    {}
  >(queryTags);

  if (isLoadingSource || isLoadingTags) {
    return <Loader />;
  }

  if (errorSource) {
    return <span>{errorSource.message}</span>;
  }
  if (errorTags) {
    return <span>{errorTags.message}</span>;
  }

  if (!sourceData || !tagsData) {
    return null;
  }

  const {
    syndication: { source }
  } = sourceData;

  if (!source) {
    return null;
  }

  const ids = source.tags ? source.tags.map(({ id }) => id) : [];

  const {
    syndication: { tags }
  } = tagsData;

  return (
    <div className={classes.container}>
      <FormTag />
      <Table className={classes.table} size="small">
        <TableHead>
          <TableRow>
            <TableCell align="center">Label</TableCell>
            <TableCell align="center">&nbsp;</TableCell>
          </TableRow>
        </TableHead>
        <TableBody>
          {tags.results.map(tag => (
            <TableRow key={tag.id}>
              <TableCell>
                <EditTag tag={tag} />
              </TableCell>
              <TableCell align="right">
                <TagButton
                  source={source}
                  ids={ids}
                  tag={tag}
                  className={classes.checkbox}
                />
                <DeleteTagButton tag={tag} />
              </TableCell>
            </TableRow>
          ))}
        </TableBody>
      </Table>
    </div>
  );
});
