import mongoose from "mongoose";

const PostSchema = new mongoose.Schema({
  _id: mongoose.Schema.Types.ObjectId,
  title: {
    type: String,
    required: true,
  },
  text: {
    type: String,
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
