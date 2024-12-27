import { fetchREST } from "./middleware.js";
import { exilesEndpoint, itemsEndpoint } from "./config.js";

export const resolvers = {
  Query: {
    items: async (_, __, { dataSources }) => {
      const data = await dataSources.itemsAPI.getItems();
      console.log(data);
      return data;
    },
    item: async (_, { id }, { dataSources }) => {
      return await dataSources.itemsAPI.getItem(id);
    },
    exiles: async (_, __, { dataSources }) => {
      return await dataSources.exileAPI.getExiles();
    },
    exile: async (_, { id }, { dataSources }) => {
      return await dataSources.exileAPI.getExile(id);
    },
  },
};
