import { RESTDataSource } from "@apollo/datasource-rest";

export class ItemsAPI extends RESTDataSource {
  constructor() {
    super();
    this.baseURL = "http://localhost:8080";
  }

  /**
   *
   * @returns {Promise<ItemsResponse>} The response data.
   */
  async getItems() {
    return this.get("items");
  }

  /**
   * @param {string} id - The item id
   * @returns {Promise<ItemsResponse>} The response data.
   */
  async getItem(id) {
    return this.get(`items/${id}`);
  }

  /**
   * @param {Item} item - The item to create
   * @returns {Promise<ItemsResponse>} The response data.
   */
  async createItem(item) {
    return this.post("items", item);
  }

  /**
   * @param {Item} item - The item to update
   * @returns {Promise<ItemsResponse>} The response data.
   */
  async updateItem(item) {
    return this.put("items", item);
  }

  /**
   * @param {string} id - The item id
   * @returns {Promise<ItemsResponse>} The response data.
   */
  async deleteItem(id) {
    return this.delete(`items/${id}`);
  }
}
