import fetch from "node-fetch";

/**
 * Fetch data from a given URL.
 * @param {string} url - The URL to fetch data from.
 * @returns {Promise<ItemsResponse>} The fetched data.
 * @throws Will throw an error if the fetch operation fails.
 */
export async function fetchItemData(url) {
  try {
    const response = await fetch(url);
    if (!response.ok) {
      throw new Error(`HTTP error! status: ${response.status}`);
    }
    const data = await response.json();
    return data;
  } catch (error) {
    console.error("Error fetching data:", error);
    throw error;
  }
}