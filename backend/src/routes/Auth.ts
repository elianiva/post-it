import {
  FastifyInstance,
  FastifyPluginCallback,
  FastifyPluginOptions,
} from "fastify";
import { IncomingMessage, Server, ServerResponse } from "http";
import { nanoid } from "nanoid";
import { User } from "../entities/User";
import { compare, hash } from "../utils/hash";
import { createAccessToken, createRefreshToken } from "../utils/jwt";
import { getRepository } from "typeorm";

export const AuthRoutes: FastifyPluginCallback<FastifyPluginOptions, Server> = (
  server: FastifyInstance<Server, IncomingMessage, ServerResponse>,
  _opts,
  done
) => {
  const userRepo = getRepository(User);

  server.post("/register", async (req, reply) => {
    const { email, username, password } = req.body as Record<string, string>;
    if (!email || !username || !password) {
      return reply.status(400).send({
        status: 400,
        msg: "Invalid request!",
        data: [],
      });
    }

    const id = nanoid();

    const user = await userRepo.findOne({
      where: { email },
    });

    if (user) {
      return {
        status: 400,
        msg: `User with email of ${email} is already exists!`,
        data: [],
      };
    }

    try {
      await userRepo.save({
        id,
        email,
        username,
        password: await hash(password),
      });

      return reply.status(201).send({
        status: 201,
        msg: `User with ID of ${id} has been successfully registered`,
        data: [],
      });
    } catch (err) {
      server.log.error(err);
      return reply.status(500).send({
        status: 500,
        msg: "Internal server error",
        data: [],
      });
    }
  });

  server.post("/login", async (req, reply) => {
    const { email, username, password } = req.body as Record<string, string>;

    const user = await User.findOne({
      where: [{ email }, { username }],
    });

    if (!user) {
      return {
        status: 400,
        msg: `User with email of ${email} does not exists!`,
        data: [],
      };
    }

    const isVerified = await compare(password, user.password);

    if (!isVerified) {
      return {
        status: 400,
        msg: `Wrong password!`,
        data: [],
      };
    }

    const refreshToken = createRefreshToken({id: user.id});
    const accessToken = createAccessToken({id: user.id});

    reply
      .setCookie("_tkn", refreshToken, {
        // save the refresh token in the httpOnly cookie
        httpOnly: true,
        signed: true,
      })
      .send({
        status: 200,
        msg: "Logged in successfully",
        data: {
          // save the actual token in the memory later
          token: accessToken,
        },
      });
  });
  done();
};
