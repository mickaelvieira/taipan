const CACHE_LONG_LIVED = "long-lived";
const CACHE_SHORT_LIVED = "short-lived";
const activeCaches = [
  CACHE_LONG_LIVED,
  CACHE_SHORT_LIVED
];

/**
 * @param {Cache} cache
 * @param {Request} request
 * @param {Response} response
 *
 * @returns {Response}
 */
function addToCacheIfSuccessful(cache, request, response) {
  if (response.ok) {
    cache.put(request, response.clone());
  }
  return response;
}

/**
 * @param {Request} request
 * @param {Cache} cache
 *
 * @returns {Promise}
 */
async function fetchAndAddToCacheIfSuccessful(request, cache) {
  const response = await fetch(request);
  return addToCacheIfSuccessful(cache, request, response);
}

/**
 * @param {Request} request
 *
 * @returns Promise<Response, Error>
 */
async function cacheFirstWithNetworkFallback(request) {
  const response = await self.caches.match(request);
  return response || fetch(request);
}

/**
 * @param {Request} request
 *
 * @returns Promise<Response, Error>
 */
async function networkWithCacheFallback(request) {
  const cache = await self.caches.open(CACHE_SHORT_LIVED);
  try {
    const response = await fetch(request);
    return addToCacheIfSuccessful(cache, request, response);
  } catch (e) {
    return cache.match(request);
  }
}

/**
 * @param {Request} request
 *
 * @returns {Promise}
 */
async function cacheOnTheFly(request) {
  const cache = await self.caches.open(CACHE_LONG_LIVED);
  const fromCache = await cache.match(request);
  return fromCache || fetchAndAddToCacheIfSuccessful(request, cache);
}

/**
 * Delete outdated caches during the activation
 */
self.addEventListener("activate", event => {
  self.caches.keys().then(names => {
    event.waitUntil(
      Promise.all(
        names
          .filter(name => !activeCaches.includes(name))
          .map(name => self.caches.delete(name))
      )
    );
  });
});

/**
 * @param {String} pathname
 *
 * @returns {Boolean}
 */
function isFont(pathname) {
  return /\.(ttf|ttc|otf|eot|woff|woff2)$/.test(pathname);
}

/**
 * @param {String} pathname
 *
 * @returns {Boolean}
 */
function isScript(pathname) {
  return /\.(css|js)$/.test(pathname);
}

/**
 * @param {String} pathname
 *
 * @returns {Boolean}
 */
function isImage(pathname) {
  return /\.(jpg|jpeg|gif|png|svg|webp)$/.test(pathname);
}

/**
 * @param {String} hostname
 *
 * @returns {Boolean}
 */
function isHostCacheable(hostname) {
  return hostname === "fonts.googleapis.com";
}

/**
 * Handle assets caching on the fly
 */
self.addEventListener("fetch", event => {
  let response;

  const url = new URL(event.request.url);
  const pathname = url.pathname;
  const hostname = url.hostname;
  const request = event.request;

  if (isHostCacheable(hostname) || isImage(pathname) || isScript(pathname) || isFont(pathname)) {
    response = cacheOnTheFly(request);
  } else {
    response = cacheFirstWithNetworkFallback(request);
  }

  event.respondWith(response);
});
