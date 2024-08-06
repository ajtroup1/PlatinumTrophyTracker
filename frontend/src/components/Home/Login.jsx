import { useState } from "react";
import "../../css/Home.css";

function Login() {
  const [loginInfo, setLoginInfo] = useState({
    username: "",
    password: "",
  });

  const [signupInfo, setSignupInfo] = useState({
    username: "",
    password: "",
    confPassword: "",
    firstname: "",
    lastname: "",
    email: "",
    profileImgURL: "",
  });

  const handleLoginChange = (e) => {
    const { id, value } = e.target;
    setLoginInfo((prevInfo) => ({
      ...prevInfo,
      [id]: value,
    }));
  };

  const handleSignupChange = (e) => {
    const { id, value } = e.target;
    setSignupInfo((prevInfo) => ({
      ...prevInfo,
      [id]: value,
    }));
  };

  const handleLoginSubmit = (e) => {
    e.preventDefault();
    console.log("Login Info:", loginInfo);
  };

  const handleSignupSubmit = (e) => {
    e.preventDefault();
    // Handle signup form submission
    console.log("Signup Info:", signupInfo);
    if (signupInfo.password != signupInfo.confPassword) {
      alert("Passwords must match");
      return;
    }
    payload = {
      username: signupInfo.username,
      password: signupInfo.password,
      firstname: signupInfo.firstname,
      lastname: signupInfo.lastname,
      email: signupInfo.email,
      profileURL: signupInfo.profileImgURL,
    };
  };

  return (
    <div className="login-main">
      <div className="login-container">
        <p className="login-head-text">Log in</p>
        <form onSubmit={handleLoginSubmit} className="login-form">
          <div className="form-group">
            <label htmlFor="username">Username / Email</label>
            <input
              type="text"
              id="username"
              value={loginInfo.username}
              onChange={handleLoginChange}
              required
            />
          </div>
          <div className="form-group">
            <label htmlFor="password">Password</label>
            <input
              type="password"
              id="password"
              value={loginInfo.password}
              onChange={handleLoginChange}
              autoComplete="off"
              required
            />
          </div>
          <button type="submit" className="login-button">
            Log in
          </button>
        </form>
      </div>
      <p className="signup-text">Don't have an account? Sign up below!</p>
      <div className="signup-container">
        <p className="login-head-text">Sign up</p>
        <form onSubmit={handleSignupSubmit} className="login-form">
          <div className="form-group">
            <label htmlFor="username">Username</label>
            <input
              type="text"
              id="username"
              value={signupInfo.username}
              onChange={handleSignupChange}
              required
            />
          </div>
          <div className="form-group-split">
            <div>
              <label htmlFor="password">Password</label>
              <input
                type="password"
                id="password"
                value={signupInfo.password}
                autoComplete="off"
                onChange={handleSignupChange}
                required
              />
            </div>
            <div>
              <label htmlFor="confPassword">Confirm Password</label>
              <input
                type="password"
                id="confPassword"
                value={signupInfo.confPassword}
                autoComplete="off"
                onChange={handleSignupChange}
                required
              />
            </div>
          </div>
          <div className="form-group">
            <label htmlFor="email">Email</label>
            <input
              type="email"
              id="email"
              value={signupInfo.email}
              onChange={handleSignupChange}
              required
            />
          </div>
          <div className="form-group-split">
            <div>
              <label htmlFor="firstname">First Name</label>
              <input
                type="text"
                id="firstname"
                value={signupInfo.firstname}
                onChange={handleSignupChange}
                required
              />
            </div>
            <div>
              <label htmlFor="lastname">Last Name</label>
              <input
                type="text"
                id="lastname"
                value={signupInfo.lastname}
                onChange={handleSignupChange}
                required
              />
            </div>
          </div>
          <div className="form-group">
            <label htmlFor="profileImgURL">Profile Image URL</label>
            <input
              type="text"
              id="profileImgURL"
              value={signupInfo.profileImgURL}
              onChange={handleSignupChange}
            />
          </div>
          <button type="submit" className="login-button">
            Register
          </button>
        </form>
      </div>
    </div>
  );
}

export default Login;
