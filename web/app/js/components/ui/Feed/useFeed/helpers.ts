export function calculateHeightUpToFirstIndex(
  index: number,
  heights: number[]
): number {
  let h = 0;
  const l = heights.length - 1;
  for (let i = 0; i < l; i++) {
    if (i < index) {
      h = h + heights[i];
    } else {
      break;
    }
  }
  return h;
}

export function calculateHeightFromIndex(
  index: number,
  heights: number[]
): number {
  let h = 0;
  const l = 0;
  for (let i = heights.length - 1; i > l; i--) {
    if (i > index) {
      h = h + heights[i];
    } else {
      break;
    }
  }
  return h;
}

export function calculateCursor(gap: number, heights: number[]): number {
  const total = heights.length;
  let height = 0;
  let cursor = 0;

  for (cursor = 0; cursor < total; cursor++) {
    height = height + heights[cursor];
    if (height <= gap) {
      continue;
    } else {
      break;
    }
  }
  return cursor;
}

function interval(a: number, b: number): number {
  return b - a + 1;
}

export function calculateInterval(
  cursor: number,
  total: number
): [number, number] {
  const page = 10;
  const halfPage = page / 2;
  const firstIndex = 0;
  const lastIndex = total === 0 ? total : total - 1;

  // if we have less than 15 items, we take all of them, regardless the cursor's position
  if (total <= page + halfPage) {
    return [firstIndex, lastIndex];
  }

  // we pick 11 items, 5 before and after the cursor
  let startIndex = cursor - halfPage;
  let endIndex = cursor + halfPage;

  // if the distance between the start index and the first one
  // is less than a page, we then take the full page
  if (startIndex < page) {
    startIndex = firstIndex;
  }

  // if the distance between the end index and the last one
  // is less than a page, we then take the full page
  if (total - endIndex < page) {
    endIndex = lastIndex;
  }

  // if we go below the first item,
  // that means the cursor is close to the beginning
  // so we pick a 10 items starting from the start
  if (startIndex < firstIndex) {
    startIndex = firstIndex;
    endIndex = startIndex + (page - 1);
  } else if (endIndex > lastIndex) {
    // similarly if we go above the last item,
    // that means the cursor is close to the end
    // so we pick a 10 items starting from the end
    endIndex = lastIndex;
    startIndex = endIndex - (page - 1);
  }

  const ch = interval(startIndex, endIndex);
  if (ch < 10) {
    // console.log(`Interval size should be equal to 10: got ${ch}`);
  }

  return [startIndex, endIndex];
}
