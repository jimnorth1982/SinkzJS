const itemsHost = "http://localhost:8080";
const itemsPath = "/items";

const exilesHost = "http://localhost:8080";
const exilesPath = "/exiles";

/**
 * 
 * @returns {Url} The URL configuration object.
 */
export const itemsEndpoint = () => {
  return {
    url: `${itemsHost}${itemsPath}`,
    host: itemsHost,
    path: itemsPath,
  };
};

/**
 * 
 * @returns {Url} The URL configuration object.
 */
export const exilesEndpoint = () => {
  return {
    url: `${exilesHost}${exilesPath}`,
    host: exilesHost,
    path: exilesPath,
  };
};
