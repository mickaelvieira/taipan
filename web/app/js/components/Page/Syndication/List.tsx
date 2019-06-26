import React from "react";
import List from "@material-ui/core/List";
import { makeStyles } from "@material-ui/core/styles";
import Button from "@material-ui/core/Button";
import SyndicationQuery from "../../apollo/Query/Syndication";
import Item from "./Item";

const useStyles = makeStyles({
  list: {
    width: "100%"
  }
});

export default function SyndicationList(): JSX.Element {
  const classes = useStyles();
  return (
    <List className={classes.list}>
      <SyndicationQuery>
        {({ data, loading, error, fetchMore }) => {
          if (loading) {
            return <span>Loading...</span>;
          }

          if (error) {
            return <span>{error.message}</span>;
          }

          if (!data) {
            return null;
          }

          return (
            <>
              {data.syndication.sources.results.map(source => {
                return <Item key={source.id} source={source} />;
              })}
              <Button
                onClick={() =>
                  fetchMore({
                    variables: {
                      pagination: {
                        limit: 50,
                        offset: data.syndication.sources.results.length
                      }
                    },
                    updateQuery: (prev, { fetchMoreResult: next }) => {
                      if (!next) {
                        return prev;
                      }
                      return {
                        syndication: {
                          ...prev.syndication,
                          sources: {
                            ...prev.syndication.sources,
                            limit: next.syndication.sources.limit,
                            offset: next.syndication.sources.offset,
                            results: [
                              ...prev.syndication.sources.results,
                              ...next.syndication.sources.results
                            ]
                          }
                        }
                      };
                    }
                  })
                }
              >
                Load more
              </Button>
            </>
          );
        }}
      </SyndicationQuery>
    </List>
  );
}
