import { createServer } from "./server";

(async () => {
  await createServer({ logger: true });
})();
