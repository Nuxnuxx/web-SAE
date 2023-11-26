import { Router } from "express";
import { handleCreatePlaylist, handleDeleteList, handleGetList } from "./playlist.controller";

const playlistRouter = Router();

playlistRouter.post("/createList", handleCreatePlaylist)
playlistRouter.put("/modifyList", handleCreatePlaylist)
playlistRouter.delete("/deleteList", handleDeleteList)
playlistRouter.get("/displayLists", handleGetList)

export default playlistRouter;
