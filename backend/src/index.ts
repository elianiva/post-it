import Fastify, { FastifyInstance } from "fastify";
import { IncomingMessage, Server, ServerResponse } from "http";

class App {
  private PORT = 3000;
  private server: FastifyInstance<Server, IncomingMessage, ServerResponse>;

  constructor(
    server: FastifyInstance<Server, IncomingMessage, ServerResponse>
  ) {
    this.server = server;
    this.init();
    this.registerRoutes();
  }

  init() {
    this.server
      .listen(this.PORT)
      .then(() => console.log(`Server started on port: ${this.PORT}`))
      .catch(err => {
        this.server.log.error(err);
        process.exit(1);
      });
  }

  registerRoutes() {
    this.server.get("/", async () => ({ hello: "world" }));
  }
}

new App(Fastify({ logger: true }));
