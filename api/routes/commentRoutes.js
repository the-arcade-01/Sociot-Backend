import express from "express";
import commentControllers from "../controllers/commentControllers.js";
import verifyToken from "../controllers/verifyToken.js";

const router = express.Router();

router.get("/userComments", verifyToken, commentControllers.get_user_comments);

router.post("/create", verifyToken, commentControllers.create_comment);

router.delete("/:_id", verifyToken, commentControllers.delete_comment);

export default router;
