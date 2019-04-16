self.importScripts("/caches.js");

/* eslint no-undef: "off" */
const CACHE_ASSETS_VERSION = "assets-" + manifest.assets.version;
const CACHE_HTML_VERSION = "html-" + manifest.html.version;
const CACHE_FONTS_VERSION = "fonts-1";
const CACHE_SHORT_LIVED = "short-lived";
const activeCaches = [
  CACHE_ASSETS_VERSION,
  CACHE_HTML_VERSION,
  CACHE_FONTS_VERSION,
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
async function cacheWithNetworkFallback(request) {
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
async function cacheFontsOnTheFly(request) {
  const cache = await self.caches.open(CACHE_FONTS_VERSION);
  const fromCache = await cache.match(request);
  return fromCache || fetchAndAddToCacheIfSuccessful(request, cache);
}

/**
 * @param {Array} files
 *
 * @returns {Promise}
 */
async function installAssets(files) {
  const cache = await self.caches.open(CACHE_ASSETS_VERSION);
  return await Promise.all(files.map(file => cache.add(file)));
}

/**
 * @param {Array}        files
 *
 * @returns {Promise}
 */
async function installHTML(files) {
  const cache = await self.caches.open(CACHE_HTML_VERSION);
  return await Promise.all(files.map(file => cache.add(file)));
}

/**
 * Cache application's assets during the install
 */
self.addEventListener("install", event => {
  event.waitUntil(
    Promise.all([
      installAssets(manifest.assets.files),
      installHTML(manifest.html.files)
    ])
  );
});

/**
 * Delete outdated caches during the activation
 */
self.addEventListener("activate", async event => {
  const names = await self.caches.keys();
  event.waitUntil(
    Promise.all(
      names
        .filter(name => !activeCaches.includes(name))
        .map(name => self.caches.delete(name))
    )
  );
});

/**
 * Handle assets caching on the fly
 */
self.addEventListener("fetch", event => {
  let response;

  const url = new URL(event.request.url);
  const pathname = url.pathname;
  const request = event.request;
  const method = request.method;

  if (/^\/dist\/fonts/.test(pathname)) {
    response = cacheFontsOnTheFly(request);
  } else if (/^\/bookmark\/\d+/.test(pathname)) {
    response = networkWithCacheFallback(request);
  } else {
    response = cacheWithNetworkFallback(request);
  }

  event.respondWith(response);
});
