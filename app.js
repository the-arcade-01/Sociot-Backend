import express from "express";
import dotenv from "dotenv";
import cors from "cors";
import morgan from "morgan";
import mongoose from "mongoose";

// routes

import route from "./api/routes/route.js";
import authRoutes from "./api/routes/authRoutes.js";
import postRoutes from "./api/routes/postRoutes.js";
import commmentRoutes from "./api/routes/commentRoutes.js";

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
app.use("/postUploads", express.static("postUploads"));
app.use("/userUploads", express.static("userUploads"));

// routes

app.use("/", route);
app.use("/api/", authRoutes);
app.use("/api/posts", postRoutes);
app.use("/api/comments", commmentRoutes);

export default app;
