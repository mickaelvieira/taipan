const getDocumentElement = (): HTMLElement =>
  "documentElement" in window.document
    ? document.documentElement
    : document.body;

const getWindowDimensions = (): { width: number; height: number } => {
  const width = window.innerWidth;
  const height = window.innerHeight;

  return { width, height };
};

const getDocumentDimensions = (): { width: number; height: number } => {
  const doc = getDocumentElement();
  const width = Math.max(doc.scrollWidth, doc.offsetWidth, doc.clientWidth);
  const height = Math.max(doc.scrollHeight, doc.offsetHeight, doc.clientHeight);

  return { width, height };
};

const hasScrollBars = (): { bottom: boolean; right: boolean } => {
  const doc = getDocumentElement();
  return {
    bottom: doc.scrollWidth > doc.clientWidth,
    right: doc.scrollHeight > doc.clientHeight,
  };
};

const getScrollPosition = (): { x: number; y: number } => {
  const x = window.pageXOffset;
  const y = window.pageYOffset;

  return { x, y };
};

const isInViewport = (element: HTMLElement | null): boolean => {
  if (!element) {
    return false;
  }

  const bounding = element.getBoundingClientRect();
  const bottom = window.innerHeight || document.documentElement.clientHeight;
  const right = window.innerWidth || document.documentElement.clientWidth;

  return (
    bounding.top >= 0 &&
    bounding.left >= 0 &&
    bounding.bottom <= bottom &&
    bounding.right <= right
  );
};

const willBeSoonInViewport = (element: HTMLElement | null): boolean => {
  if (!element) {
    return false;
  }

  const bounding = element.getBoundingClientRect();
  const bottom = window.innerHeight || document.documentElement.clientHeight;

  return bounding.top >= bottom + 200;
};

const hasReachedTheBottom = (gap = 400): boolean => {
  const win = getWindowDimensions();
  const doc = getDocumentDimensions();
  const scroll = getScrollPosition();

  const posY = win.height + scroll.y;
  const boundY = doc.height - gap;

  return posY >= boundY;
};

export {
  isInViewport,
  willBeSoonInViewport,
  hasReachedTheBottom,
  getWindowDimensions,
  getDocumentDimensions,
  getScrollPosition,
  hasScrollBars,
};
