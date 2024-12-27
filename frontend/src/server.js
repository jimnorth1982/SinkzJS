import { ApolloServer } from "@apollo/server";
import { expressMiddleware } from "@apollo/server/express4";
import { ApolloServerPluginDrainHttpServer } from "@apollo/server/plugin/drainHttpServer";
import express from "express";
import http from "http";
import cors from "cors";
import { readFileSync } from "fs";
import { fileURLToPath } from "url";
import { dirname, join } from "path";

import { resolvers } from "./resolvers.js";
import { ExileAPI } from "./exileApi.js";
import { config } from "process";
import { itemsEndpoint } from "./config.js";
import { ItemsAPI } from "./itemsApi.js";

const pathUrl = dirname(fileURLToPath(import.meta.url));

const typeDefs = readFileSync(
  join(pathUrl, "../schema/schema.graphql"),
  "utf8"
);

/**
 * @typedef {Object} ContextValue
 * @property {Object} config
 * @property {Object} dataSources
 * @property {ExileAPI} dataSources.exileAPI
 * @property {ItemsAPI} dataSources.itemsAPI
 */

export async function startApolloServer(port = 4001) {
  const app = express();
  const httpServer = http.createServer(app);

  const server = new ApolloServer({
    typeDefs,
    resolvers,
    plugins: [ApolloServerPluginDrainHttpServer({ httpServer })],
  });

  await server.start();

  app.use(
    "/",
    cors(),
    express.json(),
    expressMiddleware(server, {
      context: async ({ req }) => ({
        token: req.headers.token,
        dataSources: {
          exileAPI: new ExileAPI(),
          itemsAPI: new ItemsAPI(),
        },
      }),
    })
  );

  httpServer.listen({ port }, () => {
    console.log(`🚀 Server ready at http://localhost:${port}`);
  });
}
