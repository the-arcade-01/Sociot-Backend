import mongoose from "mongoose";
import Comment from "../models/Comment.js";
import Post from "../models/Post.js";

const create_comment = async (req, res, next) => {
  const { text, postId } = req.body;

  const comment = new Comment({
    _id: new mongoose.Types.ObjectId(),
    text,
    _creator: req.user._id,
    _post: postId,
  });

  await comment
    .save()
    .then(async (newComment) => {
      await Post.findByIdAndUpdate(postId, {
        $push: { _comments: newComment._id },
      })
        .then((post) => {
          console.log(post);
        })
        .catch((err) => {
          console.log(err);
        });
      console.log(newComment);
      return res.status(201).json({
        message: "Comment created",
        newComment,
      });
    })
    .catch((err) => {
      console.log(err);
      return res.status(500).json({
        error: err,
      });
    });
};

export default { create_comment };
