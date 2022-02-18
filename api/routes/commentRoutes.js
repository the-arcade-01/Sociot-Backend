import express from "express";
import commentControllers from "../controllers/commentControllers.js";
import verifyToken from "../controllers/verifyToken.js";

const router = express.Router();

router.get("/", (req, res, next) => {
  return res.send({
    message: "from comment",
  });
});

router.post("/create", verifyToken, commentControllers.create_comment);

export default router;
