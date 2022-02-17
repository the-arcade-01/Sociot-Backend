import express from "express";
import postControllers from "../controllers/postControllers.js";

import verifyToken from "../controllers/verifyToken.js";

const router = express.Router();

router.get("/", postControllers.get_all);

router.get("/:_id", verifyToken, postControllers.get_one);

router.post("/create", verifyToken, postControllers.create_post);

router.patch("/:_id", verifyToken, postControllers.update_post);

router.delete("/:_id", verifyToken, postControllers.delete_post);

export default router;
