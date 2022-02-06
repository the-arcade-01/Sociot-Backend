import express from "express";

const router = express.Router();
import verifyToken from "../controllers/verifyToken.js";

router.get("/", (req, res, next) => {
  return res.send({
    message: "working from route",
  });
});

router.get("/private", verifyToken, (req, res, next) => {
  return res.send({
    message: "Private Route",
  });
});

export default router;
