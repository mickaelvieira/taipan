function getRequest(url) {
  const mode = "cors";
  const headers = new Headers();
  const credentials = "include";

  headers.append("Accept", "application/vnd.collection+json");
  headers.append("X-Requested-With", "XMLHttpRequest");

  return new Request(url, {
    method: "GET",
    headers,
    credentials,
    mode
  });
}

const inFlight = {};

onmessage = function(event) {
  const urls = !Array.isArray(event.data) ? [event.data] : event.data;
  const sent = Object.keys(inFlight);
  const promises = urls
    .filter(url => !sent.includes(url))
    .map(url => {
      inFlight[url] = true;
      return url;
    })
    .map(url => fetch(getRequest(url)));

  Promise.all(promises)
    .then(response => {
      delete inFlight[response.url];
      return response;
    })
    .then(responses => Promise.all(responses.map(response => response.json())))
    .then(results => results.forEach(result => postMessage(result)))
    .catch(error => {
      throw error;
    });
};
