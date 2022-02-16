import mongoose from "mongoose";

import Post from "../models/Post.js";

const create_post = async (req, res, next) => {
  const { title, text } = req.body;

  const post = new Post({
    _id: new mongoose.Types.ObjectId(),
    title,
    text,
    _creator: req.user._id,
  });

  await post
    .save()
    .then((newPost) => {
      console.log(newPost);
      return res.status(201).json({
        newPost,
      });
    })
    .catch((err) => {
      console.log(err);
      return res.status(500).json({
        error: err,
      });
    });
};

export default { create_post };
