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

  repeat_password: Joi.ref("password"),
});

const loginValidationSchema = Joi.object({
  email: Joi.string().required().email(),

  password: Joi.string().min(3).required(),
});

const register_user = async (req, res, next) => {
  const validation = userValidationSchema.validate(req.body);
  if (validation.error) {
    return res.status(400).json({
      error: validation.error.details[0].message,
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

  //   hash password
  const salt = await bcrypt.genSalt(10);
  const hashPassword = await bcrypt.hash(password, salt);

  const user = new User({
    _id: new mongoose.Types.ObjectId(),
    name,
    username,
    email,
    password: hashPassword,
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
  return res.send({
    message: "login",
  });
};

export default { register_user, login_user };
