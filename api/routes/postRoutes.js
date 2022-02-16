import express from "express";
import postControllers from "../controllers/postControllers.js";

import verifyToken from "../controllers/verifyToken.js";

const router = express.Router();

router.get("/", verifyToken, (req, res, next) => {
  return res.send({ message: "from posts" });
});

router.post("/create", verifyToken, postControllers.create_post);

export default router;
