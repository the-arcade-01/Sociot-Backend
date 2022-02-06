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

app.use("/", route);

export default app;
