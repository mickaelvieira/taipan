import combine from "../../lib/helper/combine";

const templateCaches = {};

/**
 * Custom Error class for missing templates
 */
class MissingTemplate extends Error {
  /**
   * @param {String} name
   */
  constructor(name) {
    super(`No template matches the name '${name}'`);
  }
}

/**
 * Helper class carrying the templates available
 */
export class Templates {
  /**
   * @param {Object} templates
   */
  constructor(templates = {}) {
    this.templates = templates;
  }

  /**
   * Returns the template's compiled version
   * @param {String} name
   * @returns {Function}
   */
  compile(name) {
    if (!this.has(name)) {
      throw new MissingTemplate(name);
    }

    return _.template(this.templates[name]);
  }

  /**
   * Returns the template's rendered HTML version
   * @param {String} name
   * @param {Object} data
   * @returns {String}
   */
  render(name, data = {}) {
    const compiled = this.compile(name);

    return compiled(data);
  }

  /**
   * @param {String} name
   * @returns {Boolean}
   */
  has(name) {
    return this.templates.hasOwnProperty(name);
  }

  /**
   * @param {String} name
   * @returns {String}
   */
  get(name) {
    if (!this.has(name)) {
      throw new MissingTemplate(name);
    }

    return this.templates[name];
  }

  /**
   * @returns {Array}
   */
  get names() {
    return Object.keys(this.templates);
  }

  /**
   * @returns {Object}
   */
  get all() {
    return this.templates;
  }
}

/**
 * Normalize the URL as we can get URLs such as `/test/../common/templates/test.html`
 * and return the path
 *
 * @param {String} url
 *
 * @returns {String}
 */
function getURLPathName(url) {
  return new URL(url).pathname.replace(/^\//, "");
}

/**
 * @param {String} sourceScript
 *
 * @returns {String}
 */
function getBaseUrlFromCurrentScript(sourceScript) {
  if (!sourceScript) {
    throw "Cannot determine baseUrl";
  }

  const srcParts = sourceScript.split("/");
  srcParts.pop();

  return srcParts.length ? srcParts.join("/") + "/" : "/";
}

/**
 * Appends the file's hash to the URL if it exists
 *
 * @param {String} url
 *
 * @returns {String}
 */
function getUrlWithHash(url) {
  const pathname = getURLPathName(url);
  return templateCaches.hasOwnProperty(pathname)
    ? (url += "?" + templateCaches[pathname])
    : url;
}

/**
 * Extract the name from the URL
 * For instance:
 * - templates/test.html will return `test`
 * - templates/test/test.html will return `test_test`
 *
 * @param {String} url
 *
 * @returns {String}
 */
function getTemplateNameFromUrl(url) {
  const path = url.replace(/^(.*)templates\//, "");
  const parts = path.split("/");
  const name = parts.pop().replace(/\.html$/, "");

  parts.push(name);

  return parts.join("_");
}

/**
 * Fetch all templates
 *
 * @param {String} baseUrl
 * @param {Array} templates
 *
 * @returns {Array}
 */
function fetchAll(baseUrl, templates) {
  return templates.map(template =>
    window.fetch(getUrlWithHash(baseUrl + template), {
      credentials: "same-origin"
    })
  );
}

/**
 * As a 404 does not constitute a network error
 * We make sure all fetches were successful,
 * otherwise we trigger a exception to reject the promise
 *
 * @param {Array} responses
 *
 * @returns {Promise}
 */
function checkResponses(responses) {
  return new Promise(function(resolve) {
    responses.forEach(function(response) {
      if (!response.ok) {
        throw new Error(
          `Request failed: ${response.status} ${response.statusText} ${
            response.url
          }`
        );
      }
    });

    resolve(responses);
  });
}

/**
 * Extract the file names without the extension from the URL
 * to bind them later on to the returned templates
 *
 * @param {Array} responses
 *
 * @returns {Promise<Templates>}
 */
function appendTemplatesNames(responses) {
  return Promise.resolve().then(function() {
    const names = responses.map(response =>
      getTemplateNameFromUrl(response.url)
    );

    return {
      responses,
      names
    };
  });
}

/**
 * Resolve all text responses and bind the html contents to their corresponding name
 *
 * @param {Object} responses{ responses, names }
 *
 * @returns {Promise<Templates>}
 */
function bindTemplatesNamesToTextResponses(responses) {
  // prettier-ignore
  return Promise
    .all(responses.responses.map(response =>
      response.text())
    )
    .then(contents =>
      new Templates(combine(responses.names, contents))
    );
}

/**
 * It will pass to the `then` method, called against this Promise,
 * the Templates object that can be used to handle the HTML templates.
 *
 * For instance:
 * - templates/foo.html
 * - templates/foo/bar.html
 * - templates/foo/bar_bar.html
 *
 * To get the compile version:
 * - Template.compile('foo')
 * - Template.compile('foo_bar')
 * - Template.compile('foo_bar_bar')
 *
 * To get the rendered version:
 * - Template.render('foo', { param: value })
 * - Template.render('foo_bar', { param: value })
 * - Template.render('foo_bar_bar', { param: value })
 *
 * @param {String} baseUrl, Script's path calling the loader
 * @param {Array}  templates, List of templates paths
 *
 * @returns {Promise<Templates>}
 */
export default function(baseUrl, templates) {
  if (!Array.isArray(templates)) {
    templates = [templates];
  }

  // const baseUrl = getBaseUrlFromCurrentScript(sourceScript);

  return Promise.all(fetchAll(baseUrl, templates))
    .then(responses => checkResponses(responses))
    .then(responses => appendTemplatesNames(responses))
    .then(responses => bindTemplatesNamesToTextResponses(responses))
    .catch(function(error) {
      throw error;
    });
}
