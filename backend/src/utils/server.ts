import Fastify, { FastifyInstance, FastifyServerOptions } from "fastify";
import fastifyCookie from "fastify-cookie";
import { IncomingMessage, Server, ServerResponse } from "http";
import { createConnection } from "typeorm";
import { AuthRoutes } from "../routes/Auth";
import { PostRoutes } from "../routes/Post";

const PORT = 3000;

type ServerInstance = FastifyInstance<Server, IncomingMessage, ServerResponse>;

export const createServer = async (
  opts: FastifyServerOptions
): Promise<ServerInstance> => {
  const server = Fastify(opts);

  // connect to the database
  try {
    await createConnection();
  } catch (err) {
    server.log.error(err);
    process.exit(1);
  }

  // register plugins
  server.register(fastifyCookie, {
    secret: process.env.COOKIE_SECRET as string,
  });

  // register routes
  server.get("/", async () => ({
    status: 200,
    msg: "Hello, World!",
  }));
  server.register(AuthRoutes, { prefix: "/api/auth" });
  server.register(PostRoutes, { prefix: "/api/post" });

  server
    .listen(PORT)
    .then(() => console.log(`Server started on port: ${PORT}`))
    .catch(err => {
      server.log.error(err);
      process.exit(1);
    });

  return server;
};
