import express from "express";
import multer from "multer";

import postControllers from "../controllers/postControllers.js";
import verifyToken from "../controllers/verifyToken.js";

const router = express.Router();

const storage = multer.diskStorage({
  destination: function (req, file, cb) {
    cb(null, "./postUploads/");
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

router.get("/", postControllers.get_all);

router.get("/:_id", verifyToken, postControllers.get_one);

router.post(
  "/create",
  verifyToken,
  upload.single("postImage"),
  postControllers.create_post
);

router.patch("/:_id", verifyToken, postControllers.update_post);

router.delete("/:_id", verifyToken, postControllers.delete_post);

export default router;
