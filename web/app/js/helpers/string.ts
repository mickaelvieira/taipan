const truncate = (input: string, len: number = 250) => {
  if (input.length <= len) {
    return input;
  }

  const words = input
    .split(/\s/)
    .filter(word => word !== "")
    .map(word => word.trim());

  let output = "";
  let next = "";

  do {
    output = next;
    const token = words.shift();

    if (!token) {
      break;
    }

    next = output === "" ? token : `${output} ${token}`;
  } while (output.length < len && next.length <= len);

  return `${output}...`;
};

export { truncate };
