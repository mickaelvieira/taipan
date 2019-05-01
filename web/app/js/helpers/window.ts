const getDocumentElement = () =>
  "documentElement" in window.document
    ? window.document.documentElement
    : window.document.body;

const getWindowDimensions = () => {
  const width = window.innerWidth;
  const height = window.innerHeight;

  return { width, height };
};

const getDocumentDimensions = () => {
  const doc = getDocumentElement();
  const width = Math.max(doc.scrollWidth, doc.offsetWidth, doc.clientWidth);
  const height = Math.max(doc.scrollHeight, doc.offsetHeight, doc.clientHeight);

  return { width, height };
};

const hasScrollBars = () => {
  const doc = getDocumentElement();
  return {
    bottom: doc.scrollWidth > doc.clientWidth,
    right: doc.scrollHeight > doc.clientHeight
  };
};

const getScrollPosition = () => {
  const x = window.pageXOffset;
  const y = window.pageYOffset;

  return { x, y };
};

const hasReachedTheBottom = (gap = 400) => {
  const win = getWindowDimensions();
  const doc = getDocumentDimensions();
  const scroll = getScrollPosition();

  const posY = win.height + scroll.y;
  const boundY = doc.height - gap;

  return posY >= boundY;
};

export {
  hasReachedTheBottom,
  getWindowDimensions,
  getDocumentDimensions,
  getScrollPosition,
  hasScrollBars
};
