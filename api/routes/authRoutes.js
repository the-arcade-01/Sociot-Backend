import express from "express";

const router = express.Router();

import authControllers from "../controllers/authControllers.js";
import verifyToken from "../controllers/verifyToken.js";

router.post("/register", authControllers.register_user);

router.post("/login", authControllers.login_user);

router.post("/verifyUser", verifyToken, authControllers.verify_user);

export default router;
