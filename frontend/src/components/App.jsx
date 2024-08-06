import { useState } from 'react'
import { Routes, Route, useLocation } from "react-router-dom";
import '../css/App.css'
import User from "./User/User"
import UserHome from "./User/Home"
import Profile from "./User/Profile"
import UserGames from "./User/Games"
import UserStats from "./User/Statistics"
import UserCompleted from "./User/CompletedGames"
import Home from './Home/Home'
import Login from './Home/Login'
import Navbar from './Common/Navbar'

function App() {
  const location = useLocation();
  const isUserRoute = location.pathname.startsWith("/user");
  return (
    <div className="main">
      {!isUserRoute && <Navbar />}
      <Routes>
        <Route path="/" element={<Home />} />
        <Route path="/login" element={<Login />} />
        <Route path="user" element={<User />}>
          <Route index element={<UserHome />} />
          <Route path="profile" element={<Profile />} />
          <Route path="games" element={<UserGames />} />
          <Route path="stats" element={<UserStats />} />
          <Route path="completed" element={<UserCompleted />} />
        </Route>
      </Routes>
    </div>
  );
}

export default App
