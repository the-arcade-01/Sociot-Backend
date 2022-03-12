import mongoose from "mongoose";

const PostSchema = new mongoose.Schema({
  _id: mongoose.Schema.Types.ObjectId,
  text: {
    type: String,
    required: true,
  },
  link: {
    type: String,
  },
  category: {
    type: String,
    required: true,
    default: "All",
  },
  postImage: { type: String },
  createdAt: {
    type: Date,
    default: Date.now,
  },
  _creator: {
    type: mongoose.Schema.ObjectId,
    ref: "User",
  },
  _comments: [{ type: mongoose.Schema.ObjectId, ref: "Comment" }],
});

export default mongoose.model("Post", PostSchema);
