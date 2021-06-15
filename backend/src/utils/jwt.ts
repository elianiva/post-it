import * as jwt from "jsonwebtoken";
import dotenv from "dotenv";
dotenv.config();

const SECRET = process.env.JWT_SECRET as string;
const REFRESH_SECRET = process.env.JWT_REFRESH_SECRET as string;

export const sign = (
  data: Record<string, string>,
  refreshToken: boolean
): string => {
  return jwt.sign(data, refreshToken ? REFRESH_SECRET : SECRET, {
    expiresIn: "10m",
  });
};

export const verify = async (token: string): Promise<unknown> => {
  return new Promise((resolve, reject) => {
    jwt.verify(token, SECRET, (err, data) => {
      if (err) reject(err);
      resolve(data);
    });
  });
};
