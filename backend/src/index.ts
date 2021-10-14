import Fastify from "fastify";

const PORT = process.env.PORT || 3000;

const app = Fastify({ logger: true });

app.get("/", async () => {
  return "Hello World";
});

app.listen(PORT).then(() => console.log(`Server Listening on port ${PORT}`));
