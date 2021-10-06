import { verify, sign } from "jsonwebtoken";
import "dotenv/config";

export async function verifyAccessTokenAsync(token: string): Promise<unknown> {
  return new Promise((resolve, reject) => {
    verify(token, process.env.ACCESS_TOKEN_SECRET!, (err, data) => {
      if (err) reject(err);
      resolve(data);
    });
  });
}

export async function verifyRefreshTokenAsync(token: string): Promise<unknown> {
  return new Promise((resolve, reject) => {
    verify(token, process.env.REFRESH_TOKEN_SECRET!, (err, data) => {
      if (err) reject(err);
      resolve(data);
    });
  });
}

export function createAccessToken(payload: { id: string }): string {
  return sign(payload, process.env.ACCESS_TOKEN_SECRET!, {
    expiresIn: "15m",
  });
}

export function createRefreshToken(payload: { id: string }): string {
  return sign(payload, process.env.REFRESH_TOKEN_SECRET!, {
    expiresIn: "7d",
  });
}
