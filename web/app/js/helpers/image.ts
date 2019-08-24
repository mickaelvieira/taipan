import { FEED_SM_WIDTH, FEED_LG_WIDTH } from "../constant/layout";

type Size = "sm" | "lg";

export function getImageWidth(size: Size): number {
  const windowWidth = window.innerWidth;
  if (size === "sm") {
    return FEED_SM_WIDTH < windowWidth ? FEED_SM_WIDTH : windowWidth;
  } else if (size === "lg") {
    return FEED_LG_WIDTH < windowWidth ? FEED_LG_WIDTH : windowWidth;
  } else {
    throw new Error("Incorrect image size");
  }
}
