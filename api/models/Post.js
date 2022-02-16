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
  createdAt: {
    type: Date,
    default: Date.now,
  },
  _creator: {
    type: mongoose.Schema.ObjectId,
    ref: "User",
  },
});

export default mongoose.model("Post", PostSchema);
