import { fetchItemData } from "./middleware.js";
import { itemsEndpoint } from "./config.js";

export const resolvers = {
  Query: {
    items: async () => {
      const rawApiResponse = await fetchItemData(itemsEndpoint()).catch(
        (error) => {
          return { message: error, items: [] };
        }
      );
      return rawApiResponse.items;
    },
  },
};
