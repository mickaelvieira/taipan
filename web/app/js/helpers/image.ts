import { SM_WIDTH, LG_WIDTH } from "../constant/layout";

type Size = "sm" | "lg";

export function getImageWidth(size: Size): number {
  const windowWidth = window.innerWidth;
  if (size === "sm") {
    return SM_WIDTH < windowWidth ? SM_WIDTH : windowWidth;
  } else if (size === "lg") {
    return LG_WIDTH < windowWidth ? LG_WIDTH : windowWidth;
  } else {
    throw new Error("Incorrect image size");
  }
}
