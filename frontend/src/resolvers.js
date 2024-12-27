import { fetchREST } from "./middleware.js";
import { exilesEndpoint, itemsEndpoint } from "./config.js";

export const resolvers = {
  Query: {
    items: async () => {
      const rawApiResponse = await fetchREST(itemsEndpoint().url).catch(
        (error) => {
          return { message: error, items: [] };
        }
      );
      return rawApiResponse.items;
    },
    item: async (_, { id }) => {
      const rawApiResponse = await fetchREST(itemsEndpoint().url, {id: id}).catch(
        (error) => {
          return { message: error, items: [] };
        }
      );
      return rawApiResponse.items.find((item) => item.id === id);
    },
    exiles: async () => {
      const rawApiResponse = await fetchREST(exilesEndpoint().url).catch(
        (error) => {
          return { message: error, exiles: [] };
        }
      );
      return rawApiResponse.exiles;
    },
    exile: async (_, { id }) => {
      const rawApiResponse = await fetchREST(exilesEndpoint().url, {id: id}).catch(
        (error) => {
          return { message: error, exiles: [] };
        }
      );
      return rawApiResponse.exiles.find((exile) => exile.id === id);
    },
  },
};
