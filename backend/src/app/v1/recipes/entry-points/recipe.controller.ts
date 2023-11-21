import { Request, Response } from "express";

export const getProduct = async (req: Request, res: Response) => {
	const { id } = req.params;
	// const result = await GetRecipeById(id)
	res.send(id);
};
