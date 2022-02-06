import express from "express";

const router = express.Router();

import authControllers from "../controllers/authControllers.js";

router.post("/register", authControllers.register_user);

router.get("/login", authControllers.login_user);

export default router;
