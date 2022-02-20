import express from "express";
import multer from "multer";

import authControllers from "../controllers/authControllers.js";
import verifyToken from "../controllers/verifyToken.js";

const router = express.Router();

const storage = multer.diskStorage({
  destination: function (req, file, cb) {
    cb(null, "./userUploads/");
  },
  filename: function (req, file, cb) {
    cb(null, new Date().toISOString() + file.originalname);
  },
});

const fileFilter = (req, file, cb) => {
  if (
    file.mimetype === "image/jpeg" ||
    file.mimetype === "image/jpg" ||
    file.mimetype === "image/png"
  ) {
    cb(null, true);
  } else {
    cb(null, false);
  }
};

const upload = multer({
  storage: storage,
  limits: {
    fileSize: 1024 * 1024 * 5,
  },
  fileFilter: fileFilter,
});

router.post(
  "/register",
  upload.single("profileImage"),
  authControllers.register_user
);

router.post("/login", authControllers.login_user);

router.post("/verifyUser", verifyToken, authControllers.verify_user);

export default router;
