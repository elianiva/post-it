import dotenv from "dotenv";
import Fastify, { FastifyInstance } from "fastify";
import fastifyCookie from "fastify-cookie";
import { IncomingMessage, Server, ServerResponse } from "http";
import { createConnection } from "typeorm";
import { Post } from "./entities/Post";
import { User } from "./entities/User";
import { AuthRoutes } from "./routes/Auth";
import { PostRoutes } from "./routes/Post";

dotenv.config();
const { DB_HOST, DB_NAME, DB_USERNAME, DB_PASSWORD } = process.env;

export class App {
  private PORT = 3000;
  public server: FastifyInstance<Server, IncomingMessage, ServerResponse>;

  constructor(
    server: FastifyInstance<Server, IncomingMessage, ServerResponse>
  ) {
    this.server = server;
    (async () => {
      // wait for the database to be ready
      await this.connectDatabase();
      this.registerPlugins();
      this.registerRoutes();
      this.init();
    })();
  }

  init(): void {
    this.server
      .listen(this.PORT)
      .then(() => console.log(`Server started on port: ${this.PORT}`))
      .catch(err => {
        this.server.log.error(err);
        process.exit(1);
      });
  }

  registerRoutes(): void {
    this.server.get("/", async () => ({ hello: "world" }));

    this.server.register(AuthRoutes, { prefix: "/api/auth" });
    this.server.register(PostRoutes, { prefix: "/api/post" });
  }

  registerPlugins(): void {
    this.server.register(fastifyCookie, {
      secret: process.env.COOKIE_SECRET as string,
    });
  }

  async connectDatabase(): Promise<void> {
    try {
      await createConnection({
        type: "postgres",
        host: DB_HOST,
        database: DB_NAME,
        username: DB_USERNAME,
        password: DB_PASSWORD,
        synchronize: true,
        logging: false,
        entities: [User, Post],
      });
    } catch (err) {
      this.server.log.error(err);
      process.exit(1);
    }
  }
}

new App(Fastify({ logger: true }));
