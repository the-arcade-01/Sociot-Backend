import mongoose from "mongoose";
import Comment from "../models/Comment.js";
import Post from "../models/Post.js";

const get_user_comments = async (req, res, next) => {
  const _id = req.user._id;
  await Comment.find({ _creator: _id })
    .then((comments) => {
      return res.status(200).json({
        comments,
      });
    })
    .catch((err) => {
      console.log(err);
      return res.status(500).json({
        error: err,
      });
    });
};

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

const delete_comment = async (req, res, next) => {
  const { _id } = req.params;
  await Comment.findOneAndDelete({ _id: _id })
    .then((comment) => {
      console.log(comment);
      return res.status(200).json({
        message: "Comment Deleted",
        comment,
      });
    })
    .catch((err) => {
      console.log(err);
      return res.status(500).json({
        error: err,
      });
    });
};

export default { create_comment, get_user_comments, delete_comment };
