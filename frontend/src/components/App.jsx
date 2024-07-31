import { useState } from 'react'
import { Routes, Route } from 'react-router-dom'
import '../css/App.css'
import User from "./User/User"
import UserHome from "./User/Home"
import Profile from "./User/Profile"
import UserGames from "./User/Games"
import UserStats from "./User/Statistics"
import UserCompleted from "./User/CompletedGames"

function App() {
  return (
    <div className="main">
      <Routes>
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
