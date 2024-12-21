const { itemsEndpoint } = require("./config");
const middleware = require("./middleware");

const resolvers = {
  Query: {
    items: async () => {
      const rawApiResponse = await middleware.fetchData(itemsEndpoint());
      return rawApiResponse.items;
    },
  },
};

module.exports = resolvers;
