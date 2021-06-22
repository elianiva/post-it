import { createServer } from "./utils/server";

(async () => {
  await createServer({ logger: true });
})();
