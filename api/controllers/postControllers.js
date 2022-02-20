import mongoose from "mongoose";

import Post from "../models/Post.js";
import Comment from "../models/Comment.js";

const get_all = async (req, res, next) => {
  await Post.find()
    .populate({
      path: "_creator",
      select: "name email",
    })
    .populate({ path: "_comments", select: "text _creator createdAt" })
    .then((posts) => {
      return res.status(200).json({
        count: posts.length,
        posts,
      });
    })
    .catch((err) => {
      console.log(err);
      return res.status(500).json({
        error: err,
      });
    });
};

const get_one = async (req, res, next) => {
  const { _id } = req.params;
  await Post.findOne({ _id: _id })
    .populate({
      path: "_creator",
      select: "name email",
    })
    .populate({ path: "_comments", select: "text _creator createdAt" })
    .then((post) => {
      console.log(post);
      return res.status(200).json({
        post,
      });
    })
    .catch((err) => {
      console.log(err);
      return res.status(500).json({
        error: err,
      });
    });
};

const create_post = async (req, res, next) => {
  const { title, text } = req.body;

  const post = new Post({
    _id: new mongoose.Types.ObjectId(),
    title,
    text,
    postImage: req.file.path,
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

const update_post = async (req, res, next) => {
  const { _id } = req.params;
  const { title, text } = req.body;

  const updatedPost = { _id: _id, title, text, _creator: req.user._id };

  await Post.findByIdAndUpdate(_id, updatedPost, { new: true })
    .then((post) => {
      console.log(post);
      return res.status(200).json({
        message: "Updated Post",
        post,
      });
    })
    .catch((err) => {
      console.log(err);
      return res.status(500).json({
        error: err,
      });
    });
};

const delete_post = async (req, res, next) => {
  const { _id } = req.params;
  await Post.findOneAndDelete({ _id: _id })
    .populate({
      path: "_creator",
      select: "name email",
    })
    .then(async (post) => {
      console.log(post);
      await Comment.deleteMany({ _post: _id });
      return res.status(200).json({
        message: "Post deleted",
        post,
      });
    })
    .catch((error) => {
      console.log(error);
      return res.status(500).json({
        error: error,
      });
    });
};

export default { create_post, get_all, delete_post, get_one, update_post };
