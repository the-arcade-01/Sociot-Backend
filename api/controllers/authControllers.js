import mongoose from "mongoose";
import bcrypt from "bcryptjs";
import Joi from "joi";
import jwt from "jsonwebtoken";

import User from "../models/User.js";

const userValidationSchema = Joi.object({
  name: Joi.string().min(1).required(),

  username: Joi.string().min(3).max(30).required(),

  email: Joi.string().required().email(),

  password: Joi.string().min(3).required(),
});

const loginValidationSchema = Joi.object({
  email: Joi.string().required().email(),

  password: Joi.string().min(3).required(),
});

const register_user = async (req, res, next) => {
  const validation = userValidationSchema.validate(req.body);
  if (validation.error) {
    return res.status(400).json({
      message: validation.error.details[0].message,
    });
  }

  const emailExist = await User.findOne({
    email: req.body.email,
  });
  if (emailExist) {
    return res.status(400).json({
      message: "Email already exists",
    });
  }

  const usernameExist = await User.findOne({
    username: req.body.username,
  });
  if (usernameExist) {
    return res.status(400).json({
      message: "Username already exists",
    });
  }

  const { name, username, email, password } = req.body;

  const profileImage =
    req.file != null ? req.file.path : "userUploads/default-user.png";

  //   hash password
  const salt = await bcrypt.genSalt(10);
  const hashPassword = await bcrypt.hash(password, salt);

  const user = new User({
    _id: new mongoose.Types.ObjectId(),
    name,
    username,
    email,
    password: hashPassword,
    profileImage,
  });

  await user
    .save()
    .then((user) => {
      console.log(user);
      return res.status(201).json({
        _id: user._id,
        name: user.name,
      });
    })
    .catch((err) => {
      console.log(err);
      return res.status(400).json({
        error: err,
      });
    });
};

const login_user = async (req, res, next) => {
  const validation = loginValidationSchema.validate(req.body);
  if (validation.error) {
    return res.status(400).json({
      message: validation.error.details[0].message,
    });
  }

  const { email, password } = req.body;

  const user = await User.findOne({
    email: email,
  });

  if (!user) {
    return res.status(400).json({
      message: "Email dosn't exists",
    });
  }

  const validPassword = await bcrypt.compare(password, user.password);

  if (!validPassword) {
    return res.status(400).json({
      message: "Invalid password",
    });
  }

  //   create and assign token

  const token = jwt.sign({ _id: user._id }, process.env.SECRET_KEY);

  res.header("auth-token", token);

  return res.status(200).json({
    token: token,
    user: {
      id: user._id,
      name: user.name,
      username: user.username,
      email: user.email,
    },
  });
};

const verify_user = async (req, res, next) => {
  const userID = req.user._id;

  const userData = await User.findById(userID);

  if (!userData) {
    return res.status(400).json({ message: "Invalid user" });
  }
  return res.status(200).json({
    user: {
      _id: userData._id,
      name: userData.name,
      username: userData.username,
      email: userData.email,
    },
  });
};

export default { register_user, login_user, verify_user };
