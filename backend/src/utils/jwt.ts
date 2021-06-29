import * as jwt from "jsonwebtoken";
import "dotenv/config";

export const verifyAccessToken = async (token: string): Promise<unknown> => {
  return new Promise((resolve, reject) => {
    jwt.verify(token, process.env.ACCESS_TOKEN_SECRET!, (err, data) => {
      if (err) reject(err);
      resolve(data);
    });
  });
};

export const verifyRefreshToken = async (token: string): Promise<unknown> => {
  return new Promise((resolve, reject) => {
    jwt.verify(token, process.env.REFRESH_TOKEN_SECRET!, (err, data) => {
      if (err) reject(err);
      resolve(data);
    });
  });
};

export const createAccessToken = (payload: { id: string }): string => {
  return jwt.sign(payload, process.env.ACCESS_TOKEN_SECRET!, {
    expiresIn: "15m",
  });
};

export const createRefreshToken = (payload: { id: string}): string => {
  return jwt.sign(payload, process.env.REFRESH_TOKEN_SECRET!, {
    expiresIn: "7d",
  });
};
