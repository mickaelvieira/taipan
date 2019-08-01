export function calculateTopGap(index: number, heights: number[]): number {
  let n = 0;
  for (let i = 0, l = heights.length - 1; i < l; i++) {
    if (i < index) {
      n = n + heights[i];
    } else {
      break;
    }
  }
  return n;
}

export function calculateBottomGap(index: number, heights: number[]): number {
  let n = 0;
  for (let i = heights.length - 1, l = 0; i > l; i--) {
    if (i > index) {
      n = n + heights[i];
    } else {
      break;
    }
  }
  return n;
}

export function calculateFirstIndex(gap: number, heights: number[]): number {
  let n = 0;
  let f = 0;

  const t = heights.length;

  for (f = 0; f < t; f++) {
    n = n + heights[f];
    if (n < gap) {
      continue;
    } else {
      break;
    }
  }
  return f;
}

export function calculateBoudaries(
  first: number,
  total: number
): [number, number] {
  const page = 12;
  let f = first - page / 2;

  if (f < 0) {
    f = 0;
  }

  let l = f + page - 1;

  if (l > total) {
    l = total;
  }

  return [f, l];
}
