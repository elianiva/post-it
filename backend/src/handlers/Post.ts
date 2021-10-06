import { nanoid } from "nanoid";
import { getRepository } from "typeorm";
import type {
  FastifyInstance,
  FastifyPluginCallback,
  FastifyPluginOptions,
} from "fastify";
import type { IncomingMessage, Server, ServerResponse } from "http";
import { Post } from "#/business/entities/Post";
import { verifyAccessTokenAsync } from "#/platform/jwt/jwt";

export const PostRoutes: FastifyPluginCallback<FastifyPluginOptions, Server> = (
  server: FastifyInstance<Server, IncomingMessage, ServerResponse>,
  _opts,
  done
) => {
  const postRepo = getRepository(Post);

  server.decorateRequest("user", null);
  server.addHook("onRequest", (req, reply, done) => {
    const token = req.headers.authorization?.split(" ")[1];

    if (!token) {
      reply.status(401).send({
        status: 401,
        msg: "Unauthorized!",
        data: [],
      });
      done();
      return;
    }

    try {
      const data = verifyAccessTokenAsync(token);
      // @ts-ignore
      req.user = data;
    } catch (err) {
      reply.send({
        status: 401,
        // @ts-ignore
        msg: err.message,
        data: [],
      });
    } finally {
      done();
    }
  });

  server.get("/", async () => {
    const posts = await postRepo.find();
    return {
      status: 200,
      msg: "Posts retrieved successfully.",
      data: posts,
    };
  });

  server.post("/create", async req => {
    const id = nanoid();
    // @ts-ignore
    const { id: userId } = req.user as { id: string };
    const { content } = req.body as Record<string, string>;

    try {
      const post = await postRepo.save({
        id,
        user_id: userId,
        content,
      });
      return {
        status: 201,
        msg: `Post with ID of ${id} has been successfully created.`,
        data: post,
      };
    } catch (err) {
      return {
        status: 400,
        msg: err,
        data: [],
      };
    }
  });

  done();
};
