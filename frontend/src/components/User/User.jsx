import { useState } from "react";
import { Outlet } from "react-router-dom";
import "../../css/User.css";
import Navbar from "../Common/UserNavbar";

function User() {
  return (
    <div className="user-main">
      <Navbar />
      <div className="user-middle">
        {/* RENDER CONTENTS IN HERE */}
        <Outlet />
      </div>
    </div>
  );
}

export default User;
