import { RESTDataSource } from "@apollo/datasource-rest";
// startApolloServer(4001);

export class ExileAPI extends RESTDataSource {
  constructor() {
    super();
    this.baseURL = "http://localhost:8080";
  }

  /**
   *
   * @returns {Promise<ExilesResponse>} The response data.
   */
  async getExiles() {
    return this.get("exiles");
  }

  /**
   * @param {string} id - The exile id
   * @returns {Promise<ExilesResponse>} The response data.
   */
  async getExile(id) {
    return this.get(`exiles/${id}`);
  }
}
