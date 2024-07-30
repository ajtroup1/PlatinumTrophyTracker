import { useState } from "react";
import "../../css/User.css";

function Games() {
  const trackedGames = [
    {
      id: 1,
      title: "Dark Souls 3",
      achievementsComp: 70,
      achievementsTot: 80,
      img: "https://m.media-amazon.com/images/M/MV5BYzJhYTgzYzYtYjdjOC00ZDYyLTg0NjYtZDEwMDlkODA3OWI4XkEyXkFqcGdeQXVyMTk2OTAzNTI@._V1_.jpg",
    },
    {
      id: 2,
      title: "Bloodborne",
      achievementsComp: 50,
      achievementsTot: 60,
      img: "https://upload.wikimedia.org/wikipedia/en/6/68/Bloodborne_Cover_Wallpaper.jpg",
    },
    {
      id: 3,
      title: "Sekiro",
      achievementsComp: 30,
      achievementsTot: 40,
      img: "https://image.api.playstation.com/vulcan/img/rnd/202010/2723/knxU5uU5aKvQChKX5OvWtSGC.png",
    },
    {
      id: 4,
      title: "The Witcher 3",
      achievementsComp: 20,
      achievementsTot: 50,
      img: "https://upload.wikimedia.org/wikipedia/en/thumb/0/0c/Witcher_3_cover_art.jpg/220px-Witcher_3_cover_art.jpg",
    },
    {
      id: 5,
      title: "Hades",
      achievementsComp: 80,
      achievementsTot: 100,
      img: "https://upload.wikimedia.org/wikipedia/en/thumb/c/cc/Hades_cover_art.jpg/220px-Hades_cover_art.jpg",
    },
    {
      id: 6,
      title: "Celeste",
      achievementsComp: 90,
      achievementsTot: 120,
      img: "https://upload.wikimedia.org/wikipedia/commons/0/0f/Celeste_box_art_full.png",
    },
    {
      id: 7,
      title: "Stardew Valley",
      achievementsComp: 40,
      achievementsTot: 70,
      img: "https://upload.wikimedia.org/wikipedia/en/f/fd/Logo_of_Stardew_Valley.png",
    },
    {
      id: 8,
      title: "Minecraft",
      achievementsComp: 10,
      achievementsTot: 30,
      img: "https://assets.nintendo.com/image/upload/ar_16:9,c_lpad,w_1240/b_white/f_auto/q_auto/ncom/software/switch/70010000000964/a28a81253e919298beab2295e39a56b7a5140ef15abdb56135655e5c221b2a3a",
    },
  ];


  const getProgressBarColor = (percentage) => {
    if (percentage < 10) return "#f44336"; // Very Red
    if (percentage < 20) return "#e57373"; // Light Red
    if (percentage < 30) return "#ff8a65"; // Coral
    if (percentage < 40) return "#ffb74d"; // Light Orange
    if (percentage < 50) return "#ff9800"; // Orange
    if (percentage < 60) return "#ffeb3b"; // Yellow
    if (percentage < 70) return "#cddc39"; // Lime
    if (percentage < 80) return "#8bc34a"; // Light Green
    if (percentage < 90) return "#4caf50"; // Green
    if (percentage < 100) return "#388e3c"; // Dark Green
    return "#2c6b2f"; // Very Dark Green
  };

  return (
    <div className="user-games-main">
        <div className="user-games-header-container">
            <p className="tracked-games-title">Your tracked games...</p>
        </div>
      <div className="tracked-games-list">
        {trackedGames.map((game) => {
          const progressPercentage =
            (game.achievementsComp / game.achievementsTot) * 100;

          return (
            <div className="indiv-track-game-container" key={game.id}>
              <div className="tracked-game-img-container">
                <img
                  src={game.img}
                  alt={game.title}
                  className="tracked-game-img"
                />
              </div>
              <div className="right-tracked-container">
                <p>{game.title}</p>
                <div className="progress-container">
                  <div className="progress-bar">
                    <div
                      className="progress-bar-fill"
                      style={{
                        width: `${progressPercentage}%`,
                        backgroundColor:
                          getProgressBarColor(progressPercentage),
                      }}
                    >
                      <span className="progress-text">
                        {game.achievementsComp} / {game.achievementsTot}
                      </span>
                    </div>
                  </div>
                </div>
              </div>
            </div>
          );
        })}
      </div>
    </div>
  );
}

export default Games;
