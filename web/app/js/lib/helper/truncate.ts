export default (str: string, len = 20) =>
  str.length > len ? `${str.substr(0, len)}...` : str;
