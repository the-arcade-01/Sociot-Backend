import express from "express";
import dotenv from "dotenv";
import cors from "cors";
import morgan from "morgan";
import mongoose from "mongoose";

dotenv.config();

const app = express();

// db connect

mongoose.connect(process.env.DB_CONNECT, () => {
  console.log("Connected to mongoDB");
});

// middleware

app.use(cors());
app.use(morgan("dev"));
app.use(express.urlencoded({ extended: false }));
app.use(express.json());

// routes

import route from "./api/routes/route.js";
import authRoutes from "./api/routes/authRoutes.js";
import postRoutes from "./api/routes/postRoutes.js";

app.use("/", route);
app.use("/api/", authRoutes);
app.use("/api/posts", postRoutes);

export default app;
