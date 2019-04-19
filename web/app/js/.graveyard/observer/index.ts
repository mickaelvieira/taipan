import { isEqual } from "lodash";

/**
 * Notify observers who are watching the store's changes
 */
export class Broadcaster {
  /**
   * @param {Object} store
   */
  constructor(store) {
    this.store = store;
    this.store.subscribe(() => this.publish());
    this.observers = new Map();
  }

  /**
   * @param {Object} observer
   *
   * @returns {Broadcaster}
   */
  subscribe(observer) {
    const initState = observer.storeMapper(this.store.getState());
    this.observers.set(observer, initState);
    if (typeof observer === "function") {
      observer(initState);
    } else {
      observer.onChange(initState);
    }

    return this;
  }

  /**
   * @returns {Broadcaster}
   */
  publish() {
    const state = this.store.getState();
    this.observers.forEach((prevState, observer, map) => {
      const nextState = observer.storeMapper(state);
      if (!isEqual(prevState, nextState)) {
        this.observers.set(observer, nextState);
        if (typeof observer === "function") {
          observer(nextState);
        } else {
          observer.onChange(nextState);
        }
      }
    });

    return this;
  }

  /**
   * @param {Object} observer
   *
   * @returns {Broadcaster}
   */
  unsubscribe(observer) {
    if (this.observers.has(observer)) {
      this.observers.delete(observer);
    }
    return this;
  }
}
