import { Request, Response } from "express";

export const handleCreatePlaylist = async (req: Request, res: Response) => {
	res.send("create");
};

export const handleModifyList = async (req: Request, res: Response) => {
	res.send("modify")
}

export const handleDeleteList = async (req: Request, res: Response) => {
	res.send("delete list")
}
export const handleGetList = async (req: Request, res: Response) => {
	res.send("get list")
}
