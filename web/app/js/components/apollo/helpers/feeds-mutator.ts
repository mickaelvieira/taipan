import { ApolloClient } from "apollo-client";
import { mutation as favoriteMutation } from "../Mutation/Bookmarks/Favorite";
import { mutation as unfavoriteMutation } from "../Mutation/Bookmarks/Unfavorite";
import { mutation as bookmarkMutation } from "../Mutation/Bookmarks/Bookmark";
import { mutation as unbookmarkMutation } from "../Mutation/Bookmarks/Unbookmark";
import FeedsUpdater from "./feeds-updater";
import { Bookmark } from "../../../types/bookmark";
import { Document } from "../../../types/document";

export default class FeedsMutator {
  client: ApolloClient<object>;

  cacheUpdater: FeedsUpdater;

  mutation = {
    favorite: favoriteMutation,
    unfavorite: unfavoriteMutation,
    bookmark: bookmarkMutation,
    unbookmark: unbookmarkMutation
  };

  constructor(client: ApolloClient<object>, cacheUpdater: FeedsUpdater) {
    this.client = client;
    this.cacheUpdater = cacheUpdater;
  }

  favorite(bookmark: Bookmark): void {
    this.client.mutate({
      mutation: this.mutation["favorite"],
      variables: {
        url: bookmark.url
      },
      update: (_, result) => {
        const { data } = result;
        this.cacheUpdater.favorite(data.bookmarks.favorite);
      }
    });
  }

  unfavorite(bookmark: Bookmark): void {
    this.client.mutate({
      mutation: this.mutation["unfavorite"],
      variables: {
        url: bookmark.url
      },
      update: (_, result) => {
        const { data } = result;
        this.cacheUpdater.unfavorite(data.bookmarks.unfavorite);
      }
    });
  }

  bookmark(document: Document, isFavorite: boolean): void {
    this.client.mutate({
      mutation: this.mutation["bookmark"],
      variables: {
        url: document.url,
        isFavorite
      },
      update: (_, result) => {
        const { data } = result;
        this.cacheUpdater.bookmark(data.bookmarks.add);
      }
    });
  }

  unbookmark(bookmark: Bookmark): void {
    this.client.mutate({
      mutation: this.mutation["unbookmark"],
      variables: {
        url: bookmark.url
      },
      update: (_, result) => {
        const { data } = result;
        this.cacheUpdater.unbookmark(data.bookmarks.remove);
      }
    });
  }
}
