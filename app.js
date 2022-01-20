import express from "express";

const app = express();

app.get("/", (req, res, next) => {
  return res.json({
    message: "working",
  });
});

export default app;
