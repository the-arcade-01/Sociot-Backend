import mongoose from "mongoose";

const CommentSchema = new mongoose.Schema({
  _id: mongoose.Schema.Types.ObjectId,
  text: {
    type: String,
    required: true,
  },
  _creator: {
    type: mongoose.Schema.Types.ObjectId,
    ref: "User",
  },
  _post: {
    type: mongoose.Schema.Types.ObjectId,
    ref: "Post",
  },
  createdAt: {
    type: Date,
    default: Date.now,
  },
});

const autoPopulateCreator = function (next) {
  this.populate({
    path: "_creator",
    select: "_id name email username",
  });
  next();
};

CommentSchema.pre("find", autoPopulateCreator);

export default mongoose.model("Comment", CommentSchema);
