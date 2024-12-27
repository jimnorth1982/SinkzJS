import fetch from "node-fetch";

/**
 * Fetch item data from the specified URL with an optional JSON payload.
 *
 * @param {string} url - The URL to fetch data from.
 * @param {Object} [payload] - The optional JSON payload to send with the request.
 * @returns {Promise<Object>} - The response data.
 */
export async function fetchREST(url, payload = null) {
  const options = {
    method: payload ? 'POST' : 'GET',
    headers: {
      'Content-Type': 'application/json',
    },
  };

  if (payload) {
    options.body = JSON.stringify(payload);
  }

  const response = await fetch(url, options);
  if (!response.ok) {
    throw new Error(`HTTP error! status: ${response.status}`);
  }
  return response.json();
}

